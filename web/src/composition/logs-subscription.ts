import { gql } from "graphql-tag";
import { useSubscription } from "@vue/apollo-composable";
import { Ref } from "vue";

export interface AppLog {
  timestamp: String;
  message: String;
}

interface LogArgs {
  appId: string;
}

export function useLogsSubscription(args: Ref<LogArgs>) {
  return useSubscription<{ log: AppLog }, LogArgs>(
    gql`
      subscription logs($appId: ID!) {
        log: appLogs(id: $appId, excludeStderr: false, excludeStdout: false) {
          timestamp
          message
        }
      }
    `,
    args
  );
}
