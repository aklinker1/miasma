import { MaybeRef } from 'vue';
import { useQuery } from 'vue-query';
import { ListNodeOptions } from './useDocker';

export default function (options?: MaybeRef<ListNodeOptions>) {
  const docker = useDocker();

  return useDockerQuery<
    Docker.GetNodeListResponse200,
    H3Error<Docker.GetNodeListResponse500 | Docker.GetNodeListResponse503>
  >({
    queryKey: [QueryKeys.Nodes, options],
    queryFn() {
      return docker.listNodes(unref(options));
    },
  });
}
