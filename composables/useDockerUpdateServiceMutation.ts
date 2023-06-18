import { useMutation, useQueryClient } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';
import { routes } from '~/utils/routes';

export default function () {
  const client = useQueryClient();
  const router = useRouter();

  return useMutation<
    Docker.PostServiceUpdateResponse200 | Docker.PostServiceCreateResponse201,
    H3Error<
      | Docker.PostServiceUpdateResponse400
      | Docker.PostServiceUpdateResponse404
      | Docker.PostServiceUpdateResponse500
      | Docker.PostServiceUpdateResponse503
    >,
    { service: Docker.Service; newSpec: Docker.ServiceSpec }
  >({
    async mutationFn({ service, newSpec }) {
      if (service.Spec?.Name !== newSpec.Name) {
        return docker.renameService(service, newSpec);
      } else {
        return docker.updateService(service, newSpec);
      }
    },
    async onSuccess(res, { service }) {
      // All the services queries
      client.invalidateQueries(QueryKeys.Services);
      // Service by ID
      client.invalidateQueries([QueryKeys.Service, service.ID]);

      if ('ID' in res) {
        // If renamed, go to new service URL
        router.replace(routes.service(res.ID));
      }
    },
  });
}
