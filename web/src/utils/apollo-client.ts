import {
  ApolloClient,
  createHttpLink,
  gql,
  InMemoryCache,
} from "@apollo/client/core";
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

// Cache implementation
const cache = new InMemoryCache();

// Create the apollo client
export const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
});
