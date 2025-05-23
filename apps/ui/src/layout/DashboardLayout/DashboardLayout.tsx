import { ReactNode } from 'react';
import { redirect } from 'next/navigation';

import { Content, Footer, Header } from 'antd/lib/layout/layout';
import { Flex, Layout } from 'antd';
import Title from 'antd/lib/typography/Title';

import { AccountDropdown } from '@self/components/AccountDropdown';
import { ApiFactoryProvider } from '@self/providers/ApiFactoryProvider';
import { AppBreadcrumb } from '@self/components/AppBreadcrumb/AppBreadcrumb';
import { AppSider } from '@self/components/AppSider';
import { AuthProvider } from '@self/providers/AuthProvider';
import { Logo } from '@self/components/Logo';
import { auth } from '@self/lib/auth';

import rootClasses from '@self/layout/MainLayout/MainLayout.module.css';

import classes from './DashboardLayout.module.css';

type DashboardLayoutProps = {
  children: ReactNode;
};

async function DashboardLayout({ children }: DashboardLayoutProps) {
  const session = await auth();

  if (!session) {
    redirect('/');
  }

  return (
    <ApiFactoryProvider>
      <AuthProvider>
        <Layout className={rootClasses.rootLayout}>
          <Header className={`${rootClasses.header} ${classes.header}`}>
            <Flex
              className={rootClasses.headerChildren}
              align="center"
              justify="space-between"
            >
              <Flex align="center">
                <Logo width="3em" height="3em" />

                <Title
                  className={`${rootClasses.appName} ${classes.appName}`}
                  level={4}
                >
                  Follytics
                </Title>
              </Flex>
              <AccountDropdown />
            </Flex>
          </Header>
          <Layout hasSider>
            <AppSider />
            <Layout className={classes.contentLayout}>
              <AppBreadcrumb />
              <Content className={`${rootClasses.content} ${classes.content}`}>
                {children}
              </Content>
              <Footer className={`${rootClasses.footer} ${classes.footer}`}>
                Follytics ©2025-{new Date().getFullYear()}
              </Footer>
            </Layout>
          </Layout>
        </Layout>
      </AuthProvider>
    </ApiFactoryProvider>
  );
}

export { DashboardLayout };
