import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";

export interface Node {
  id: string;
  os: string;
  architecture: string;
  status: string;
  statusMessage?: string;
  hostname: string;
  labels: Record<string, string>;
  ip: string;
  services: Array<{
    name: string;
  }>;
}

export function useNodesQuery() {
  return useQuery<{ nodes: Node[] }, {}>(
    gql`
      query nodes {
        nodes {
          id
          os
          architecture
          status
          statusMessage
          hostname
          labels
          ip
          services {
            name
          }
        }
      }
    `,
    { pollInterval: 5e3 }
  );
}
