import { isAxiosError } from 'axios';

const getApiErrorMessage = (error: unknown): string => {
  if (!error) return '';

  if (typeof error === 'string') return error;

  if (!isAxiosError(error)) return 'Unknown error';

  const { response, message } = error;
  const data = response?.data;

  if (!data) return message;

  if (typeof data === 'string') return data;

  if (typeof data === 'object') {
    return (
      data.message ??
      (typeof data.error === 'string'
        ? data.error
        : JSON.stringify(data.error ?? data))
    );
  }

  return message;
};

export { getApiErrorMessage };
