import React, { useState } from 'react';
import graphql from 'babel-plugin-relay/macro';
import { createFragmentContainer } from 'react-relay';
import { QueryRenderer } from 'react-relay';
import { useParams } from 'react-router-dom';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableCell from '@mui/material/TableCell';
import Checkbox from '@mui/material/Checkbox';
import { useTranslation } from 'react-i18next';
import { Map } from 'immutable';
import { useSnackbar } from 'notistack';

import { ProjectDetails_projectRepository$data } from './__generated__/ProjectDetails_projectRepository.graphql';
import { ProjectDetails_project$data } from './__generated__/ProjectDetails_project.graphql';
import { ProjectDetails_Query$data } from './__generated__/ProjectDetails_Query.graphql';
import { ProjectDetailsQuery, ProjectDetailsQuery$data } from './__generated__/ProjectDetailsQuery.graphql';

import { default as createEnvironment } from '../../../framework/relay/Environment';
import LoadingContainer from '../../common/loading/LoadingContainer';
import GenericErrorContainer from '../../common/generic-error/GenericErrorContainer';
import { default as useStartMonitoringProjectRepositoryMutation } from './StartMonitoringProjectRepositoryMutation';
import { default as useStopMonitoringProjectRepositoryMutation } from './StopMonitoringProjectRepositoryMutation';

export const enNZTranslation = {
  code: 'Code',
  name: 'Name',
  monitored: 'Monitored?',
  gitCloneUrl: 'Git Clone Url',
};

const rootQuery = graphql`
  query ProjectDetailsQuery($code: String!, $pageSize: Int) {
    ...ProjectDetails_Query
  }
`;

const Header = React.memo(() => {
  const { t } = useTranslation();

  return (
    <TableHead>
      <TableRow>
        <TableCell>{t('projectDetailsTable.monitored')}</TableCell>
        <TableCell>{t('projectDetailsTable.gitCloneUrl')}</TableCell>
      </TableRow>
    </TableHead>
  );
});

interface ProjectRepositoryRowProps {
  projectRepository: ProjectDetails_projectRepository$data;
  isMonitored: boolean;
  onProjectRepositoryMonitorClick: (id: string) => void;
}

const ProjectRepositoryRow = React.memo<ProjectRepositoryRowProps>(
  ({ projectRepository: { id, gitCloneUrl }, isMonitored, onProjectRepositoryMonitorClick }) => {
    return (
      <TableRow>
        <TableCell padding="checkbox">
          <Checkbox checked={isMonitored} onClick={() => onProjectRepositoryMonitorClick(id)} />
        </TableCell>
        <TableCell>
          <a href={gitCloneUrl}>{gitCloneUrl}</a>
        </TableCell>
      </TableRow>
    );
  },
);

const ProjectRepositoryRowRelayed = createFragmentContainer(ProjectRepositoryRow, {
  projectRepository: graphql`
    fragment ProjectDetails_projectRepository on ProjectRepository {
      id
      gitCloneUrl
    }
  `,
});

interface ProjectRepositoriesTableProps {
  project: ProjectDetails_project$data;
  isMonitoredStates: Map<string, boolean>;
  onProjectRepositoryMonitorClick: (id: string) => void;
}

const ProjectRepositoriesTable = React.memo<ProjectRepositoriesTableProps>(({ project, isMonitoredStates, onProjectRepositoryMonitorClick }) => {
  const getProjectRepositoriesTable = (project: ProjectDetails_project$data) => {
    // @ts-ignore: Object is possibly 'null'.
    return project.projectRepositories.edges.map((edge) => (
      <ProjectRepositoryRowRelayed
        key={edge?.node?.id}
        // @ts-ignore: Object is possibly 'null'.
        projectRepository={edge?.node}
        onProjectRepositoryMonitorClick={onProjectRepositoryMonitorClick}
        // @ts-ignore: Object is possibly 'null'.
        isMonitored={isMonitoredStates.get(edge?.node?.id)}
      />
    ));
  };

  return (
    <Paper>
      <Table size="small">
        <Header />
        <TableBody>{getProjectRepositoriesTable(project)}</TableBody>
      </Table>
    </Paper>
  );
});

const ProjectRepositoriesTableRelayed = createFragmentContainer(ProjectRepositoriesTable, {
  project: graphql`
    fragment ProjectDetails_project on Project {
      projectRepositories(first: $pageSize, orderBy: [{ direction: ASC, field: gitCloneUrl }]) @connection(key: "Project_projectRepositories") {
        pageInfo {
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
        edges {
          cursor
          node {
            id
            ...ProjectDetails_projectRepository
          }
        }
      }
    }
  `,
});

interface ProjectProps {
  response: ProjectDetails_Query$data;
}

const Project = React.memo<ProjectProps>(({ response }) => {
  const commitStartMonitoringProjectRepositoryMutation = useStartMonitoringProjectRepositoryMutation();
  const commitStopMonitoringProjectRepositoryMutation = useStopMonitoringProjectRepositoryMutation();
  const { enqueueSnackbar } = useSnackbar();
  const [isMonitoredStates, setIsMonitoredStates] = useState(
    // @ts-ignore: Object is possibly 'null'.
    response.project.projectRepositories?.edges?.reduce(
      // @ts-ignore: Object is possibly 'null'.
      (reducation, val) => reducation.set(val?.node?.id, val?.node?.isMonitored),
      Map<string, boolean>(),
    ),
  );

  if (!response.project) {
    return <GenericErrorContainer message="project not found!!!" />;
  }

  const handleProjectRepositoryMonitorClick = (id: string) => {
    // @ts-ignore: Object is possibly 'null'.
    const projectRepositoryNode = response.project.projectRepositories?.edges?.find((edge) => edge?.node?.id === id);

    if (!projectRepositoryNode) {
      return;
    }

    const projectRepository = projectRepositoryNode.node;

    if (!projectRepository) {
      return;
    }

    const isMonitoredState = isMonitoredStates?.get(id);

    if (isMonitoredState === true) {
      commitStopMonitoringProjectRepositoryMutation(
        {
          gitCloneUrl: projectRepository.gitCloneUrl,
          input: {
            id,
          },
        },
        {
          onSuccess: (mutationResponse) => {
            // @ts-ignore: Object is possibly 'null'.
            setIsMonitoredStates(
              isMonitoredStates?.set(
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.projectRepository?.id,
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.projectRepository?.isMonitored,
              ),
            );

            enqueueSnackbar(`Stopped monitoring ${projectRepository.gitCloneUrl}`, { variant: 'success' });
          },
          onError: (error) => {
            enqueueSnackbar(`Failed to stop monitoring ${projectRepository.gitCloneUrl}. Error: ${error.message}`, { variant: 'error' });
          },
        },
      );
    } else {
      commitStartMonitoringProjectRepositoryMutation(
        {
          gitCloneUrl: projectRepository.gitCloneUrl,
          input: {
            id,
          },
        },
        {
          onSuccess: (mutationResponse) => {
            setIsMonitoredStates(
              isMonitoredStates?.set(
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.projectRepository?.id,
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.projectRepository?.isMonitored,
              ),
            );

            enqueueSnackbar(`Started monitoring ${projectRepository.gitCloneUrl}`, { variant: 'success' });
          },
          onError: (error) => {
            enqueueSnackbar(`Failed to start monitoring ${projectRepository.gitCloneUrl}. Error: ${error.message}`, { variant: 'error' });
          },
        },
      );
    }
  };

  return (
    <React.Fragment>
      <ProjectRepositoriesTableRelayed
        project={response.project}
        // @ts-ignore: Object is possibly 'null'.
        isMonitoredStates={isMonitoredStates}
        onProjectRepositoryMonitorClick={handleProjectRepositoryMonitorClick}
      />
    </React.Fragment>
  );
});

const ProjectRelayed = createFragmentContainer(Project, {
  response: graphql`
    fragment ProjectDetails_Query on Query {
      project(code: $code, activityPeriod: LastWeek) {
        name
        projectRepositories(first: $pageSize, orderBy: [{ direction: ASC, field: gitCloneUrl }]) @connection(key: "Project_projectRepositories") {
          edges {
            cursor
            node {
              id
              gitCloneUrl
              isMonitored
            }
          }
        }
        ...ProjectDetails_project
      }
    }
  `,
});

interface ProjectDetailsContainerProps {
  response: ProjectDetailsQuery$data;
}

const ProjectDetailsContainer = React.memo<ProjectDetailsContainerProps>(({ response }) => {
  return (
    <React.Fragment>
      <ProjectRelayed response={response} />
    </React.Fragment>
  );
});

export default React.memo(() => {
  const { appToken, code } = useParams();

  if (!appToken) {
    return <GenericErrorContainer message="No app token provided!!!" />;
  }

  if (!code) {
    return <GenericErrorContainer message="No code provided!!!" />;
  }

  const relay = {
    environment: createEnvironment(appToken),
  };

  return (
    <QueryRenderer<ProjectDetailsQuery>
      environment={relay.environment}
      query={rootQuery}
      variables={{
        code: code,
      }}
      render={({ props, error }) => {
        if (error) {
          return <GenericErrorContainer message={error.message} />;
        } else if (props) {
          return <ProjectDetailsContainer response={props} />;
        }

        return <LoadingContainer />;
      }}
    />
  );
});
