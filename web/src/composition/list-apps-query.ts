import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";
import { Ref } from "vue";
import { appFragment } from "../utils/apollo-client";
import { RuntimeStatus } from "./app-status-query";

export interface App {
  id: string;
  name: string;
  group?: string;
  status: RuntimeStatus;
  simpleRoute?: string;
  instances?: {
    running: number;
    total: number;
  };
  availableAt: string[];
}

interface ListAppArgs {
  showHidden: Ref<boolean | undefined>;
  clusterIpAddress: String;
}

export function useAppListQuery(args: ListAppArgs) {
  return useQuery<{ apps: App[] }, ListAppArgs>(
    gql`
      query listApps($showHidden: Boolean, $clusterIpAddress: String!) {
        apps: listApps(showHidden: $showHidden) {
          ...AppListApp
          availableAt(clusterIpAddress: $clusterIpAddress)
        }
      }

      ${appFragment}
    `,
    args,
    { pollInterval: 2e3 }
  );
}
