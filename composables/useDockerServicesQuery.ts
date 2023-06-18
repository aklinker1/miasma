import { MaybeRef } from 'vue';
import { useQuery } from 'vue-query';
import { ListServiceOptions } from '~/utils/docker';
import { QueryKeys } from '~~/utils/QueryKeys';

export default function (options?: MaybeRef<ListServiceOptions>) {
  return useQuery<
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
