import { Ref } from 'vue';

export default function (service: Ref<Docker.Service | undefined>) {
  const { data } = useDockerSystemInfoQuery();
  return computed<string[] | undefined>(() => {
    const publishedPorts = (service.value?.Spec?.EndpointSpec?.Ports ?? []).filter(
      p => p.PublishMode === 'ingress',
    );
    const ip = data.value?.Swarm?.NodeAddr;
    if (!publishedPorts.length || !ip) return;

    const portLocations = publishedPorts.map(
      port => `http://${ip}:${port.PublishedPort ?? port.TargetPort}`,
    );
    const ingressLocations: string[] = [];
    return [...portLocations, ...ingressLocations];
  });
}
