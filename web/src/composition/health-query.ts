import { gql } from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";
import type { OptionsParameter } from "@vue/apollo-composable/dist/useQuery";

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

export function useHealthQuery(
  options?: OptionsParameter<{ health: Health }, {}>
) {
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
    {},
    options ?? {}
  );
}
