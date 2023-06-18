import { MaybeRef, Ref } from 'vue';
import { UseQueryOptions, useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

type Error = H3Error<Docker.GetServiceListResponse500 | Docker.GetServiceListResponse503>;
type Options = UseQueryOptions<Docker.Service, Error>;

export default function (id: MaybeRef<string>, options?: Options) {
  return useQuery<Docker.Service, Error>({
    queryKey: [QueryKeys.Service, id],
    queryFn() {
      return docker.getService(unref(id));
    },
    staleTime: 1e3,
    refetchInterval: 1e3,
    ...options,
  });
}
