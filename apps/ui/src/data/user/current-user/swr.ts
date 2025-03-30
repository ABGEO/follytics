import useSWR from 'swr';

import { type ApiFactoryInterface } from '@self/lib/api/api-factory';
import useApiFactory from '@self/hooks/use-api-factory';

import fetchCurrentUser from './fetcher';

function useCurrentUser() {
  const apiFactory = useApiFactory();

  const key = apiFactory ? ['userApi/getCurrentUser'] : null;

  return useSWR(key, () => fetchCurrentUser(apiFactory as ApiFactoryInterface));
}

export default useCurrentUser;
