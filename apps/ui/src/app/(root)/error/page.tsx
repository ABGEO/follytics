'use client';

import { Suspense, useMemo } from 'react';
import Link from 'next/link';
import { useSearchParams } from 'next/navigation';

import { Button, Flex, Image } from 'antd';
import { HomeOutlined } from '@ant-design/icons';
import Paragraph from 'antd/lib/typography/Paragraph';
import Text from 'antd/lib/typography/Text';
import Title from 'antd/lib/typography/Title';

import classes from './error.module.css';

const DEFAULT_ERROR_MESSAGE =
  'Some cosmic turbulence disrupted the connection. Maybe it’s best to return to base.';
const ERROR_MESSAGES = {
  AccessDenied:
    'Access denied by the interstellar gatekeepers. Either you need higher authorization, or this part of the galaxy is restricted. Better head back before you attract unwanted attention.',
  AuthUnknownError:
    'Some mysterious force blocked your access — we’re not sure if it was a cosmic ray or a rogue AI. Either way, something weird happened. Try again or return to base while we decode the disturbance.',
};

function ErrorMessage() {
  const searchParams = useSearchParams();

  const error = searchParams.get('error');

  const errorMessage = useMemo(() => {
    if (error) {
      const errorMessage = ERROR_MESSAGES[error as keyof typeof ERROR_MESSAGES];

      if (errorMessage) {
        return errorMessage;
      }
    }

    return DEFAULT_ERROR_MESSAGE;
  }, [error]);

  return (
    <>
      {error && <Text strong>Error Code: {error}</Text>}
      <Paragraph className={classes.message}>{errorMessage}</Paragraph>
    </>
  );
}

function Error() {
  return (
    <Flex vertical align="center">
      <div>
        <Image
          src="/img/to-the-moon.svg"
          preview={false}
          alt="Error"
          height={300}
        />
      </div>

      <Flex vertical align="center" className={classes.contentWrapper} gap={16}>
        <Title className={classes.title}>Glitch in the Galaxy</Title>

        <Suspense
          fallback={
            <Text>
              <Paragraph className={classes.message}>
                {DEFAULT_ERROR_MESSAGE}
              </Paragraph>
            </Text>
          }
        >
          <ErrorMessage />
        </Suspense>

        <div>
          <Link href="/" passHref>
            <Button type="primary" size="large" icon={<HomeOutlined />}>
              Return to Base
            </Button>
          </Link>
        </div>
      </Flex>
    </Flex>
  );
}

export default Error;
