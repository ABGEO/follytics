'use client';

import { Suspense } from 'react';

import { Card, Col, Row } from 'antd';
import Title from 'antd/lib/typography/Title';

import { FollowersCount } from '@self/components/FollowersCount';
import { FollowersTimeline } from '@self/components/FollowersTimeline';
import { NewFollowers } from '@self/components/NewFollowers';
import { UserFollowEvents } from '@self/components/UserFollowEvents';
import { useAuth } from '@self/providers/AuthProvider';

function Dashboard() {
  const auth = useAuth();

  return (
    <>
      <Title level={2}>Dashboard</Title>

      <Row gutter={[24, 24]}>
        <Col xs={24} sm={10} md={8} lg={4}>
          <Card size="small" title="Followers" variant="borderless">
            <Suspense fallback={<div>Loading...</div>}>
              {auth.user?.id && <FollowersCount userId={auth.user?.id} />}
            </Suspense>
          </Card>
        </Col>

        <Col xs={24} sm={14} md={16} lg={6}>
          <Card size="small" title="New Followers" variant="borderless">
            <Suspense fallback={<div>Loading...</div>}>
              {auth.user?.id && (
                <NewFollowers userId={auth.user?.id} limit={9} />
              )}
            </Suspense>
          </Card>
        </Col>

        <Col md={24} lg={14}>
          <Card size="small" title="Follow Events" variant="borderless">
            <Suspense fallback={<div>Loading...</div>}>
              {auth.user?.id && <UserFollowEvents userId={auth.user?.id} />}
            </Suspense>
          </Card>
        </Col>
      </Row>

      <Row>
        <Col span={24}>
          <Card size="small" title="Followers Timeline" variant="borderless">
            <Suspense fallback={<div>Loading...</div>}>
              {auth.user?.id && <FollowersTimeline userId={auth.user?.id} />}
            </Suspense>
          </Card>
        </Col>
      </Row>
    </>
  );
}

export default Dashboard;
