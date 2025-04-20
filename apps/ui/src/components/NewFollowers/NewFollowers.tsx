'use client';

import Image from 'next/image';

import { Avatar, List } from 'antd';
import { UserOutlined } from '@ant-design/icons';

import Link from 'antd/lib/typography/Link';

import useUserFollowEvents from '@self/data/user/user-follow-events/swr';

import { NewFollowersSkeleton } from './Skeleton';

type NewFollowersProps = {
  userId: string;
  limit?: number;
};

function NewFollowers({ userId, limit = 10 }: NewFollowersProps) {
  const { data: events, isLoading } = useUserFollowEvents({
    id: userId,
    limit: limit,
    filter: ['type||eq||FOLLOW'],
  });

  const data = events?.data?.map((event) => {
    return {
      username: event.user?.username,
      avatar: event.user?.avatar,
    };
  });

  if (isLoading) {
    return <NewFollowersSkeleton size={limit} />;
  }

  return (
    <List
      dataSource={data}
      renderItem={(user) => (
        <List.Item>
          <List.Item.Meta
            avatar={
              <Avatar
                icon={
                  !user?.avatar ? (
                    <UserOutlined />
                  ) : (
                    <Image
                      src={user?.avatar}
                      alt={user?.username ?? 'Default User Avatar'}
                      sizes="100%"
                      fill
                    />
                  )
                }
              />
            }
            title={
              <Link
                href={`https://github.com/${user?.username}`}
                target="_blank"
              >
                @{user.username}
              </Link>
            }
          />
        </List.Item>
      )}
    />
  );
}

export { NewFollowers };
