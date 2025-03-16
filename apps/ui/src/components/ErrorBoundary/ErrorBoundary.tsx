import { ReactNode, useMemo } from 'react';
import { Alert } from 'antd';

import { getApiErrorMessage } from '@self/lib/api/error';

type ErrorBoundaryProps = {
  error: unknown;
  heading?: string;
  children: ReactNode;
};

function ErrorBoundary({
  error,
  heading = 'Error',
  children,
}: ErrorBoundaryProps) {
  const errorMessage = useMemo(() => getApiErrorMessage(error), [error]);

  if (!error) return <>{children}</>;

  return (
    <Alert message={heading} description={errorMessage} type="error" showIcon />
  );
}

export { ErrorBoundary };
