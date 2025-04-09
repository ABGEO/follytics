'use client';

import { useEffect, useMemo, useState } from 'react';

import { Alert } from 'antd';

import { normalizeError } from '@self/lib/error';

type ErrorProps = {
  error: Error & { digest?: string };
  reset: () => void;
};

function Error({ error }: ErrorProps) {
  const [errorMessage, setErrorMessage] = useState<string>('');

  const normalizedError = useMemo(() => normalizeError(error), [error]);

  useEffect(() => {
    console.error(normalizedError.errorObject);

    setErrorMessage(normalizedError.message);
  }, [normalizedError]);

  return (
    <div>
      <Alert message="Error" description={errorMessage} type="error" showIcon />
    </div>
  );
}

export default Error;
