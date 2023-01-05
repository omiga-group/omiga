import graphql from 'babel-plugin-relay/macro';
import React from 'react';
import { createFragmentContainer, QueryRenderer } from 'react-relay';
import { useParams } from 'react-router-dom';

import { OrderDetailsQuery, OrderDetailsQuery$data } from './__generated__/OrderDetailsQuery.graphql';
import { OrderDetails_Query$data } from './__generated__/OrderDetails_Query.graphql';

import { default as createEnvironment } from '../../../framework/relay/Environment';
import GenericErrorContainer from '../../common/generic-error/GenericErrorContainer';
import LoadingContainer from '../../common/loading/LoadingContainer';

export const enNZTranslation = {
  id: 'Id',
};

const rootQuery = graphql`
  query OrderDetailsQuery($id: ID!) {
    ...OrderDetails_Query
  }
`;

interface OrderProps {
  response: OrderDetails_Query$data;
}

const Order = React.memo<OrderProps>(({ response }) => {
  if (!response.order) {
    return <GenericErrorContainer message="order not found!!!" />;
  }

  return <React.Fragment></React.Fragment>;
});

const OrderRelayed = createFragmentContainer(Order, {
  response: graphql`
    fragment OrderDetails_Query on Query {
      order(where: { id: $id }) {
        id
      }
    }
  `,
});

interface OrderDetailsContainerProps {
  response: OrderDetailsQuery$data;
}

const OrderDetailsContainer = React.memo<OrderDetailsContainerProps>(({ response }) => {
  return (
    <React.Fragment>
      <OrderRelayed response={response} />
    </React.Fragment>
  );
});

export default React.memo(() => {
  const { id } = useParams();

  if (!id) {
    return <GenericErrorContainer message="No Order Id provided!!!" />;
  }

  const relay = {
    environment: createEnvironment(),
  };

  return (
    <QueryRenderer<OrderDetailsQuery>
      environment={relay.environment}
      query={rootQuery}
      variables={{
        id: id,
      }}
      render={({ props, error }) => {
        if (error) {
          return <GenericErrorContainer message={error.message} />;
        } else if (props) {
          return <OrderDetailsContainer response={props} />;
        }

        return <LoadingContainer />;
      }}
    />
  );
});
