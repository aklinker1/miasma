import { gql } from "graphql-tag";
import { useMutation } from "@vue/apollo-composable";
import { appFragment } from "../utils/apollo-client";
import { App } from "./list-apps-query";

export function useStopAppMutation() {
  return useMutation<App, { id: string }>(
    gql`
      mutation stopApp($id: ID!) {
        app: stopApp(id: $id) {
          ...AppListApp
        }
      }
      ${appFragment}
    `
  );
}
