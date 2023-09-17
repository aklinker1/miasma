export function buildTraefikService(config: TraefikConfig): Docker.ServiceSpec {
  const command = ['traefik'];
  if (config.httpsEnabled) {
    if (!config.certsEmail) {
      throw Error('Https is enabled, but a certificate email was not provided');
    }
    if (!config.certsDir) {
      throw Error('Https is enabled, but no directory was provided to store your certs');
    }

    command.push(
      '--entrypoints.web.address=:80',
      '--entrypoints.websecure.address=:443',
      // Use LetsEncrypt to manage certs: https://doc.traefik.io/traefik/https/acme/#configuration-examples
      `--certificatesresolvers.${TRAEFIK_CERTS_RESOLVER_NAME}.acme.email=${config.certsEmail}`,
      `--certificatesresolvers.${TRAEFIK_CERTS_RESOLVER_NAME}.acme.storage=/letsencrypt/acme.json`,
      `--certificatesresolvers.${TRAEFIK_CERTS_RESOLVER_NAME}.acme.httpchallenge.entrypoint=web`,
      // Redirect HTTP -> HTTPS: https://doc.traefik.io/traefik/routing/entrypoints/#redirection
      '--entrypoints.web.http.redirections.entrypoint.to=websecure',
      '--entrypoints.web.http.redirections.entrypoint.scheme=https',
    );
  }
  command.push('--api.insecure=true', '--providers.docker', '--providers.docker.swarmmode');

  const ports: Docker.EndpointPortConfig[] = [
    {
      PublishedPort: 8080,
      TargetPort: 8080,
    },
    {
      PublishedPort: 80,
      TargetPort: 80,
    },
  ];
  if (config.httpsEnabled) {
    ports.push({
      PublishedPort: 443,
      TargetPort: 443,
    });
  }

  const volumes: Docker.Mount[] = [
    {
      Source: '/var/run/docker.sock',
      Target: '/var/run/docker.sock',
    },
  ];
  if (config.httpsEnabled) {
    volumes.push({
      Source: config.certsDir,
      Target: '/letsencrypt',
    });
  }

  const labels = {
    [MiasmaLabels.System]: '',
    [MiasmaLabels.Group]: 'Plugins',
  };

  const containerLabels = {
    [MiasmaLabels.System]: '',
  };

  return {
    Name: 'miasma-traefik-plugin',
    EndpointSpec: {
      Ports: ports,
    },
    TaskTemplate: {
      ContainerSpec: {
        Image: 'traefik:2.7',
        Command: command,
        Mounts: volumes,
        Labels: containerLabels,
      },
      Placement: {
        Constraints: ['node.role == manager'],
      },
    },
    Labels: labels,
  };
}

export const TRAEFIK_SERVICE_NAME = 'miasma-traefik-plugin';
export const TRAEFIK_CERTS_RESOLVER_NAME = 'miasmaresolver';

export interface TraefikConfig {
  httpsEnabled?: boolean;
  /**
   * Email used with let's encrypt
   */
  certsEmail?: string;
  /**
   * Server directory let's encyrpt will store certs in.
   *
   * Ex: "/root/letsencrypt"
   */
  certsDir?: string;
}

export function userInputToTraefikRule(input: string): string {
  if (input === '') return input;

  if (input.includes('(')) return input;
  if (input.includes('/')) {
    const slashIndex = input.indexOf('/');
    const hostname = input.substring(0, slashIndex);
    const pathPrefix = input.substring(slashIndex + 1);
    return `Host(\`${hostname}\`) && PathPrefix(\`/${pathPrefix}\`)`;
  }
  return `Host(\`${input}\`)`;
}

export function traefikRuleToUserInput(rule: string): string {
  let match;
  if ((match = rule.match(/Host\(`(.+?)`\)\s*&&\s?PathPrefix\(`\/(.+?)`\)/)))
    return `${match[1]}/${match[2]}`;
  if ((match = rule.match(/Host\(`(.+?)`\)/))) return match[1];
  return rule;
}

export function applyTraefikLabels(
  labels: Record<string, string>,
  name: string,
  rule: string,
  ports: Docker.EndpointPortConfig[],
  config: TraefikConfig,
) {
  // Remove old labels
  removeTraefikLabels(labels);

  const hostname = name;
  const port = ports[0]?.TargetPort;

  labels[`traefik.enable`] = 'true';
  labels[`traefik.docker.network`] = 'ingress';
  if (port) {
    labels[`traefik.http.services.${hostname}.loadbalancer.server.port`] = String(port);
  }
  labels[`traefik.http.routers.${hostname}.rule`] = rule;

  if (config.httpsEnabled) {
    labels[`traefik.http.routers.${hostname}.tls`] = 'true';
    labels[`traefik.http.routers.${hostname}.tls.certresolver`] = TRAEFIK_CERTS_RESOLVER_NAME;
  }
}

export function removeTraefikLabels(labels: Record<string, string>) {
  Object.keys(labels).forEach(key => {
    if (key.startsWith('traefik.')) delete labels[key];
  });
}
