import { useQuery } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';

export default function () {
  return useQuery<Docker.GetSystemInfoResponse200, H3Error<Docker.GetSystemInfoResponse500>>({
    queryKey: [QueryKeys.SystemInfo],
    queryFn: () => $fetch('/api/docker/info'),
  });
}
