import graphql from 'babel-plugin-relay/macro';
import { useCallback } from 'react';
import { useMutation } from 'react-relay';
import { PayloadError } from 'relay-runtime';
import cuid from 'cuid';

import { SubmitOrderInput, SubmitOrderMutation$data } from './__generated__/SubmitOrderMutation.graphql';

const mutation = graphql`
  mutation SubmitOrderMutation($input: SubmitOrderInput!) @raw_response_type {
    submitOrder(input: $input) {
      clientMutationId
      order {
        id
      }
    }
  }
`;

export interface MutationReuqest {
  input: SubmitOrderInput;
}

export interface Callback {
  onSuccess?: (mutationResponse: SubmitOrderMutation$data) => void;
  onError?: (error: Error) => void;
}

const useSubmitOrderMutation = () => {
  const [commit] = useMutation(mutation);

  return useCallback(
    (request: MutationReuqest, callback?: Callback) => {
      commit({
        variables: {
          input: request.input,
        },
        optimisticResponse: {
          submitOrder: {
            clientMutationId: request.input.clientMutationId ?? cuid(),
            order: {
              id: 0,
            },
          },
        },
        onCompleted: (response: any, errors: PayloadError[] | null) => {
          if (errors && errors.length > 0) {
            return;
          }

          if (!callback || !callback.onSuccess) {
            return;
          }

          callback.onSuccess(response.submitOrder);
        },
        onError: (error: Error) => {
          if (!callback || !callback.onError) {
            return;
          }

          callback.onError(error);
        },
      });
    },
    [commit],
  );
};

export default useSubmitOrderMutation;
