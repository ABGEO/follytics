import { Avatar, List, Skeleton } from 'antd';

type UserFollowEventsSkeletonProps = {
  size: number;
};

function NewFollowersSkeleton({ size }: UserFollowEventsSkeletonProps) {
  return (
    <List
      dataSource={[...Array(size).keys()]}
      renderItem={(_) => (
        <List.Item>
          <List.Item.Meta
            avatar={<Avatar icon={<Skeleton.Avatar active size="default" />} />}
            title={
              <Skeleton.Input style={{ width: 150 }} active size="small" />
            }
          />
        </List.Item>
      )}
    />
  );
}

export { NewFollowersSkeleton };
