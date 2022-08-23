import React from 'react';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';

interface GenericErrorContainerProps {
  message: string;
}

export default React.memo<GenericErrorContainerProps>(({ message }) => {
  return (
    <div>
      <Container component="main" maxWidth="sm">
        <Typography variant="h2" gutterBottom>
          Something went wrong
        </Typography>
        <Typography variant="h5" gutterBottom>
          Error: {message}
        </Typography>
      </Container>
    </div>
  );
});
