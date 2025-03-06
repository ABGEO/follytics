'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

import { useSession } from 'next-auth/react';

import { Breadcrumb, Flex, Layout, Menu, Spin, theme } from 'antd';
import { DashboardOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';

const { Header, Content, Sider, Footer } = Layout;

import { AccountDropdown } from '@self/components/account-dropdown';

const menuItems: MenuProps['items'] = [
  {
    key: 'dashboard',
    icon: <DashboardOutlined />,
    label: 'Dashboard',
  },
];

export default function DashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [collapsed, setCollapsed] = useState(false);
  const [authLoading, setAuthLoading] = useState<boolean>(true);

  const { status } = useSession();
  const router = useRouter();
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  useEffect(() => {
    if (status === 'loading') {
      setAuthLoading(false);
    }

    if (status === 'unauthenticated') {
      router.push('/');
    }
  }, [status, router]);

  if (authLoading) {
    return <Spin fullscreen size="large" tip="Loading" />;
  }

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Header style={{ display: 'flex', alignItems: 'center' }}>
        <div className="demo-logo" />

        <Flex style={{ width: '100%' }} justify="end">
          <AccountDropdown />
        </Flex>
      </Header>
      <Layout>
        <Sider
          collapsible
          collapsed={collapsed}
          onCollapse={(value) => setCollapsed(value)}
          width={200}
          style={{ background: colorBgContainer }}
        >
          <Menu
            mode="inline"
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            style={{ height: '100%', borderRight: 0 }}
            items={menuItems}
          />
        </Sider>
        <Layout style={{ padding: '0 24px 24px' }}>
          <Breadcrumb
            items={[{ title: 'Dashboard' }, { title: 'App' }]}
            style={{ margin: '16px 0' }}
          />
          <Content
            style={{
              padding: 24,
              margin: 0,
              marginBlock: 25,
              minHeight: 280,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            {children}
          </Content>
          <Footer style={{ textAlign: 'center' }}>
            Follytics ©2025-{new Date().getFullYear()}
          </Footer>
        </Layout>
      </Layout>
    </Layout>
  );
}
