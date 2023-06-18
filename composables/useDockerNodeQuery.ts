import { MaybeRef, Ref } from 'vue';
import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export default function (id: MaybeRef<string>) {
  return useQuery<
    Docker.Node,
    H3Error<Docker.GetNodeListResponse500 | Docker.GetNodeListResponse503>
  >({
    queryKey: [QueryKeys.Node, id],
    queryFn() {
      return docker.getNode(unref(id));
    },
    staleTime: 1e3,
    refetchInterval: 1e3,
  });
}
