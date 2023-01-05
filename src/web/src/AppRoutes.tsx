import React from 'react';
import { Route, Routes } from 'react-router-dom';

import './App.css';
import NotificationHandlerContainer from './components/common/notification-handler/NotificationHandlerContainer';

import OrderDetails from './components/pages/order/OrderDetails';
import OrderList from './components/pages/order/OrderList';
import SubmitOrder from './components/pages/order/SubmitOrder';

const AppRoutes = () => {
  return (
    <React.Fragment>
      <Routes>
        <Route path="/" element={<OrderList />} />
        <Route path="/submitOrder" element={<SubmitOrder />} />
        <Route path="/:id" element={<OrderDetails />} />
      </Routes>
      <NotificationHandlerContainer />
    </React.Fragment>
  );
};

export default AppRoutes;
