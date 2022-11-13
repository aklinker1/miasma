import { Ref } from 'vue';

/**
 * Whether or not the service is fully running
 */
export default function (
  service: Ref<Docker.GetServiceInspectResponse200 | undefined>,
): Ref<'running' | 'degraded' | 'stopped'> {
  return computed(() => {
    const desired = service.value?.ServiceStatus?.DesiredTasks ?? 0;
    if (desired === 0) return 'stopped';

    const running = service.value?.ServiceStatus?.RunningTasks ?? 0;
    if (running < desired) return 'degraded';
    return 'running';
  });
}
