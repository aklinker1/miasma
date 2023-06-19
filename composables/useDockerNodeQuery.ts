import { MaybeRef } from 'vue';
import { UseQueryOptions, useQuery } from 'vue-query';

type Errors = H3Error<Docker.GetNodeListResponse500 | Docker.GetNodeListResponse503>;
type Options = UseQueryOptions<Docker.Node, Errors>;

export default function (id: MaybeRef<string>, options?: Options) {
  const docker = useDocker();

  return useDockerQuery<Docker.Node, Errors>({
    queryKey: [QueryKeys.Node, id],
    queryFn() {
      return docker.getNode(unref(id));
    },
    staleTime: 1e3,
    refetchInterval: 1e3,
    ...options,
  });
}
