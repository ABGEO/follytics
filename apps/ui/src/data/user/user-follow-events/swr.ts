import useSWR from 'swr';

import type { UserApiGetUserFollowEventsRequest } from '@follytics/sdk';

import { type ApiFactoryInterface } from '@self/lib/api/api-factory';
import useApiFactory from '@self/hooks/use-api-factory';

import fetchUserFollowEvents from './fetcher';

function useUserFollowEvents(request: UserApiGetUserFollowEventsRequest) {
  const apiFactory = useApiFactory();

  const key =
    apiFactory && request.id
      ? ['userApi/getUserFollowEvents', request.id, request.page, request.limit]
      : null;

  return useSWR(key, () =>
    fetchUserFollowEvents(apiFactory as ApiFactoryInterface, request)
  );
}

export default useUserFollowEvents;
