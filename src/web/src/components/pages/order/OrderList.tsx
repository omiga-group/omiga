import React from 'react';
import graphql from 'babel-plugin-relay/macro';
import { createFragmentContainer, createPaginationContainer } from 'react-relay';
import { Environment } from 'relay-runtime';
import { QueryRenderer } from 'react-relay';
import { useNavigate } from 'react-router-dom';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableCell from '@mui/material/TableCell';
import Link from '@mui/material/Link';
import Button from '@mui/material/Button';
import { useTranslation } from 'react-i18next';

import { OrderList_order$data } from './__generated__/OrderList_order.graphql';
import { OrderList_Query$data } from './__generated__/OrderList_Query.graphql';
import { OrderListQuery, OrderListQuery$data } from './__generated__/OrderListQuery.graphql';

import { default as createEnvironment } from '../../../framework/relay/Environment';
import LoadingContainer from '../../common/loading/LoadingContainer';
import GenericErrorContainer from '../../common/generic-error/GenericErrorContainer';

export const enNZTranslation = {
  loadMore: 'Load more',
  id: 'Id',
};

const pageSize = 10000;

const rootQuery = graphql`
  query OrderListQuery($count: Int!, $after: Cursor) {
    orders(first: $count, after: $after) @connection(key: "Query_orders") {
      edges {
        cursor
        node {
          id
        }
      }
    }
    ...OrderList_Query
  }
`;

const Header = React.memo(() => {
  const { t } = useTranslation();

  return (
    <TableHead>
      <TableRow>
        <TableCell>{t('orderList.id')}</TableCell>
      </TableRow>
    </TableHead>
  );
});

interface OrderRowProps {
  order: OrderList_order$data;
  onOrderClick: (code: string) => void;
}

const OrderRow = React.memo<OrderRowProps>(({ order: { id }, onOrderClick }) => {
  const url = `${window.location.href}/${id}`;

  return (
    <TableRow>
      <TableCell>{id}</TableCell>
      <TableCell>
        <Link onClick={() => onOrderClick(id)}>{id}</Link>
      </TableCell>
      <TableCell>
        <a href={url}>{url}</a>
      </TableCell>
    </TableRow>
  );
});

const OrderRowRelayed = createFragmentContainer(OrderRow, {
  order: graphql`
    fragment OrderList_order on Order {
      id
    }
  `,
});

interface OrdersTableProps {
  response: OrderList_Query$data;
  onOrderClick: (code: string) => void;
  readonly relay: {
    environment: Environment;
    hasMore: () => boolean;
    isLoading: () => boolean;
    loadMore: (count: number, callBack: () => void) => boolean;
  };
}

const OrdersTable = React.memo<OrdersTableProps>(({ response, onOrderClick, relay }) => {
  const { t } = useTranslation();

  const getOrdersTable = (response: OrderList_Query$data) => {
    // @ts-ignore: Object is possibly 'null'.
    return response.orders.edges.map((edge) => (
      <OrderRowRelayed
        key={edge?.node?.id}
        // @ts-ignore: Object is possibly 'null'.
        order={edge?.node}
        onOrderClick={onOrderClick}
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
        {t('orderList.loadMore')}
      </Button>
      <Table size="small">
        <Header />
        <TableBody>{getOrdersTable(response)}</TableBody>
      </Table>
    </Paper>
  );
});

const OrdersTableRelayed = createPaginationContainer(
  OrdersTable,
  {
    response: graphql`
      fragment OrderList_Query on Query {
        orders(first: $count, after: $after) @connection(key: "Query_orders") {
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
              ...OrderList_order
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
    getConnectionFromProps: (props) => props.response && props.response.orders,
  },
);

interface OrderListContainerProps {
  response: OrderListQuery$data;
  readonly relay: {
    environment: Environment;
  };
}

const OrderListContainer = React.memo<OrderListContainerProps>(({ response, relay: { environment } }) => {
  const navigate = useNavigate();

  const handleOrderClick = (code: string) => {
    navigate(code);
  };

  return (
    <React.Fragment>
      <OrdersTableRelayed
        response={response}
        // @ts-ignore: Object is possibly 'null'.
        isReviewedStates={isReviewedStates}
        onOrderClick={handleOrderClick}
      />
    </React.Fragment>
  );
});

export default React.memo(() => {
  const relay = {
    environment: createEnvironment(),
  };

  return (
    <QueryRenderer<OrderListQuery>
      environment={relay.environment}
      query={rootQuery}
      variables={{
        count: pageSize,
      }}
      render={({ props, error }) => {
        if (error) {
          return <GenericErrorContainer message={error.message} />;
        } else if (props) {
          return <OrderListContainer response={props} relay={relay} />;
        }

        return <LoadingContainer />;
      }}
    />
  );
});
