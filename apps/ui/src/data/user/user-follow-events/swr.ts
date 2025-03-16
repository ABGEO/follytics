import useSWR from 'swr';

import type { UserApiGetUserFollowEventsRequest } from '@follytics/sdk';

import { fetchUserFollowEvents } from './fetcher';
import useApiFactory from '@self/hooks/use-api-factory';

function useUserFollowEvents(request: UserApiGetUserFollowEventsRequest) {
  const apiFactory = useApiFactory();

  return useSWR(
    ['userApi/getUserFollowEvents', request.id, request.page, request.limit],
    () => fetchUserFollowEvents(request, apiFactory)
  );
}

export default useUserFollowEvents;
