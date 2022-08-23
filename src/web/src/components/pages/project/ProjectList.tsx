import React, { useState } from 'react';
import graphql from 'babel-plugin-relay/macro';
import { createFragmentContainer, createPaginationContainer } from 'react-relay';
import { Environment } from 'relay-runtime';
import { QueryRenderer } from 'react-relay';
import { useParams } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableCell from '@mui/material/TableCell';
import Link from '@mui/material/Link';
import Button from '@mui/material/Button';
import Checkbox from '@mui/material/Checkbox';
import { useTranslation } from 'react-i18next';
import { Map } from 'immutable';
import { useSnackbar } from 'notistack';

import { ProjectList_project$data } from './__generated__/ProjectList_project.graphql';
import { ProjectList_Query$data } from './__generated__/ProjectList_Query.graphql';
import { ProjectListQuery, ProjectListQuery$data } from './__generated__/ProjectListQuery.graphql';

import { default as createEnvironment } from '../../../framework/relay/Environment';
import LoadingContainer from '../../common/loading/LoadingContainer';
import GenericErrorContainer from '../../common/generic-error/GenericErrorContainer';
import { default as useSetProjectIsReviewedMutation } from './SetProjectIsReviewedMutation';
import { default as useClearProjectIsReviewedMutation } from './ClearProjectIsReviewedMutation';

export const enNZTranslation = {
  loadMore: 'Load more',
  rank: 'Rank',
  code: 'Code',
  name: 'Name',
  link: 'Link',
  reviewed: 'Reviewed?',
};

const pageSize = 10000;

const rootQuery = graphql`
  query ProjectListQuery($count: Int!, $after: Cursor) {
    projects(first: $count, after: $after, activityPeriod: LastWeek, orderBy: [{ direction: ASC, field: rank }]) @connection(key: "Query_projects") {
      edges {
        cursor
        node {
          id
          name
          isReviewed
        }
      }
    }
    ...ProjectList_Query
  }
`;

const Header = React.memo(() => {
  const { t } = useTranslation();

  return (
    <TableHead>
      <TableRow>
        <TableCell>{t('projectList.reviewed')}</TableCell>
        <TableCell>{t('projectList.rank')}</TableCell>
        <TableCell>{t('projectList.name')}</TableCell>
        <TableCell>{t('projectList.link')}</TableCell>
      </TableRow>
    </TableHead>
  );
});

interface ProjectRowProps {
  project: ProjectList_project$data;
  isReviewed: boolean;
  onProjectClick: (code: string) => void;
  onProjectReviewClick: (id: string) => void;
}

const ProjectRow = React.memo<ProjectRowProps>(({ project: { id, rank, code, name }, isReviewed, onProjectClick, onProjectReviewClick }) => {
  const url = `${window.location.href}/${code}`;

  return (
    <TableRow>
      <TableCell padding="checkbox">
        <Checkbox checked={isReviewed} onClick={() => onProjectReviewClick(id)} />
      </TableCell>
      <TableCell>{rank}</TableCell>
      <TableCell>
        <Link onClick={() => onProjectClick(code)}>{name}</Link>
      </TableCell>
      <TableCell>
        <a href={url}>{url}</a>
      </TableCell>
    </TableRow>
  );
});

const ProjectRowRelayed = createFragmentContainer(ProjectRow, {
  project: graphql`
    fragment ProjectList_project on Project {
      id
      rank
      code
      name
    }
  `,
});

interface ProjectsTableProps {
  response: ProjectList_Query$data;
  isReviewedStates: Map<string, boolean>;
  onProjectClick: (code: string) => void;
  onProjectReviewClick: (id: string) => void;
  readonly relay: {
    environment: Environment;
    hasMore: () => boolean;
    isLoading: () => boolean;
    loadMore: (count: number, callBack: () => void) => boolean;
  };
}

const ProjectsTable = React.memo<ProjectsTableProps>(({ response, isReviewedStates, onProjectClick, onProjectReviewClick, relay }) => {
  const { t } = useTranslation();

  const getProjectsTable = (response: ProjectList_Query$data) => {
    // @ts-ignore: Object is possibly 'null'.
    return response.projects.edges.map((edge) => (
      <ProjectRowRelayed
        key={edge?.node?.id}
        // @ts-ignore: Object is possibly 'null'.
        project={edge?.node}
        // @ts-ignore: Object is possibly 'null'.
        isReviewed={isReviewedStates.get(edge?.node?.id)}
        onProjectClick={onProjectClick}
        onProjectReviewClick={onProjectReviewClick}
      />
    ));
  };

  const loadMore = () => {
    if (!relay.hasMore() || relay.isLoading()) {
      return;
    }

    relay.loadMore(pageSize, () => {});
  };

  return (
    <Paper>
      <Button variant="contained" onClick={loadMore} color="primary" disabled={!relay.hasMore()}>
        {t('projectList.loadMore')}
      </Button>
      <Table size="small">
        <Header />
        <TableBody>{getProjectsTable(response)}</TableBody>
      </Table>
    </Paper>
  );
});

const ProjectsTableRelayed = createPaginationContainer(
  ProjectsTable,
  {
    response: graphql`
      fragment ProjectList_Query on Query {
        projects(first: $count, after: $after, activityPeriod: LastWeek, orderBy: [{ direction: ASC, field: rank }])
          @connection(key: "Query_projects") {
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
              ...ProjectList_project
            }
          }
        }
      }
    `,
  },
  {
    direction: 'forward',
    query: rootQuery,
    getVariables: (_, { count, cursor }) => ({
      count,
      after: cursor,
    }),
    getFragmentVariables: (previousVars, totalCount) => ({
      ...previousVars,
      totalCount,
    }),
    getConnectionFromProps: (props) => props.response && props.response.projects,
  },
);

interface ProjectListContainerProps {
  response: ProjectListQuery$data;
  readonly relay: {
    environment: Environment;
  };
}

const ProjectListContainer = React.memo<ProjectListContainerProps>(({ response, relay: { environment } }) => {
  const navigate = useNavigate();
  const commitSetProjectIsReviewedMutation = useSetProjectIsReviewedMutation();
  const commitClearProjectIsReviewedMutationMutation = useClearProjectIsReviewedMutation();
  const { enqueueSnackbar } = useSnackbar();
  const [isReviewedStates, setIsReviewedStates] = useState(
    // @ts-ignore: Object is possibly 'null'.
    response.projects?.edges?.reduce(
      // @ts-ignore: Object is possibly 'null'.
      (reducation, val) => reducation.set(val?.node?.id, val?.node?.isReviewed),
      Map<string, boolean>(),
    ),
  );

  const handleProjectClick = (code: string) => {
    navigate(code);
  };

  const handleProjectReviewClick = (id: string) => {
    // @ts-ignore: Object is possibly 'null'.
    const projectNode = response.projects?.edges?.find((edge) => edge?.node?.id === id);

    if (!projectNode) {
      return;
    }

    const project = projectNode.node;

    if (!project) {
      return;
    }

    const isReviewedState = isReviewedStates?.get(id);

    if (isReviewedState === true) {
      commitClearProjectIsReviewedMutationMutation(
        {
          name: project.name,
          input: {
            id,
          },
        },
        {
          onSuccess: (mutationResponse) => {
            // @ts-ignore: Object is possibly 'null'.
            setIsReviewedStates(
              isReviewedStates?.set(
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.project?.id,
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.project?.isReviewed,
              ),
            );

            enqueueSnackbar(`Cleared reviewed for ${project.name}`, {
              variant: 'success',
            });
          },
          onError: (error) => {
            enqueueSnackbar(`Failed to clear reviewed ${project.name}. Error: ${error.message}`, { variant: 'error' });
          },
        },
      );
    } else {
      commitSetProjectIsReviewedMutation(
        {
          name: project.name,
          input: {
            id,
          },
        },
        {
          onSuccess: (mutationResponse) => {
            setIsReviewedStates(
              isReviewedStates?.set(
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.project?.id,
                // @ts-ignore: Object is possibly 'null'.
                mutationResponse.project?.isReviewed,
              ),
            );

            enqueueSnackbar(`Set reviewed for ${project.name}`, {
              variant: 'success',
            });
          },
          onError: (error) => {
            enqueueSnackbar(`Failed to set reviewed for ${project.name}. Error: ${error.message}`, { variant: 'error' });
          },
        },
      );
    }
  };

  return (
    <React.Fragment>
      <ProjectsTableRelayed
        response={response}
        // @ts-ignore: Object is possibly 'null'.
        isReviewedStates={isReviewedStates}
        onProjectClick={handleProjectClick}
        onProjectReviewClick={handleProjectReviewClick}
      />
    </React.Fragment>
  );
});

export default React.memo(() => {
  const { appToken } = useParams();

  if (!appToken) {
    return <GenericErrorContainer message="No app token provided!!!" />;
  }

  const relay = {
    environment: createEnvironment(appToken),
  };

  return (
    <QueryRenderer<ProjectListQuery>
      environment={relay.environment}
      query={rootQuery}
      variables={{
        count: pageSize,
      }}
      render={({ props, error }) => {
        if (error) {
          return <GenericErrorContainer message={error.message} />;
        } else if (props) {
          return <ProjectListContainer response={props} relay={relay} />;
        }

        return <LoadingContainer />;
      }}
    />
  );
});
