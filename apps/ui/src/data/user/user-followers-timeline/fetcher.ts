import type { UserApiGetUserFollowersTimelineRequest } from '@follytics/sdk';

import type { ApiFactoryInterface } from '@self/lib/api/api-factory';

async function fetchUserFollowersTimeline(
  apiFactory: ApiFactoryInterface,
  request: UserApiGetUserFollowersTimelineRequest
) {
  const { data } = await apiFactory
    .getUserApi()
    .getUserFollowersTimeline(request);

  return data;
}

export default fetchUserFollowersTimeline;
