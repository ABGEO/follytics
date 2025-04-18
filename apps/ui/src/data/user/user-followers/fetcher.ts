import type { UserApiGetUserFollowersRequest } from '@follytics/sdk';

import type { ApiFactoryInterface } from '@self/lib/api/api-factory';

async function fetchUserFollowers(
  apiFactory: ApiFactoryInterface,
  request: UserApiGetUserFollowersRequest,
) {
  const { data } = await apiFactory.getUserApi().getUserFollowers(request);

  return data;
}

export default fetchUserFollowers;
