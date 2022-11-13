export default function () {
  return useDockerUpdateServiceMutation((service, spec) => {
    const desiredTasks = service.ServiceStatus?.DesiredTasks ?? 0;
    if (desiredTasks === 0) throw Error('Service is already stopped');

    spec.Mode!.Replicated!.Replicas = 0;
    return spec;
  });
}
