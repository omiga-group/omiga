import { Environment, Network, RecordSource, Store, RequestParameters, Variables } from 'relay-runtime';

// Create a network layer from the fetch function

const createEnvironment = () => {
  return new Environment({
    network: Network.create(async (request: RequestParameters, variables: Variables) => {
      const requestHeaders: HeadersInit = new Headers();

      requestHeaders.set('Accept', 'application/json');
      requestHeaders.set('Content-Type', 'application/json');

      const response = await fetch(window._env_.GRAPHQL_API_ENDPOINT, {
        method: 'POST',
        headers: requestHeaders,
        body: JSON.stringify({
          query: request.text,
          variables,
        }),
      });

      if (response.status !== 200) {
        throw new Error(response.statusText);
      }

      const result = await response.json();

      if (result.errors && result.errors.length > 0) {
        throw new Error(
          result.errors.map((error: Error) => error.message).reduce((reduction: string, message: string) => `${reduction}\n${message}`),
        );
      }

      return result;
    }),
    store: new Store(new RecordSource()),
  });
};

export default createEnvironment;
