import { MaybeRef, Ref } from 'vue';
import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export default function (id: MaybeRef<string>) {
  return useQuery<
    Docker.Service,
    H3Error<Docker.GetServiceListResponse500 | Docker.GetServiceListResponse503>
  >({
    queryKey: [QueryKeys.Service, id],
    queryFn() {
      return docker.getService(unref(id));
    },
    staleTime: 1e3,
    refetchInterval: 1e3,
  });
}
