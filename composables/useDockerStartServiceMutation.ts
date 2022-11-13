import { MiasmaLabels } from '~~/utils/labels';

export default function () {
  return useDockerUpdateServiceMutation((service, spec) => {
    const desiredTasks = service.ServiceStatus?.DesiredTasks ?? 0;
    if (desiredTasks > 0) throw Error('Service is already running');

    const replicas = Number(service.Spec?.Labels?.[MiasmaLabels.InstanceCount] ?? '1');
    spec.Mode!.Replicated!.Replicas = replicas;

    return spec;
  });
}
