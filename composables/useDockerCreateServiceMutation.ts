import { useMutation } from 'vue-query';

export default function () {
  return useMutation<
    Docker.PostServiceCreateResponse201,
    H3Error<
      | Docker.PostServiceCreateResponse400
      | Docker.PostServiceCreateResponse403
      | Docker.PostServiceCreateResponse409
      | Docker.PostServiceCreateResponse500
      | Docker.PostServiceCreateResponse503
    >,
    Docker.ServiceSpec
  >({
    mutationFn(spec) {
      return $fetch('/api/docker/services/create', {
        method: 'POST',
        body: spec,
      });
    },
  });
}
