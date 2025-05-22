import { GoogleAnalytics } from '@next/third-parties/google';
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
  title: 'Follytics - Track Your GitHub Followers Over Time',
  description:
    'Follytics is a free and open-source tool that lets you track, visualize, and analyze your GitHub followers using an event-sourcing pattern.',
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
      <GoogleAnalytics gaId="G-TE0GT7TD6Y" />
    </html>
  );
}

export default AppLayout;
