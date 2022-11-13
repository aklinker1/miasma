import { useMutation, useQueryClient } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';
import { clone } from '~~/utils/clone';

export default function (
  servicelyChanges: (
    service: Readonly<Docker.Service>,
    spec: Docker.ServiceSpec,
  ) => Docker.ServiceSpec,
) {
  const client = useQueryClient();

  return useMutation<
    Docker.PostServiceUpdateResponse200,
    H3Error<
      | Docker.PostServiceUpdateResponse400
      | Docker.PostServiceUpdateResponse404
      | Docker.PostServiceUpdateResponse500
      | Docker.PostServiceUpdateResponse503
    >,
    Docker.Service
  >({
    mutationFn(service) {
      const version = service.Version?.Index ?? 0;
      const newSpec: Docker.ServiceSpec = clone(service.Spec!);

      return $fetch(`/api/docker/services/${service.ID}/update?version=${version}`, {
        method: 'POST',
        body: servicelyChanges(service, newSpec),
      });
    },
    async onSuccess(_, service) {
      // All the services queries
      client.invalidateQueries(QueryKeys.Services);
      // Service by ID
      client.invalidateQueries([QueryKeys.Service, service.ID]);
    },
  });
}
