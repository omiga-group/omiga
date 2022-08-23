import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import { Provider } from 'react-redux';
import { I18nextProvider } from 'react-i18next';
import { BrowserRouter } from 'react-router-dom';
import { SnackbarProvider } from 'notistack';

import './App.css';
import i18n from './i18n';
import { store } from './framework/redux/Store';
import theme from './theme';
import AppRoutes from './AppRoutes';

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
            <SnackbarProvider maxSnack={5}>
              <CssBaseline />
              <AppRoutes />
            </SnackbarProvider>
          </ThemeProvider>
        </BrowserRouter>
      </Provider>
    </I18nextProvider>
  );
};

export default App;
