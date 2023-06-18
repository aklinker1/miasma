import { useMutation, useQueryClient } from 'vue-query';

export default function () {
  const client = useQueryClient();

  return useMutation<
    void,
    H3Error<
      | Docker.PostNodeUpdateResponse400
      | Docker.PostNodeUpdateResponse404
      | Docker.PostNodeUpdateResponse500
      | Docker.PostNodeUpdateResponse503
    >,
    { node: Docker.Node; newSpec: Docker.NodeSpec }
  >({
    mutationFn({ node, newSpec }) {
      return docker.updateNode(node, newSpec);
    },
    async onSuccess(_, { node }) {
      // All the nodes queries
      await client.invalidateQueries(QueryKeys.Nodes);
      // Node by ID
      await client.invalidateQueries([QueryKeys.Nodes, node.ID]);
    },
  });
}
