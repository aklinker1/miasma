import { Ref } from 'vue';
import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export type TaskFilters = {
  'desired-state'?: Array<'running' | 'shutdown' | 'accepted'>;
  id?: string[];
  /**
   * Each array itme should be in the form of "key=value".
   */
  label?: string[];
  /**
   * The task name.
   */
  name?: string[];
  /**
   * ID or name of the node.
   */
  node?: string[];
  service?: string[];
};

export default function (filters?: Ref<TaskFilters>) {
  return useQuery<
    Docker.GetTaskListResponse200,
    H3Error<Docker.GetTaskListResponse500 | Docker.GetTaskListResponse503>
  >({
    queryKey: [QueryKeys.Nodes, filters],
    queryFn() {
      let url = '/api/docker/tasks';
      if (filters?.value) {
        const filtersStr = encodeURIComponent(JSON.stringify(toRaw(filters.value)));
        url += `?filters=${filtersStr}`;
      }
      return $fetch(url);
    },
    staleTime: 1e3,
    cacheTime: 1e3,
    refetchInterval: 1e3,
  });
}
