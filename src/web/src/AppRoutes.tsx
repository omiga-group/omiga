import React from 'react';
import { Routes, Route } from 'react-router-dom';

import './App.css';
import NotificationHandlerContainer from './components/common/notification-handler/NotificationHandlerContainer';

import OrderList from './components/pages/order/OrderList';
import OrderDetails from './components/pages/order/OrderDetails';

const AppRoutes = () => {
  return (
    <React.Fragment>
      <Routes>
        <Route path="/" element={<OrderList />} />
        <Route path="/:id" element={<OrderDetails />} />
      </Routes>
      <NotificationHandlerContainer />
    </React.Fragment>
  );
};

export default AppRoutes;
