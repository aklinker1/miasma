import { gql } from "graphql-tag";
import { appFragment } from "../utils/apollo-client";
import { useMutation } from "@vue/apollo-composable";

export function useStartAppMutation() {
  return useMutation(
    gql`
      mutation startApp($id: ID!) {
        app: startApp(id: $id) {
          ...AppListApp
        }
      }
      ${appFragment}
    `
  );
}
