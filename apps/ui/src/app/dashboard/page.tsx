import Title from 'antd/lib/typography/Title';

import { UserFollowEvents } from '@self/components/UserFollowEvents';

function Dashboard() {
  return (
    <>
      <Title level={2}>Follow Events</Title>
      <UserFollowEvents />
    </>
  );
}

export default Dashboard;
