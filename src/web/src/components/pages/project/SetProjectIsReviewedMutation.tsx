import graphql from 'babel-plugin-relay/macro';
import { useCallback } from 'react';
import { useMutation } from 'react-relay';
import { PayloadError } from 'relay-runtime';
import cuid from 'cuid';

import { SetProjectIsReviewedInput, SetProjectIsReviewedMutation$data } from './__generated__/SetProjectIsReviewedMutation.graphql';

const mutation = graphql`
  mutation SetProjectIsReviewedMutation($input: SetProjectIsReviewedInput!) @raw_response_type {
    setProjectIsReviewed(input: $input) {
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
  input: SetProjectIsReviewedInput;
  name: string;
}

export interface Callback {
  onSuccess?: (mutationResponse: SetProjectIsReviewedMutation$data) => void;
  onError?: (error: Error) => void;
}

const useSetProjectIsReviewedMutation = () => {
  const [commit] = useMutation(mutation);

  return useCallback(
    (request: MutationReuqest, callback?: Callback) => {
      commit({
        variables: {
          input: request.input,
        },
        optimisticResponse: {
          setProjectIsReviewed: {
            clientMutationId: request.input.clientMutationId ?? cuid(),
            project: {
              id: request.input.id,
              name: request.name,
              isReviewed: true,
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

          callback.onSuccess(response.setProjectIsReviewed);
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

export default useSetProjectIsReviewedMutation;
