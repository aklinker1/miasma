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
  filters?: unknown;
}

export interface ListTaskOptions {
  filters?: unknown;
}

/**
 * URL class that strips the origin from the `href` getter.
 */
class DockerURL extends URL {
  static ORIGIN = 'http://docker';

  constructor(path: string) {
    super(DockerURL.ORIGIN + path);
  }

  get href() {
    return super.href.replace(DockerURL.ORIGIN, '');
  }
}

export const docker = {
  getSwarmInfo(): Promise<Docker.GetSwarmInspectResponse200> {
    const url = new DockerURL(`/api/docker/swarm`);
    return $fetch(url.href);
  },

  getSystemInfo(): Promise<Docker.GetSystemInfoResponse200> {
    const url = new DockerURL(`/api/docker/info`);
    return $fetch(url.href);
  },

  async getService(id: string): Promise<Docker.Service> {
    const res = await docker.listServices({
      filters: { id: [id] },
      status: true,
    });
    return res[0];
  },

  async listServices(options?: ListServiceOptions): Promise<Docker.GetServiceListResponse200> {
    let url = new DockerURL(`/api/docker/services`);

    if (options?.filters != null) {
      url.searchParams.set('filters', JSON.stringify(options.filters));
    }
    if (options?.status != null) {
      url.searchParams.set('status', String(options.status));
    }

    return await $fetch(url.href);
  },

  createService(spec: Docker.ServiceSpec): Promise<Docker.PostServiceCreateResponse201> {
    const url = new DockerURL(`/api/docker/services/create`);
    return $fetch(url.href, {
      method: 'POST',
      body: spec,
    });
  },

  async deleteService(service: Pick<Docker.Service, 'ID'>): Promise<void> {
    const url = new DockerURL(`/api/docker/services/${service.ID}`);
    await $fetch(url.href, {
      method: 'DELETE',
    });
  },

  async updateService(
    service: Docker.Service,
    newSpec: Docker.ServiceSpec,
  ): Promise<Docker.PostServiceUpdateResponse200> {
    const version = service.Version?.Index ?? 0;

    return $fetch(`/api/docker/services/${service.ID}/update?version=${version}`, {
      method: 'POST',
      body: newSpec,
    });
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
    const url = new DockerURL(`/api/docker/nodes`);

    if (options?.filters) {
      url.searchParams.set('filters', JSON.stringify(options.filters));
    }

    return $fetch(url.href);
  },

  listTasks(options?: ListTaskOptions): Promise<Docker.GetTaskListResponse200> {
    const url = new DockerURL(`/api/docker/tasks`);

    if (options?.filters) {
      url.searchParams.set('filters', JSON.stringify(options.filters));
    }

    return $fetch(url.href);
  },

  async pullImage(image: string): Promise<void> {
    const url = new DockerURL(`/api/docker/images/create`);
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
