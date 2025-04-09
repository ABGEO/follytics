import useSWR from 'swr';

import { type ApiFactoryInterface } from '@self/lib/api/api-factory';
import useApiFactory from '@self/hooks/use-api-factory';

import fetchCurrentUser from './fetcher';

function useCurrentUser() {
  const apiFactory = useApiFactory();

  const key = apiFactory ? ['userApi/getCurrentUser'] : null;

  const result = useSWR(key, () =>
    fetchCurrentUser(apiFactory as ApiFactoryInterface),
  );

  if (result.error) {
    throw result.error;
  }

  return result;
}

export default useCurrentUser;
