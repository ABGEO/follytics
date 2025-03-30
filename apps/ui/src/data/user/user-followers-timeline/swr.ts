import useSWR from 'swr';

import type { UserApiGetUserFollowersTimelineRequest } from '@follytics/sdk';

import { type ApiFactoryInterface } from '@self/lib/api/api-factory';
import useApiFactory from '@self/hooks/use-api-factory';

import fetchUserFollowersTimeline from './fetcher';

function useUserFollowersTimeline(
  request: UserApiGetUserFollowersTimelineRequest
) {
  const apiFactory = useApiFactory();

  const key =
    apiFactory && request.id
      ? ['userApi/getUserFollowersTimeline', request.id]
      : null;

  return useSWR(key, () =>
    fetchUserFollowersTimeline(apiFactory as ApiFactoryInterface, request)
  );
}

export default useUserFollowersTimeline;
