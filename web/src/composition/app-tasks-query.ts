import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";
import { Ref } from "vue";

export interface AppDetailsTask {
  state: string;
  desiredState: string;
  timestamp: string;
  message: string;
  error?: string;
  exitCode?: string;
}

export interface AppTasksQueryVars {
  appId: string;
}

export function useAppTasksQuery(vars: Ref<AppTasksQueryVars>) {
  return useQuery<{ tasks: AppDetailsTask[] }, { appId: string }>(
    gql`
      query tasks($appId: ID!) {
        tasks: getAppTasks(id: $appId) {
          state
          desiredState
          message
          timestamp
          error
          exitCode
        }
      }
    `,
    vars,
    { pollInterval: 5e3 }
  );
}
