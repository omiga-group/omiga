import React from 'react';
import Stack from '@mui/material/Stack';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';

import { useAppDispatch, useAppSelector } from '../../../framework/redux/Store';
import { notificationActions, notificationSelectors, NotificationType } from './NotificationHandlerSlice';

export default React.memo(() => {
  const appDispatch = useAppDispatch();
  const storeNotifications = useAppSelector((state) => notificationSelectors.selectAll(state));

  var context = storeNotifications.map((notification) => ({
    key: notification.id,
    type: notification.type,
    message: notification.message,
    onClose: (_?: React.SyntheticEvent | Event, reason?: string) => {
      if (reason === 'clickaway') {
        return;
      }

      appDispatch(notificationActions.notificationRemoved(notification.id));
    },
  }));

  // Below code is buggy, do not use it!!!!!

  return (
    <Stack spacing={2} sx={{ width: '100%' }}>
      {context.length > 0 && (
        <Snackbar open={true} autoHideDuration={1000}>
          <React.Fragment>
            {context.map((notification) => {
              if (notification.type === NotificationType.Error) {
                return (
                  <Alert key={notification.key} elevation={6} onClose={notification.onClose} severity="error">
                    <AlertTitle>Error</AlertTitle>
                    {notification.message}
                  </Alert>
                );
              }

              if (notification.type === NotificationType.Warning) {
                return (
                  <Alert key={notification.key} elevation={6} onClose={notification.onClose} severity="warning">
                    <AlertTitle>Warning</AlertTitle>
                    {notification.message}
                  </Alert>
                );
              }

              if (notification.type === NotificationType.Info) {
                return (
                  <Alert key={notification.key} elevation={6} onClose={notification.onClose} severity="info">
                    <AlertTitle>Info</AlertTitle>
                    {notification.message}
                  </Alert>
                );
              }

              if (notification.type === NotificationType.Success) {
                return (
                  <Alert key={notification.key} elevation={6} onClose={notification.onClose} severity="success">
                    <AlertTitle>Success</AlertTitle>
                    {notification.message}
                  </Alert>
                );
              }

              throw new Error(`Notification type: ${notification.type} is not supported`);
            })}
          </React.Fragment>
        </Snackbar>
      )}
    </Stack>
  );
});
