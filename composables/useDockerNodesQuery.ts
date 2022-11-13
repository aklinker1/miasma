import { Ref } from 'vue';
import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export type NodeFilters = {
  id?: string[];
  label?: string[];
  membership?: Array<'accepted' | 'pending'>;
  name?: string[];
  'node.label'?: string[];
  role?: Array<'manager' | 'worker'>;
};

export default function (filters?: Ref<NodeFilters>) {
  return useQuery<
    Docker.GetNodeListResponse200,
    H3Error<Docker.GetNodeListResponse500 | Docker.GetNodeListResponse503>
  >({
    queryKey: [QueryKeys.Nodes, filters],
    async queryFn() {
      let url = '/api/docker/nodes';
      if (filters?.value) {
        const filtersStr = encodeURIComponent(JSON.stringify(toRaw(filters.value)));
        url += `?filters=${filtersStr}`;
      }
      return $fetch(url);
    },
  });
}
