'use client';

import { useSession } from 'next-auth/react';

import { Flex, Typography } from 'antd';

import { SignIn } from '@self/components/sign-in';

const { Title } = Typography;

export default function Home() {
  const { status } = useSession();

  return (
    <Flex align="center" vertical>
      <Title>Welcome to Follytics</Title>

      {status !== 'authenticated' && <SignIn />}
    </Flex>
  );
}
