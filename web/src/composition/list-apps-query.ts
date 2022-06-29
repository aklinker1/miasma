import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";
import { Ref } from "vue";
import { appFragment } from "../utils/apollo-client";

export interface App {
  id: string;
  name: string;
  group?: string;
  status: string;
  simpleRoute?: string;
  instances?: {
    running: number;
    total: number;
  };
}

interface ListAppArgs {
  showHidden: Ref<boolean | undefined>;
}

export function useAppListQuery(args: ListAppArgs) {
  return useQuery<{ apps: App[] }, ListAppArgs>(
    gql`
      query listApps($showHidden: Boolean) {
        apps: listApps(showHidden: $showHidden) {
          ...AppListApp
        }
      }

      ${appFragment}
    `,
    args,
    { pollInterval: 2e3 }
  );
}
