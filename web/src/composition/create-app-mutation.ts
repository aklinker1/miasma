import { gql } from "graphql-tag";
import { appFragment } from "../utils/apollo-client";
import { useMutation } from "@vue/apollo-composable";
import { App } from "./list-apps-query";

export interface AppInput {
  name: string;
  image: string;
  hidden?: boolean;
}

export function useCreateAppMutation() {
  return useMutation<{ app: App }, { app: AppInput }>(
    gql`
      mutation createApp($app: AppInput!) {
        app: createApp(input: $app) {
          ...AppListApp
        }
      }
      ${appFragment}
    `
  );
}
