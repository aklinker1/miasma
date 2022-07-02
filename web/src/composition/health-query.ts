import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";

export interface Health {
  version: string;
  dockerVersion: string;
  cluster?: {
    id: string;
    joinCommand: string;
    createdAt: string;
    updatedAt: string;
  };
}

export function useHealthQuery() {
  return useQuery<{ health: Health }, {}>(
    gql`
      query getHealth {
        health {
          version
          dockerVersion
          cluster {
            id
            joinCommand
            createdAt
            updatedAt
          }
        }
      }
    `,
    {}
  );
}
