import { gql } from "graphql-tag";
import { appFragment } from "../utils/apollo-client";
import { useMutation } from "@vue/apollo-composable";
import { App } from "./list-apps-query";

export function useDeleteAppMutation() {
  return useMutation<{ app: App }, { id: string }>(
    gql`
      mutation deleteApp($id: ID!) {
        app: deleteApp(id: $id) {
          ...AppListApp
        }
      }
      ${appFragment}
    `
  );
}
