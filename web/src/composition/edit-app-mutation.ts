import { gql } from "graphql-tag";
import { useMutation } from "@vue/apollo-composable";

export interface AppChanges {
  name?: string | null;
  image?: string | null;
  group?: string | null;
}

export function useEditAppMutation() {
  return useMutation<
    { app: { id: string } },
    { id: string; changes: AppChanges }
  >(
    gql`
      mutation editApp($id: ID!, $changes: AppChanges!) {
        app: editApp(id: $id, changes: $changes) {
          id
        }
      }
    `,
    {
      refetchQueries: ["listApps", "appDetails"],
    }
  );
}
