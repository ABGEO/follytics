import { AxiosError, isAxiosError } from 'axios';
import { ErrorObject, serializeError } from 'serialize-error';

type NormalizedError = {
  message: string;
  errorObject: ErrorObject;
};

const getAxiosErrorMessage = (error: AxiosError): string => {
  let errorMessage = '';
  const { name, message, code, response } = error;

  if (code) errorMessage += `[${code}] `;
  if (name) errorMessage += `${name}: `;

  const data = response?.data;

  if (!data) return `${errorMessage}${message}`;

  if (typeof data === 'string') return `${errorMessage}${data}`;

  if (typeof data !== 'object') {
    return `${errorMessage}${message}`;
  }

  for (const key of [
    'message',
    'error',
    'errorMessage',
    'description',
    'reason',
    'messages',
    'errors',
    'errorMessages',
    'descriptions',
    'reasons',
  ] as const) {
    const value = (data as Record<string, unknown>)[key];

    if (typeof value === 'string') {
      return `${errorMessage}${value}`;
    }

    if (Array.isArray(value)) {
      const strings = value.filter((v): v is string => typeof v === 'string');

      if (strings.length > 0) {
        return `${errorMessage}${strings.join('\n')}`;
      }
    }
  }

  return `${errorMessage}${JSON.stringify(data)}`;
};

const getSerializedErrorMessage = (serializedError: ErrorObject): string => {
  let errorMessage = '';

  const { name, message, code, digest } = serializedError;

  if (code) errorMessage += `[${code}] `;
  if (digest) errorMessage += `[${digest}] `;
  if (name) errorMessage += `${name}: `;
  if (message) errorMessage += message;
  if (errorMessage) return errorMessage;

  return 'Unknown error';
};

const normalizeError = (error: Error): NormalizedError => {
  const serializedError = serializeError(error);

  if (typeof error === 'string') {
    return { message: error, errorObject: serializedError };
  }

  if (isAxiosError(error)) {
    return {
      message: getAxiosErrorMessage(error),
      errorObject: serializedError,
    };
  }

  return {
    message: getSerializedErrorMessage(serializedError),
    errorObject: serializedError,
  };
};

export {
  type NormalizedError,
  getAxiosErrorMessage,
  getSerializedErrorMessage,
  normalizeError,
};
