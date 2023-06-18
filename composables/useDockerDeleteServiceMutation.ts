import { useMutation } from 'vue-query';

export default function () {
  return useMutation<
    void,
    H3Error<
      | Docker.DeleteServiceDeleteResponse404
      | Docker.DeleteServiceDeleteResponse500
      | Docker.DeleteServiceDeleteResponse503
    >,
    Docker.Service
  >({
    mutationFn: docker.deleteService,
  });
}
