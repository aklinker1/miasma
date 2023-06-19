import { MaybeRef } from 'vue';
import { useQuery } from 'vue-query';
import { ListServiceOptions } from './useDocker';

export default function (options?: MaybeRef<ListServiceOptions>) {
  const docker = useDocker();

  return useDockerQuery<
    Docker.GetServiceListResponse200,
    H3Error<Docker.GetServiceListResponse500 | Docker.GetServiceListResponse503>
  >({
    queryKey: [QueryKeys.Services, options],
    queryFn() {
      return docker.listServices(unref(options));
    },
    staleTime: 1e3,
    refetchInterval: 1e3,
  });
}
