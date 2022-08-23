import graphql from 'babel-plugin-relay/macro';
import { useCallback } from 'react';
import { useMutation } from 'react-relay';
import { PayloadError } from 'relay-runtime';
import cuid from 'cuid';

import {
  StopMonitoringProjectRepositoryInput,
  StopMonitoringProjectRepositoryMutation$data,
} from './__generated__/StopMonitoringProjectRepositoryMutation.graphql';

const mutation = graphql`
  mutation StopMonitoringProjectRepositoryMutation($input: StopMonitoringProjectRepositoryInput!) @raw_response_type {
    stopMonitoringProjectRepository(input: $input) {
      clientMutationId
      projectRepository {
        id
        gitCloneUrl
        isMonitored
      }
    }
  }
`;

export interface MutationReuqest {
  input: StopMonitoringProjectRepositoryInput;
  gitCloneUrl: string;
}

export interface Callback {
  onSuccess?: (mutationResponse: StopMonitoringProjectRepositoryMutation$data) => void;
  onError?: (error: Error) => void;
}

const useStopMonitoringProjectRepositoryMutation = () => {
  const [commit] = useMutation(mutation);

  return useCallback(
    (request: MutationReuqest, callback?: Callback) => {
      commit({
        variables: {
          input: request.input,
        },
        optimisticResponse: {
          stopMonitoringProjectRepository: {
            clientMutationId: request.input.clientMutationId ?? cuid(),
            projectRepository: {
              id: request.input.id,
              gitCloneUrl: request.gitCloneUrl,
              isMonitored: false,
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

          callback.onSuccess(response.stopMonitoringProjectRepository);
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

export default useStopMonitoringProjectRepositoryMutation;
