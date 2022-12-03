import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import { SnackbarProvider } from 'notistack';
import { I18nextProvider } from 'react-i18next';
import { Provider } from 'react-redux';
import { RelayEnvironmentProvider } from 'react-relay';
import { BrowserRouter } from 'react-router-dom';

import './App.css';
import AppRoutes from './AppRoutes';
import { store } from './framework/redux/Store';
import { default as createEnvironment } from './framework/relay/Environment';
import i18n from './i18n';
import theme from './theme';

const Environment = createEnvironment();

let baseUrl = '/';
const base = document.getElementsByTagName('base');

if (base && base.length === 1) {
  baseUrl = base[0].getAttribute('href') as string;
}

const App = () => {
  return (
    <I18nextProvider i18n={i18n}>
      <Provider store={store}>
        <BrowserRouter basename={baseUrl}>
          <ThemeProvider theme={theme}>
            <CssBaseline />
            <RelayEnvironmentProvider environment={Environment}>
              <SnackbarProvider maxSnack={5}>
                <AppRoutes />
              </SnackbarProvider>
            </RelayEnvironmentProvider>
          </ThemeProvider>
        </BrowserRouter>
      </Provider>
    </I18nextProvider>
  );
};

export default App;
