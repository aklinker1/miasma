import { useMutation, useQueryClient } from 'vue-query';

export default function () {
  const client = useQueryClient();
  const docker = useDocker();

  return useMutation<
    Docker.PostServiceUpdateResponse200 | Docker.PostServiceCreateResponse201,
    H3Error<
      | Docker.PostServiceUpdateResponse400
      | Docker.PostServiceUpdateResponse404
      | Docker.PostServiceUpdateResponse500
      | Docker.PostServiceUpdateResponse503
    >,
    Docker.Service
  >({
    mutationFn: docker.startService,
    async onSuccess(_, service) {
      // All the services queries
      client.invalidateQueries(QueryKeys.Services);
      // Service by ID
      client.invalidateQueries([QueryKeys.Service, service.ID]);
    },
  });
}
