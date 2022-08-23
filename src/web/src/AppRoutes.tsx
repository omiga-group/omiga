import React from 'react';
import { Routes, Route } from 'react-router-dom';

import './App.css';
import NotificationHandlerContainer from './components/common/notification-handler/NotificationHandlerContainer';

const AppRoutes = () => {
  return (
    <React.Fragment>
      <Routes>
      </Routes>
      <NotificationHandlerContainer />
    </React.Fragment>
  );
};

export default AppRoutes;
