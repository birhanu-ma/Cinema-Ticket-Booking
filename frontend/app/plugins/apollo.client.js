import {
  ApolloClient,
  InMemoryCache,
  HttpLink,
  ApolloLink,
} from "@apollo/client/core";

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig();

  const httpLink = new HttpLink({
    uri: config.public.graphqlEndpoint,
  });

  const authLink = new ApolloLink((operation, forward) => {
    operation.setContext(({ headers = {} }) => {
      let token = "";

      if (import.meta.client) {
        token = localStorage.getItem("auth-token") || "";
      }

      return {
        headers: {
          ...headers,
          ...(token && {
            Authorization: `Bearer ${token}`,
          }),
        },
      };
    });

    return forward(operation);
  });

  const apollo = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });

  return {
    provide: {
      apollo,
    },
  };
});
