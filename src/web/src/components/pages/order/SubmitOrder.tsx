import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useTheme } from '@mui/material/styles';
import { Theme } from '@mui/material/styles';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import { SxProps } from '@mui/system';
import { useTranslation } from 'react-i18next';
import { Form, Field } from 'react-final-form';
import { useSnackbar } from 'notistack';

import { renderTextField, renderAutocomplete } from '../../common/react-final-form-components';
import { default as useSubmitOrderMutation } from './SubmitOrderMutation';

export const enNZTranslation = {
  title: 'Submit New Order',
  from: 'From',
  amount: 'Amount',
  to: 'To',
  required: 'Required',
  submitOrder: 'Submit Order',
  cancel: 'Cancel',
};

const getSubmitOrderButtonStyle = (theme: Theme): SxProps => ({
  margin: theme.spacing(1, 1, 0),
});

const getCancelButtonStyle = (theme: Theme): SxProps => ({
  margin: theme.spacing(1, 1, 0),
});

interface SubmitOrderProps {}

export interface Values {
  from: string;
  amount: string;
  to: string;
}

export default React.memo<SubmitOrderProps>(() => {
  const theme = useTheme();
  const navigate = useNavigate();
  const commitSubmitOrderMutation = useSubmitOrderMutation();
  const submitOrderStyle = getSubmitOrderButtonStyle(theme);
  const cancelStyle = getCancelButtonStyle(theme);
  const { enqueueSnackbar } = useSnackbar();
  const { t } = useTranslation();

  const handleSubmit = (values: Values) => {
    commitSubmitOrderMutation(
      {
        input: {
          orderDetails: {
            baseCurrency: {
              code: values.from,
              digital: true,
              maxPrecision: 10,
              name: 'XXX',
            },
            counterCurrency: {
              code: values.to,
              digital: true,
              maxPrecision: 10,
              name: 'XXX',
            },
            price: {
              amount: parseFloat(values.amount),
              scale: 0,
            },
            quantity: {
              amount: parseFloat(values.amount),
              scale: 0,
            },
            side: 'BID',
            type: 'INSTANT',
          },
        },
      },
      {
        onSuccess: (mutationResponse) => {
          enqueueSnackbar(`Submitted order with Id ${mutationResponse.submitOrder?.order?.id}`, {
            variant: 'success',
          });
          navigate(-1);
        },
        onError: (error) => {
          enqueueSnackbar(`Failed to submit order, error: ${error.message}`, { variant: 'error' });
        },
      },
    );
  };

  const handleCancel = () => {
    navigate(-1);
  };

  const requiredValidation = (value: string) => (value && value.trim() ? undefined : t('submitOrder.required'));

  return (
    <React.Fragment>
      <Typography variant="h5">{t('submitOrder.title')}</Typography>
      <Form
        onSubmit={handleSubmit}
        initialValues={{
          name: '',
        }}
        render={({ handleSubmit, form, submitting, invalid }) => (
          <form onSubmit={handleSubmit}>
            <React.Fragment>
              <Field<string>
                name="from"
                variant="outlined"
                margin="normal"
                required
                fullWidth
                placeholder={t('submitOrder.from')}
                component={renderAutocomplete}
                autoFocus
                options={['BTC', 'ETH']}
                validate={requiredValidation}
              />
              <Field<string>
                name="amount"
                variant="outlined"
                margin="normal"
                required
                fullWidth
                placeholder={t('submitOrder.amount')}
                component={renderTextField}
                autoFocus
                validate={requiredValidation}
              />
              <Field<string>
                name="to"
                variant="outlined"
                margin="normal"
                required
                fullWidth
                placeholder={t('submitOrder.to')}
                component={renderAutocomplete}
                autoFocus
                options={['BTC', 'ETH']}
                validate={requiredValidation}
              />
            </React.Fragment>
            <Button type="submit" variant="contained" color="primary" sx={submitOrderStyle} disabled={submitting || invalid}>
              {t('submitOrder.submitOrder')}
            </Button>
            <Button type="button" variant="contained" color="secondary" disabled={submitting} sx={cancelStyle} onClick={handleCancel}>
              {t('submitOrder.cancel')}
            </Button>
          </form>
        )}
      />
    </React.Fragment>
  );
});
