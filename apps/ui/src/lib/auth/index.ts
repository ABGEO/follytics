import { AxiosError, HttpStatusCode, isAxiosError } from 'axios';
import GitHub from 'next-auth/providers/github';
import NextAuth from 'next-auth';

import { createApiFactory } from '../api/api-factory';

export const { handlers, signIn, signOut, auth } = NextAuth({
  providers: [
    GitHub({
      authorization: {
        params: { prompt: 'consent' },
      },
    }),
  ],
  pages: {
    error: '/error',
  },
  session: {
    maxAge: 7 * 24 * 60 * 60, // 7 days.
  },
  callbacks: {
    jwt: async ({ token, account }) => {
      if (account && account.access_token) {
        token.accessToken = account.access_token;

        return token;
      }

      return token;
    },
    session: async ({ session, token }) => {
      session.accessToken = token.accessToken as string;

      return session;
    },
    async signIn({ account, profile }) {
      if (account && profile) {
        const apiFactory = createApiFactory({
          basePath: process.env.NEXT_PUBLIC_API_URL,
          apiKey: account.access_token,
          baseOptions: {
            headers: {
              'X-Request-Origin': 'server',
              'X-Component': 'nextauth',
            },
          },
        });

        try {
          await apiFactory.getUserApi().trackLogin();
        } catch (error) {
          if (isAxiosError(error)) {
            const axiosError = error as AxiosError;

            switch (axiosError.response?.status) {
              case HttpStatusCode.Unauthorized:
                return false;
            }
          }

          return '/error?error=AuthUnknownError';
        }
      }

      return true;
    },
  },
});
