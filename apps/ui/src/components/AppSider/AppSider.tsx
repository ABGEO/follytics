'use client';

import { Layout } from 'antd';

import { AppMenu } from '@self/components/AppMenu';

function AppSider() {
  return (
    <Layout.Sider breakpoint="sm" collapsible>
      <AppMenu />
    </Layout.Sider>
  );
}

export { AppSider };
