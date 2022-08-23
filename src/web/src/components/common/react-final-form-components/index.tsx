import React, { ChangeEvent } from 'react';
import { FieldRenderProps } from 'react-final-form';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';

type Props = FieldRenderProps<string, any>;

export const renderTextField: React.FC<Props> = ({ input, meta: { error, dirty, submitFailed, submitError }, placeholder, ...rest }: Props) => {
  const errorMessage = error || submitError;
  const hasError = error && (dirty || submitFailed);

  return <TextField {...input} {...rest} error={hasError} helperText={hasError ? errorMessage : ''} placeholder={placeholder} />;
};

export const renderAutocomplete: React.FC<Props> = ({
  input,
  meta: { error, dirty, submitFailed, submitError },
  placeholder,
  options}: Props) => {
  const errorMessage = error || submitError;
  const hasError = error && (dirty || submitFailed);
  const { name, onChange, value, ...restInput } = input;

  return (
    <Autocomplete
      value={value}
      onChange={(_: ChangeEvent<{}>, option: string | null): void => onChange(option ? option : null)}
      options={options}
      renderInput={(params) => (
        <TextField
          {...params}
          {...restInput}
          error={hasError}
          helperText={hasError ? errorMessage : ''}
          placeholder={placeholder}
          variant="outlined"
        />
      )}
    />
  );
};
