import { RawAxiosRequestConfig } from 'axios';

import type { UserApiGetUserFollowersTimelineRequest } from '@follytics/sdk';

import type { ApiFactoryInterface } from '@self/lib/api/api-factory';

async function fetchUserFollowersTimeline(
  apiFactory: ApiFactoryInterface,
  request: UserApiGetUserFollowersTimelineRequest,
  options?: RawAxiosRequestConfig,
) {
  const { data } = await apiFactory
    .getUserApi()
    .getUserFollowersTimeline(request, options);

  return data;
}

export default fetchUserFollowersTimeline;
