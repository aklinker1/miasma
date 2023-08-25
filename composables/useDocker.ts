export default function () {
  const router = useRouter();
  const route = useRoute();
  const authHeader = useAuthHeader();

  const fetchWithAuth = $fetch.create({
    onRequest(ctx) {
      if (authHeader.value == null) return;

      const headers = new Headers(ctx.options.headers);
      headers.set('Authorization', authHeader.value);
      ctx.options.headers = headers;
    },
    onResponse(context) {
      if (context.response.status === 403 && route.path !== routes.login) {
        router.push({
          path: routes.login,
          query: { redirect: route.fullPath },
        });
      }
    },
  });

  const docker = {
    getSwarmInfo(): Promise<Docker.GetSwarmInspectResponse200> {
      const url = new DockerURL(`/swarm`);
      return fetchWithAuth(url.href);
    },

    getSystemInfo(): Promise<Docker.GetSystemInfoResponse200> {
      const url = new DockerURL(`/info`);
      return fetchWithAuth(url.href);
    },

    async getService(id: string): Promise<Docker.Service> {
      const res = await docker.listServices({
        filters: { id: [id] },
        status: true,
      });
      return res[0];
    },

    async listServices(options?: ListServiceOptions): Promise<Docker.GetServiceListResponse200> {
      let url = new DockerURL(`/services`);

      if (options?.filters != null) {
        url.searchParams.set('filters', JSON.stringify(options.filters));
      }
      if (options?.status != null) {
        url.searchParams.set('status', String(options.status));
      }

      return await fetchWithAuth(url.href);
    },

    createService(spec: Docker.ServiceSpec): Promise<Docker.PostServiceCreateResponse201> {
      const url = new DockerURL(`/services/create`);
      return fetchWithAuth(url.href, {
        method: 'POST',
        body: spec,
      });
    },

    async deleteService(service: Pick<Docker.Service, 'ID'>): Promise<void> {
      const url = new DockerURL(`/services/${service.ID}`);
      await $fetch(url.href, {
        method: 'DELETE',
      });
    },

    async updateService(
      service: Docker.Service,
      newSpec: Docker.ServiceSpec,
    ): Promise<Docker.PostServiceUpdateResponse200> {
      const url = new DockerURL(`/services/${service.ID}/update`);
      url.searchParams.set('version', String(service.Version?.Index ?? 0));

      return fetchWithAuth(url.href, { method: 'POST', body: newSpec });
    },

    async renameService(
      service: Docker.Service,
      newSpec: Docker.ServiceSpec,
    ): Promise<Docker.PostServiceCreateResponse201> {
      const oldSpec = service.Spec!;

      // Delete the old service to clear up ports
      await docker.deleteService(service);
      try {
        return await docker.createService(newSpec);
      } catch (err) {
        // Restore old service if new create failed
        await docker.createService(oldSpec);
        throw err;
      }
    },

    async startService(service: Docker.Service): Promise<Docker.PostServiceUpdateResponse200> {
      const desiredTasks = service.ServiceStatus?.DesiredTasks ?? 0;
      if (desiredTasks > 0) throw Error('Service is already running');

      const newSpec = clone(service.Spec!);
      const replicas = Number(service.Spec?.Labels?.[MiasmaLabels.InstanceCount] ?? '1');
      newSpec.Mode ??= {};
      newSpec.Mode.Replicated ??= {};
      newSpec.Mode.Replicated.Replicas = replicas;

      return await docker.updateService(service, newSpec);
    },

    async stopService(service: Docker.Service): Promise<Docker.PostServiceUpdateResponse200> {
      const desiredTasks = service.ServiceStatus?.DesiredTasks ?? 0;
      if (desiredTasks === 0) throw Error('Service is already stopped');

      const newSpec = clone(service.Spec!);
      newSpec.Mode ??= {};
      newSpec.Mode.Replicated ??= {};
      newSpec.Mode.Replicated.Replicas = 0;

      return await docker.updateService(service, newSpec);
    },

    listNodes(options?: ListNodeOptions): Promise<Docker.GetNodeListResponse200> {
      const url = new DockerURL(`/nodes`);

      if (options?.filters) {
        url.searchParams.set('filters', JSON.stringify(options.filters));
      }

      return fetchWithAuth(url.href);
    },

    async getNode(id: string): Promise<Docker.Node> {
      const res = await docker.listNodes({
        filters: { id: [id] },
      });
      return res[0];
    },

    async updateNode(node: Docker.Node, newSpec: Docker.NodeSpec): Promise<void> {
      const url = new DockerURL(`/nodes/${node.ID}/update`);
      url.searchParams.set('version', String(node.Version?.Index ?? 0));

      return fetchWithAuth(url.href, { method: 'POST', body: newSpec });
    },

    listTasks(options?: ListTaskOptions): Promise<Docker.GetTaskListResponse200> {
      const url = new DockerURL(`/tasks`);

      if (options?.filters) {
        url.searchParams.set('filters', JSON.stringify(options.filters));
      }

      return fetchWithAuth(url.href);
    },

    async pullImage(image: string): Promise<void> {
      const url = new DockerURL(`/images/create`);
      const body = {
        fromImage: image,
      };
      await $fetch(url.href, {
        method: 'POST',
        body,
      });
    },

    async pullLatest(service: Docker.Service): Promise<Docker.PostServiceUpdateResponse200> {
      const image = service.Spec?.TaskTemplate?.ContainerSpec?.Image;
      if (image == null) {
        console.warn('Image missing in service:', service);
        throw Error('Service does not have an image');
      }
      await this.pullImage(image);
      return await this.updateService(service, service.Spec!);
    },
  };

  return docker;
}

/**
 * URL class that strips the origin from the `href` getter.
 */
class DockerURL extends URL {
  static ORIGIN = 'http://docker';
  static BASE_URL = `${DockerURL.ORIGIN}/api/docker`;

  constructor(path: string) {
    super(DockerURL.BASE_URL + path);
  }

  get href() {
    return super.href.replace(DockerURL.ORIGIN, '');
  }
}

export interface ListServiceOptions {
  filters?: {
    id?: string[];
    label?: string[];
    mode?: Array<'replicated' | 'global'>;
    name?: string[];
  };
  status?: boolean;
}

export interface ListNodeOptions {
  filters?: {
    id?: string[];
  };
}

export interface ListTaskOptions {
  filters?: unknown;
}
