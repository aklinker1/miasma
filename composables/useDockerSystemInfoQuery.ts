import { useQuery } from 'vue-query';

export default function () {
  const docker = useDocker();

  return useDockerQuery<Docker.GetSystemInfoResponse200, H3Error<Docker.GetSystemInfoResponse500>>({
    queryKey: [QueryKeys.SystemInfo],
    queryFn: docker.getSystemInfo,
  });
}
