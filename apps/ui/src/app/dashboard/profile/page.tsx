import Title from 'antd/lib/typography/Title';

import fetchCurrentUser from '@self/data/user/current-user/fetcher';
import getServerApiFactory from '@self/lib/api/server-api-factory';

async function Profile() {
  const apiFactory = await getServerApiFactory();
  const user = await fetchCurrentUser(apiFactory);

  return (
    <>
      <Title level={2}>Profile</Title>
      <pre>{JSON.stringify(user?.data, null, 2)}</pre>
    </>
  );
}

export default Profile;
