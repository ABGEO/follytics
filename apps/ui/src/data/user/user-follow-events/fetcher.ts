import type { UserApiGetUserFollowEventsRequest } from '@follytics/sdk';

import type { ApiFactoryInterface } from '@self/lib/api/api-factory';

async function fetchUserFollowEvents(
  apiFactory: ApiFactoryInterface,
  request: UserApiGetUserFollowEventsRequest
) {
  const { data } = await apiFactory.getUserApi().getUserFollowEvents(request);

  return data;
}

export default fetchUserFollowEvents;
