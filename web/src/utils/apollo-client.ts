import {
  ApolloClient,
  createHttpLink,
  gql,
  split,
  InMemoryCache,
} from "@apollo/client/core";
import { WebSocketLink } from "@apollo/client/link/ws";
import { getMainDefinition } from "@apollo/client/utilities";
import { router } from "../router";

export const appFragment = gql`
  fragment AppListApp on App {
    id
    name
    group
    status
    simpleRoute
    instances {
      running
      total
    }
  }
`;

const ACCESS_TOKEN_STORAGE_KEY = "accessToken";
export function setAccessToken(token: string) {
  localStorage.setItem(ACCESS_TOKEN_STORAGE_KEY, token);
}
export function getAccessToken(): string | undefined {
  return localStorage.getItem(ACCESS_TOKEN_STORAGE_KEY) ?? undefined;
}

// HTTP connection to the API
const httpLink = createHttpLink({
  uri: __API_URL__,
  async fetch(input, init) {
    // Add the access token if necessary
    const token = getAccessToken();
    if (token) {
      // @ts-expect-error: Headers is setup by apollo always
      init.headers["Authorization"] = `Bearer ${token}`;
    }

    // Perform the request
    const res = await fetch(input, init);

    // Redirect to login on 401s
    if (res.status === 401) {
      if (!location.pathname.startsWith("/login")) {
        const url = new URL(`${location.origin}/login`);
        url.searchParams.append("redirect", window.location.href);
        router.push(url.href);
      }
    }
    return res;
  },
});

const wsLink = new WebSocketLink({
  uri: import.meta.env.DEV
    ? "ws://localhost:3000/graphql"
    : `ws://${location.host}/graphql`,
  options: {
    reconnect: true,

    connectionParams: {
      authToken: "",
    },
  },
});

// Cache implementation
const cache = new InMemoryCache();

// Create the apollo client
export const apolloClient = new ApolloClient({
  link: split(
    ({ query }) => {
      const definition = getMainDefinition(query);
      return (
        definition.kind === "OperationDefinition" &&
        definition.operation === "subscription"
      );
    },
    wsLink,
    httpLink
  ),
  cache,
});
