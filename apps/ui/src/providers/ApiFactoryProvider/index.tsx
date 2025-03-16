'use client';

import { ReactNode, createContext, useContext } from 'react';

import { type ApiFactoryInterface } from '@self/lib/api/api-factory';
import useApiFactoryHook from '@self/hooks/use-api-factory';

type ApiFactoryProviderType = {
  factory?: ApiFactoryInterface;
};

type ApiFactoryProviderProps = {
  children: ReactNode;
};

const ApiFactoryContext = createContext<ApiFactoryProviderType | null>(null);

function ApiFactoryProvider({ children }: ApiFactoryProviderProps) {
  const factory = useApiFactoryHook();

  return (
    <ApiFactoryContext.Provider value={{ factory }}>
      {children}
    </ApiFactoryContext.Provider>
  );
}

function useApiFactory() {
  const context = useContext(ApiFactoryContext);

  if (!context) {
    throw new Error('useApiFactory must be used within an ApiFactoryContext');
  }

  return context.factory;
}

export { ApiFactoryProvider, useApiFactory };
