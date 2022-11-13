import { Ref } from 'vue';
import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export type ServiceFilters = {
  id?: string[];
  label?: string[];
  mode?: Array<'replicated' | 'global'>;
  name?: string[];
};

export default function (status?: Ref<boolean>, filters?: Ref<ServiceFilters>) {
  return useQuery<
    Docker.GetServiceListResponse200,
    H3Error<Docker.GetServiceListResponse500 | Docker.GetServiceListResponse503>
  >({
    queryKey: [QueryKeys.Services, status, filters],
    queryFn() {
      let url = '/api/docker/services';
      let queryParams: string[] = [];
      if (filters != null)
        queryParams.push('filters=' + encodeURIComponent(JSON.stringify(toRaw(filters.value))));
      if (status != null) queryParams.push('status=' + encodeURIComponent(toRaw(status.value)));
      if (queryParams.length > 0) {
        url += `?${queryParams.join('&')}`;
      }
      return $fetch(url);
    },
    staleTime: 1e3,
    refetchInterval: 1e3,
  });
}
