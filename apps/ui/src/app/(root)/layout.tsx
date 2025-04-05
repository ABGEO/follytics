import { ReactNode } from 'react';

import { Content, Footer, Header } from 'antd/lib/layout/layout';
import { Flex, Layout } from 'antd';
import Title from 'antd/lib/typography/Title';

import { Logo } from '@self/components/Logo';

import classes from './Layout.module.css';

type RootLayoutProps = {
  children: ReactNode;
};

function RootLayout({ children }: RootLayoutProps) {
  return (
    <Layout className={classes.rootLayout}>
      <Header className={classes.header}>
        <Flex
          className={classes.headerChildren}
          align="center"
          justify="space-between"
        >
          <Flex align="center">
            <Logo width="3em" height="3em" />

            <Title className={classes.appName} level={4}>
              Follytics
            </Title>
          </Flex>
        </Flex>
      </Header>

      <Content className={classes.content}>{children}</Content>

      <Footer className={classes.footer}>
        <p>Follytics ©2025-{new Date().getFullYear()}.</p>
        <p>
          Made with ❤️ by{' '}
          <a href="https://www.abgeo.dev" target="_blank">
            ABGEO
          </a>
        </p>
      </Footer>
    </Layout>
  );
}

export default RootLayout;
