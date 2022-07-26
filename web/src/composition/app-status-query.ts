import type { OptionsParameter } from "@vue/apollo-composable/dist/useQuery";
import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";
import { Ref } from "vue";

export type RuntimeStatus = "RUNNING" | "STOPPED";

export interface AppStatus {
  status: RuntimeStatus;
  instances?: {
    total: number;
    running: number;
  };
}

export function useAppStatusQuery(appId: Ref<string>) {
  const variables = computed(() => ({
    id: appId.value,
  }));
  const { result } = useQuery<{ app: AppStatus }, { id: string }>(
    gql`
      query appStatus($id: ID!) {
        app: getApp(id: $id) {
          id
          status
          instances {
            total
            running
          }
        }
      }
    `,
    variables,
    { pollInterval: 1e3 }
  );

  return computed(() => result.value?.app);
}
