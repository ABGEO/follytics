import { Flex, Skeleton, Table } from 'antd';
import type { TableProps } from 'antd';

const columns: TableProps<unknown>['columns'] = [
  {
    title: 'Date',
    dataIndex: 'date',
    key: 'date',
    render: () => <Skeleton.Input style={{ width: 135 }} active size="small" />,
  },
  {
    title: 'Type',
    dataIndex: 'type',
    key: 'type',
    render: () => <Skeleton.Button style={{ width: 70 }} active size="small" />,
  },
  {
    title: 'User',
    key: 'user',
    dataIndex: 'user',
    render: () => (
      <Flex align="center" gap={8}>
        <Skeleton.Avatar active size="default" />
        <Skeleton.Input style={{ width: 150 }} active size="small" />
      </Flex>
    ),
  },
];

type UserFollowEventsSkeletonProps = {
  pageSize: number;
};

function UserFollowEventsSkeleton({ pageSize }: UserFollowEventsSkeletonProps) {
  const data = Array.from({ length: pageSize }, (_, index) => ({
    key: `skeleton-${index}`,
    date: null,
    type: null,
    user: null,
  }));

  return (
    <Table
      rowKey="key"
      size="small"
      columns={columns}
      dataSource={data}
      scroll={{ x: true }}
      pagination={{ pageSize }}
    />
  );
}

export { UserFollowEventsSkeleton };
