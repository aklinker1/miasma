import { gql } from "graphql-tag";
import { appFragment } from "../utils/apollo-client";
import { useMutation } from "@vue/apollo-composable";
import { App } from "./list-apps-query";

export function useRestartAppMutation() {
  return useMutation<{ app: App }, { id: string }>(
    gql`
      mutation restartApp($id: ID!) {
        app: restartApp(id: $id) {
          ...AppListApp
        }
      }
      ${appFragment}
    `
  );
}
