'use client';

import {
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useRef,
  useState,
} from 'react';
import { useRouter } from 'next/navigation';
import { useSession } from 'next-auth/react';

import { Spin, message } from 'antd';

import type { ResponseUser } from '@follytics/sdk';

import { Logo } from '@self/components/Logo';
import { normalizeError } from '@self/lib/error';
import { useApiFactory } from '@self/providers/ApiFactoryProvider';

import './style.css';

type AuthContextType = {
  user: ResponseUser | null;
} | null;

type AuthProviderProps = {
  children: ReactNode;
};

const AuthContext = createContext<AuthContextType>(null);

function AuthProvider({ children }: AuthProviderProps) {
  const [authLoading, setAuthLoading] = useState<boolean>(true);
  const [user, setUser] = useState<ResponseUser | null>(null);

  const initializedRef = useRef<boolean>(false);

  const router = useRouter();
  const apiFactory = useApiFactory();
  const { status } = useSession();
  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    if (status === 'unauthenticated') {
      router.push('/');
    }
  }, [status, router]);

  useEffect(() => {
    (async () => {
      if (!apiFactory || status !== 'authenticated' || initializedRef.current) {
        return;
      }

      initializedRef.current = true;

      try {
        const response = await apiFactory.getUserApi().getCurrentUser();
        setUser(response?.data.data ?? null);
      } catch (error) {
        messageApi.open({
          type: 'error',
          content: normalizeError(error as Error).message,
        });
      } finally {
        setAuthLoading(false);
      }
    })();
  }, [apiFactory, status, messageApi]);

  return (
    <AuthContext.Provider value={{ user }}>
      {contextHolder}
      {authLoading && (
        <Spin
          fullscreen
          size="large"
          tip="Loading"
          indicator={<Logo height="6rem" animate />}
        />
      )}
      {children}
    </AuthContext.Provider>
  );
}

function useAuth() {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }

  return context;
}

export { AuthProvider, useAuth };
