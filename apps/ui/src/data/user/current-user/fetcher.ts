import type { ApiFactoryInterface } from '@self/lib/api/api-factory';
import getServerApiFactory from '@self/lib/api/server-api-factory';

export async function fetchCurrentUser(
  apiFactory?: ApiFactoryInterface | null
) {
  if (!apiFactory) {
    apiFactory = await getServerApiFactory();
  }

  return apiFactory
    ?.getUserApi()
    .getCurrentUser()
    .then((res) => res.data);
}

export default fetchCurrentUser;
