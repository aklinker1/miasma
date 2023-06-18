import { useMutation, useQueryClient } from 'vue-query';
import { QueryKeys } from '~~/utils/QueryKeys';
import { routes } from '~/utils/routes';

export default function () {
  const client = useQueryClient();
  const router = useRouter();

  return useMutation<
    void,
    H3Error<
      | Docker.PostServiceUpdateResponse400
      | Docker.PostServiceUpdateResponse404
      | Docker.PostServiceUpdateResponse500
      | Docker.PostServiceUpdateResponse503
      | Docker.PostImageCreateResponse404
      | Docker.PostImageCreateResponse500
    >,
    Docker.Service
  >({
    async mutationFn(service) {
      await docker.pullLatest(service);
    },
    async onSuccess(res, service) {
      // All the services queries
      client.invalidateQueries(QueryKeys.Services);
      // Service by ID
      client.invalidateQueries([QueryKeys.Service, service.ID]);
    },
  });
}
