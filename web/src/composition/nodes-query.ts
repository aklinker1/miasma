import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";
import { appFragment } from "../utils/apollo-client";

interface NodeService {
  id: string;
  name: string;
}

export interface Node {
  id: string;
  os: string;
  architecture: string;
  status: string;
  statusMessage?: string;
  hostname: string;
  labels: Record<string, string>;
  ip: string;
  services: NodeService[];
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
            ...AppListApp
          }
        }
      }

      ${appFragment}
    `,
    {},
    { pollInterval: 5e3 }
  );
}
