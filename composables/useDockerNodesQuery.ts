import { MaybeRef, Ref } from 'vue';
import { useQuery } from 'vue-query';
import { ListNodeOptions } from '~/utils/docker';
import { QueryKeys } from '~~/utils/QueryKeys';

export default function (options?: MaybeRef<ListNodeOptions>) {
  return useQuery<
    Docker.GetNodeListResponse200,
    H3Error<Docker.GetNodeListResponse500 | Docker.GetNodeListResponse503>
  >({
    queryKey: [QueryKeys.Nodes, options],
    queryFn() {
      return docker.listNodes(unref(options));
    },
  });
}
