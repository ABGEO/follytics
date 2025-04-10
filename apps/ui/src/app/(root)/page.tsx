import Link from 'next/link';

import { Button, Flex, Image, Space } from 'antd';
import { DashboardOutlined, GithubOutlined } from '@ant-design/icons';
import Paragraph from 'antd/lib/typography/Paragraph';
import Title from 'antd/lib/typography/Title';

import { SignIn } from '@self/components/SignIn';
import { auth } from '@self/lib/auth';

async function Home() {
  const session = await auth();

  return (
    <Flex vertical align="center" justify="center">
      <Title>Track your GitHub follower trends over time</Title>

      <Paragraph style={{ maxWidth: 600, marginTop: 32 }}>
        Follytics is an open-source analytics tool that tracks every follow and
        unfollow as an event. Built with an event-sourcing-like pattern to give
        you full historical insight into your follower activity.
      </Paragraph>

      <Image
        src="/img/growing.svg"
        preview={false}
        alt="Followers Chart Widget"
        style={{ marginTop: 64 }}
      />

      <Space size="large" style={{ marginTop: 32 }}>
        {!session && <SignIn />}
        {session && (
          <Link href="/dashboard" passHref>
            <Button type="primary" size="large" icon={<DashboardOutlined />}>
              Dashboard
            </Button>
          </Link>
        )}

        <Button
          size="large"
          icon={<GithubOutlined />}
          href="https://github.com/ABGEO/follytics"
          target="_blank"
        >
          GitHub
        </Button>
      </Space>
    </Flex>
  );
}

export default Home;
