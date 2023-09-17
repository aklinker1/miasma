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
      PublishedPort: 80,
      TargetPort: 80,
    },
    {
      PublishedPort: 8080,
      TargetPort: 8080,
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
