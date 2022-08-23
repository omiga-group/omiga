import graphql from 'babel-plugin-relay/macro';
import { useCallback } from 'react';
import { useMutation } from 'react-relay';
import { PayloadError } from 'relay-runtime';
import cuid from 'cuid';

import { ClearProjectIsReviewedInput, ClearProjectIsReviewedMutation$data } from './__generated__/ClearProjectIsReviewedMutation.graphql';

const mutation = graphql`
  mutation ClearProjectIsReviewedMutation($input: ClearProjectIsReviewedInput!) @raw_response_type {
    clearProjectIsReviewed(input: $input) {
      clientMutationId
      project {
        id
        name
        isReviewed
      }
    }
  }
`;

export interface MutationReuqest {
  input: ClearProjectIsReviewedInput;
  name: string;
}

export interface Callback {
  onSuccess?: (mutationResponse: ClearProjectIsReviewedMutation$data) => void;
  onError?: (error: Error) => void;
}

const useClearProjectIsReviewedMutation = () => {
  const [commit] = useMutation(mutation);

  return useCallback(
    (request: MutationReuqest, callback?: Callback) => {
      commit({
        variables: {
          input: request.input,
        },
        optimisticResponse: {
          clearProjectIsReviewed: {
            clientMutationId: request.input.clientMutationId ?? cuid(),
            project: {
              id: request.input.id,
              name: request.name,
              isReviewed: false,
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

          callback.onSuccess(response.clearProjectIsReviewed);
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

export default useClearProjectIsReviewedMutation;
