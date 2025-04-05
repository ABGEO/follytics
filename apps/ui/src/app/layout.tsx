import { Inter } from 'next/font/google';
import type { Metadata } from 'next';
import { ReactNode } from 'react';
import { SessionProvider } from 'next-auth/react';

import '@ant-design/v5-patch-for-react-19';
import { AntdRegistry } from '@ant-design/nextjs-registry';
import { ConfigProvider } from 'antd';

import theme from '@self/theme';

import './globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Follytics',
  description: 'Follytics App',
};

type AppLayoutProps = {
  children: ReactNode;
};

function AppLayout({ children }: AppLayoutProps) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ConfigProvider theme={theme}>
          <AntdRegistry>
            <SessionProvider>{children}</SessionProvider>
          </AntdRegistry>
        </ConfigProvider>
      </body>
    </html>
  );
}

export default AppLayout;
