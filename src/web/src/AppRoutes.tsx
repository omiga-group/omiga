import React from 'react';
import { Routes, Route } from 'react-router-dom';

import './App.css';
import NotificationHandlerContainer from './components/common/notification-handler/NotificationHandlerContainer';

import OrderList from './components/pages/order/OrderList';
import SubmitOrder from './components/pages/order/SubmitOrder';
import OrderDetails from './components/pages/order/OrderDetails';

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
