import graphql from 'babel-plugin-relay/macro';
import { useCallback } from 'react';
import { useMutation } from 'react-relay';
import { PayloadError } from 'relay-runtime';
import cuid from 'cuid';

import {
  StartMonitoringProjectRepositoryInput,
  StartMonitoringProjectRepositoryMutation$data,
} from './__generated__/StartMonitoringProjectRepositoryMutation.graphql';

const mutation = graphql`
  mutation StartMonitoringProjectRepositoryMutation($input: StartMonitoringProjectRepositoryInput!) @raw_response_type {
    startMonitoringProjectRepository(input: $input) {
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
  input: StartMonitoringProjectRepositoryInput;
  gitCloneUrl: string;
}

export interface Callback {
  onSuccess?: (mutationResponse: StartMonitoringProjectRepositoryMutation$data) => void;
  onError?: (error: Error) => void;
}

const useStartMonitoringProjectRepositoryMutation = () => {
  const [commit] = useMutation(mutation);

  return useCallback(
    (request: MutationReuqest, callback?: Callback) => {
      commit({
        variables: {
          input: request.input,
        },
        optimisticResponse: {
          startMonitoringProjectRepository: {
            clientMutationId: request.input.clientMutationId ?? cuid(),
            projectRepository: {
              id: request.input.id,
              gitCloneUrl: request.gitCloneUrl,
              isMonitored: true,
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

          callback.onSuccess(response.startMonitoringProjectRepository);
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

export default useStartMonitoringProjectRepositoryMutation;
