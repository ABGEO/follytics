import type { UserApiGetUserFollowEventsRequest } from '@follytics/sdk';

import type { ApiFactoryInterface } from '@self/lib/api/api-factory';
import getServerApiFactory from '@self/lib/api/server-api-factory';

export async function fetchUserFollowEvents(
  request: UserApiGetUserFollowEventsRequest,
  apiFactory?: ApiFactoryInterface | null
) {
  if (!apiFactory) {
    apiFactory = await getServerApiFactory();
  }

  return apiFactory
    ?.getUserApi()
    .getUserFollowEvents(request)
    .then((res) => res.data);
}

export default fetchUserFollowEvents;
