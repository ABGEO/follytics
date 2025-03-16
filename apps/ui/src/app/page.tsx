import Link from 'next/link';

import { Button, Flex } from 'antd';
import Title from 'antd/lib/typography/Title';

import { SignIn } from '@self/components/SignIn';
import { auth } from '@self/lib/auth';

async function Home() {
  const session = await auth();

  return (
    <Flex align="center" vertical>
      <Title>Welcome to Follytics</Title>

      {!session && <SignIn />}
      {session && (
        <>
          <pre>{JSON.stringify(session, null, 2)}</pre>
          <Link href="/dashboard" passHref legacyBehavior>
            <Button>Dashboard</Button>
          </Link>
        </>
      )}
    </Flex>
  );
}

export default Home;
