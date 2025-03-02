import { AxiosError, HttpStatusCode, isAxiosError } from 'axios';
import GitHub from 'next-auth/providers/github';
import NextAuth from 'next-auth';

import { Configuration, UserApi } from '@follytics/sdk';

export const { handlers, signIn, signOut, auth } = NextAuth({
  providers: [GitHub],
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
        const userApi = new UserApi(
          new Configuration({ apiKey: account.access_token })
        );

        try {
          await userApi.trackLogin();
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
