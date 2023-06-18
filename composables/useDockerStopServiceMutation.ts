import { useMutation, useQueryClient } from 'vue-query';
import { MiasmaLabels } from '~~/utils/labels';

export default function () {
  const client = useQueryClient();

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
    mutationFn: docker.stopService,
    async onSuccess(_, service) {
      // All the services queries
      client.invalidateQueries(QueryKeys.Services);
      // Service by ID
      client.invalidateQueries([QueryKeys.Service, service.ID]);
    },
  });
}
