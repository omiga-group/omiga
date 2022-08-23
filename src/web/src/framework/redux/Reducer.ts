import { notificationSlice } from '../../components/common/notification-handler/NotificationHandlerSlice';

const reducer = {
  [notificationSlice.name]: notificationSlice.reducer,
};

export default reducer;
