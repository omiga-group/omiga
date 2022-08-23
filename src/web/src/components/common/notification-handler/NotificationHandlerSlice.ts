import { createSlice, createEntityAdapter } from '@reduxjs/toolkit';

export enum NotificationType {
  Error,
  Warning,
  Info,
  Success,
}

export interface Notification {
  id: string;
  type: NotificationType;
  message: string;
}

const notificationEntity = createEntityAdapter<Notification>();

export const notificationSlice = createSlice({
  name: 'notifications',
  initialState: notificationEntity.getInitialState(),
  reducers: {
    notificationAdded: notificationEntity.addOne,
    notificationRemoved: notificationEntity.removeOne,
  },
});

export const notificationActions = notificationSlice.actions;

export type NotificationSlice = {
  [notificationSlice.name]: ReturnType<typeof notificationSlice['reducer']>;
};

export const notificationSelectors = notificationEntity.getSelectors<NotificationSlice>((state) => state[notificationSlice.name]);
