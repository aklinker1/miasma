import { Ref } from "vue";
import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";

export interface AppDetails {
  id: string;
  name: string;
  image: string;
  group?: string;
  status: string;
  availableAt: string[];
  instances?: {
    running: number;
    total: number;
  };
}

export function useAppDetailsQuery(appId: Ref<string | undefined>) {
  const variables = computed(() => ({
    id: appId.value,
    clusterIpAddress: location.hostname,
  }));
  const options = computed(() => ({ enabled: !!appId.value }));

  const { result, ...query } = useQuery<
    { app: AppDetails },
    { id: string | undefined }
  >(
    gql`
      query appDetails($id: ID!, $clusterIpAddress: String!) {
        app: getApp(id: $id) {
          id
          name
          image
          group
          status
          availableAt(clusterIpAddress: $clusterIpAddress)
          instances {
            running
            total
          }
        }
      }
    `,
    variables,
    options
  );
  const app = computed<AppDetails | undefined>(() => {
    if (!options.value.enabled) return undefined;
    return result.value?.app;
  });

  return {
    ...query,
    app,
  };
}
