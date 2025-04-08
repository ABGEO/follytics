import { ReactNode } from 'react';

import { MainLayout as Layout } from '@self/layout/MainLayout';

type RootLayoutProps = {
  children: ReactNode;
};

function RootLayout({ children }: RootLayoutProps) {
  return <Layout>{children}</Layout>;
}

export default RootLayout;
