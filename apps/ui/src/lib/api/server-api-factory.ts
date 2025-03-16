import { auth } from '@self/lib/auth';

import type { ConfigurationParameters } from '@follytics/sdk';

import type { ApiFactoryInterface } from '@self/lib/api/api-factory';
import { createApiFactory } from '@self/lib/api/api-factory';

async function getServerApiFactory(
  config?: ConfigurationParameters
): Promise<ApiFactoryInterface | null> {
  const session = await auth();

  if (!session || !session.accessToken) {
    return null;
  }

  return createApiFactory({
    ...config,
    apiKey: session.accessToken,
    baseOptions: {
      headers: {
        'X-Request-Origin': 'server',
      },
    },
  });
}

export default getServerApiFactory;
