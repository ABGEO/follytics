import useSWR from 'swr';

import { fetchCurrentUser } from './fetcher';
import useApiFactory from '@self/hooks/use-api-factory';

function useCurrentUser() {
  const apiFactory = useApiFactory();

  return useSWR(['userApi/getCurrentUser'], () => fetchCurrentUser(apiFactory));
}

export default useCurrentUser;
