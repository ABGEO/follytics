import Title from 'antd/lib/typography/Title';

import { ErrorBoundary } from '@self/components/ErrorBoundary';
import { fetchCurrentUser } from '@self/data/user/current-user/fetcher';
import { fetchServerData } from '@self/data/server';

async function Profile() {
  const { data: user, error } = await fetchServerData(fetchCurrentUser);

  return (
    <>
      <Title level={2}>Profile</Title>
      <ErrorBoundary error={error}>
        <pre>{JSON.stringify(user?.data, null, 2)}</pre>
      </ErrorBoundary>
    </>
  );
}

export default Profile;
