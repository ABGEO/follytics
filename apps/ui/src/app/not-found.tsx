import Link from 'next/link';

import { Button, Flex, Image } from 'antd';
import { HomeOutlined } from '@ant-design/icons';
import Paragraph from 'antd/lib/typography/Paragraph';
import Title from 'antd/lib/typography/Title';

import { MainLayout } from '@self/layout/MainLayout';

import classes from './not-found.module.css';

function NotFound() {
  return (
    <MainLayout>
      <Flex vertical align="center">
        <div>
          <Image
            src="/img/taken.svg"
            preview={false}
            alt="Not Found"
            height={400}
          />
        </div>

        <Flex
          vertical
          align="center"
          className={classes.contentWrapper}
          gap={16}
        >
          <Title className={classes.title}>This Page Was Abducted!</Title>
          <Paragraph className={classes.message}>
            We saw it vanish into the sky. Bright lights, weird sounds, the
            whole deal. Sorry you missed it.
          </Paragraph>

          <div>
            <Link href="/" passHref legacyBehavior>
              <Button type="primary" size="large" icon={<HomeOutlined />}>
                Return to Earth
              </Button>
            </Link>
          </div>
        </Flex>
      </Flex>
    </MainLayout>
  );
}

export default NotFound;
