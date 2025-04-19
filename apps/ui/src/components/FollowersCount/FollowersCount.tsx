'use client';

import { ArrowDownOutlined, ArrowUpOutlined } from '@ant-design/icons';
import { Flex, theme } from 'antd';
import Text from 'antd/lib/typography/Text';

import useUserFollowers from '@self/data/user/user-followers/swr';
import useUserFollowersTimeline from '@self/data/user/user-followers-timeline/swr';

const { useToken } = theme;

type FollowersCountProps = {
  userId: string;
};

function FollowersCount({ userId }: FollowersCountProps) {
  const { token } = useToken();

  const { data: timeline } = useUserFollowersTimeline({
    id: userId,
  });

  const { data: followers } = useUserFollowers({
    id: userId,
  });

  // @ts-expect-error we have to fix the API Spec generation
  const count = followers?.pagination?.totalItems ?? 0;
  const timelineData = timeline?.data?.timeline ?? [];
  let change = 0;

  if (timelineData.length > 1) {
    const count = timelineData.length;

    change = timelineData[count - 1].total - timelineData[count - 2].total;
  }

  return (
    <Flex align="center">
      <Text style={{ fontSize: 26, marginRight: 16 }}>{count}</Text>

      {change !== 0 && (
        <Text type={change > 0 ? 'success' : 'danger'} style={{ fontSize: 22 }}>
          {Math.abs(change)}
        </Text>
      )}

      {change > 0 && (
        <ArrowUpOutlined style={{ color: token.colorSuccess, fontSize: 18 }} />
      )}
      {change < 0 && (
        <ArrowDownOutlined style={{ color: token.colorError, fontSize: 18 }} />
      )}
    </Flex>
  );
}

export { FollowersCount };
