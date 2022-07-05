import { gql } from "graphql-tag";
import { appFragment } from "../utils/apollo-client";
import { useMutation } from "@vue/apollo-composable";
import { App } from "./list-apps-query";

export function useStartAppMutation() {
  return useMutation<{ app: App }, { id: string }>(
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
