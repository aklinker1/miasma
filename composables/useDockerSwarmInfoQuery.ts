import { useQuery } from 'vue-query';

export default function () {
  const docker = useDocker();
  const auth = useAuthCookie();
  const hasAuth = computed(() => !!auth.value);

  return useDockerQuery<
    Docker.GetSwarmInspectResponse200,
    H3Error<
      | Docker.GetSwarmInspectResponse404
      | Docker.GetSwarmInspectResponse500
      | Docker.GetSwarmInspectResponse503
    >
  >({
    queryKey: [QueryKeys.SwarmInfo],
    queryFn: docker.getSwarmInfo,
    staleTime: 20e3,
    refetchInterval: 20e3,
    enabled: hasAuth,
  });
}
