import { useEffect, useState } from 'react';
import { useSession } from 'next-auth/react';

import type { ConfigurationParameters } from '@follytics/sdk';

import type { ApiFactoryInterface } from '@self/lib/api/api-factory';
import { createApiFactory } from '@self/lib/api/api-factory';

const useApiFactory = (config?: ConfigurationParameters) => {
  const [apiFactory, setApiFactory] = useState<ApiFactoryInterface>();

  const { data: session, status } = useSession();

  useEffect(() => {
    if (status !== 'authenticated' || !session?.accessToken) return;

    const factory = createApiFactory({
      ...config,
      apiKey: session.accessToken,
      baseOptions: { headers: {} },
    });

    setApiFactory(factory);
  }, [config, session?.accessToken, status]);

  return apiFactory;
};

export default useApiFactory;
