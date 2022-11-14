import { useMutation } from 'vue-query';

export default function () {
  const { mutateAsync: copyService } = useDockerCreateServiceMutation();
  const { mutateAsync: deleteService } = useDockerDeleteServiceMutation();

  return useMutation<
    Docker.PostServiceCreateResponse201,
    H3Error<
      // Create errors
      | Docker.PostServiceCreateResponse400
      | Docker.PostServiceCreateResponse403
      | Docker.PostServiceCreateResponse409
      | Docker.PostServiceCreateResponse500
      | Docker.PostServiceCreateResponse503
      // Delete errors
      | Docker.DeleteServiceDeleteResponse404
      | Docker.DeleteServiceDeleteResponse500
      | Docker.DeleteServiceDeleteResponse503
    >,
    {
      prevService: Docker.Service;
      newSpec: Docker.ServiceSpec;
    }
  >({
    async mutationFn({ newSpec, prevService }) {
      const res: Docker.PostServiceCreateResponse201 = await copyService(newSpec);
      await deleteService(prevService).catch(console.error);
      return res;
    },
  });
}
