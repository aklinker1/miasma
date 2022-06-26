import {
  ApolloClient,
  createHttpLink,
  gql,
  InMemoryCache,
} from "@apollo/client/core";

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

// HTTP connection to the API
const httpLink = createHttpLink({
  uri: __API_URL__,
});

// Cache implementation
const cache = new InMemoryCache();

// Create the apollo client
export const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
});
