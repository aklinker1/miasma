import { Ref } from 'vue';
import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export default function (id: Ref<string>) {
  return useQuery<
    Docker.Service,
    H3Error<Docker.GetServiceListResponse500 | Docker.GetServiceListResponse503>
  >({
    queryKey: [QueryKeys.Service, id],
    async queryFn() {
      const url =
        '/api/docker/services?status=true&filters=' +
        encodeURIComponent(JSON.stringify({ id: [id.value] }));
      const res: Docker.GetServiceListResponse200 = await $fetch(url);
      return res[0];
    },
    staleTime: 1e3,
    refetchInterval: 1e3,
  });
}
