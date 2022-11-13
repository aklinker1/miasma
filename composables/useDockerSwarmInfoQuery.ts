import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export default function () {
  return useQuery<
    Docker.GetSwarmInspectResponse200,
    H3Error<
      | Docker.GetSwarmInspectResponse404
      | Docker.GetSwarmInspectResponse500
      | Docker.GetSwarmInspectResponse503
    >
  >({
    queryKey: [QueryKeys.SwarmInfo],
    queryFn: () => $fetch('/api/docker/swarm'),
    staleTime: 20e3,
    refetchInterval: 20e3,
  });
}
