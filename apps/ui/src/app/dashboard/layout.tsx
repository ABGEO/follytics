import { ReactNode } from 'react';

import { DashboardLayout as Layout } from '@self/layout/DashboardLayout';

type DashboardLayoutProps = {
  children: ReactNode;
};

async function DashboardLayout({ children }: DashboardLayoutProps) {
  return <Layout>{children}</Layout>;
}

export default DashboardLayout;
