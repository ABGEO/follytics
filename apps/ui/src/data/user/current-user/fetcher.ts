import type { ApiFactoryInterface } from '@self/lib/api/api-factory';

export async function fetchCurrentUser(apiFactory: ApiFactoryInterface) {
  const { data } = await apiFactory.getUserApi().getCurrentUser();

  return data;
}

export default fetchCurrentUser;
