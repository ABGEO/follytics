import useSWR from 'swr';

import type { UserApiGetUserFollowersRequest } from '@follytics/sdk';

import { type ApiFactoryInterface } from '@self/lib/api/api-factory';
import useApiFactory from '@self/hooks/use-api-factory';

import fetchUserFollowers from './fetcher';

function useUserFollowers(request: UserApiGetUserFollowersRequest) {
  const apiFactory = useApiFactory();

  const key =
    apiFactory && request.id
      ? ['userApi/getUserFollowers', request.id, request.page, request.limit]
      : null;

  const result = useSWR(key, () =>
    fetchUserFollowers(apiFactory as ApiFactoryInterface, request),
  );

  if (result.error) {
    throw result.error;
  }

  return result;
}

export default useUserFollowers;
