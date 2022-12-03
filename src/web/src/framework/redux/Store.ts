import { reduxBatch } from '@manaflair/redux-batch';
import { addListener, configureStore, createListenerMiddleware, ListenerEffectAPI, TypedAddListener, TypedStartListening } from '@reduxjs/toolkit';
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
import logger from 'redux-logger';
import createSagaMiddleware from 'redux-saga';
import { all } from 'redux-saga/effects';

import getRootReducer from './Reducer';

const rootSagas = function* sagas() {
  yield all([]);
};

const listenerMiddlewareInstance = createListenerMiddleware({
  onError: () => console.error,
});

const createStore = () => {
  const sagaMiddleware = createSagaMiddleware();
  const store = configureStore({
    reducer: getRootReducer,
    middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(sagaMiddleware).concat(logger).prepend(listenerMiddlewareInstance.middleware),
    devTools: process.env.NODE_ENV !== 'production',
    enhancers: [reduxBatch],
  });

  if (process.env.NODE_ENV === 'development' && module.hot) {
    module.hot.accept('./Reducer', () => {
      const newRootReducer = require('./Reducer').default;
      store.replaceReducer(newRootReducer);
    });
  }

  sagaMiddleware.run(rootSagas);

  return store;
};

export const store = createStore();

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// @see https://redux-toolkit.js.org/usage/usage-with-typescript#getting-the-dispatch-type
export type AppDispatch = typeof store.dispatch;

export type AppListenerEffectAPI = ListenerEffectAPI<RootState, AppDispatch>;

// @see https://redux-toolkit.js.org/api/createListenerMiddleware#typescript-usage
export type AppStartListening = TypedStartListening<RootState, AppDispatch>;
export type AppAddListener = TypedAddListener<RootState, AppDispatch>;

export const startAppListening = listenerMiddlewareInstance.startListening as AppStartListening;
export const addAppListener = addListener as AppAddListener;

// Use throughout your app instead of plain `useDispatch` and `useSelector`
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
