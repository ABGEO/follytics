'use client';

import { useState } from 'react';

import { Avatar, Table, Tag } from 'antd';
import type { GetProp, TableProps } from 'antd';

import type { ResponseEventWithUserReference } from '@follytics/sdk';

import { ErrorBoundary } from '@self/components/ErrorBoundary';
import { useAuth } from '@self/providers/AuthProvider';
import useUserFollowEvents from '@self/data/user/user-follow-events/swr';

type TablePaginationConfig = Exclude<
  GetProp<TableProps, 'pagination'>,
  boolean
>;

type TableParams = {
  pagination?: TablePaginationConfig;
};

const columns: TableProps<ResponseEventWithUserReference>['columns'] = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    render: (text) => <a>{text}</a>,
  },
  {
    title: 'Created At',
    dataIndex: 'createdAt',
    key: 'createdAt',
  },
  {
    title: 'Type',
    dataIndex: 'type',
    key: 'type',
    render: (type) => {
      let color = 'default';
      if (type === 'FOLLOW') {
        color = 'success';
      }

      if (type === 'UNFOLLOW') {
        color = 'error';
      }

      return <Tag color={color}>{type}</Tag>;
    },
  },

  {
    title: 'User',
    key: 'user',
    dataIndex: 'userName',
    render: (_, { user }) => <>{user?.username}</>,
  },
  {
    title: 'Avatar',
    key: 'user',
    dataIndex: 'userAvatar',
    render: (_, { user }) => (
      <>
        <Avatar src={user?.avatar} />
      </>
    ),
  },
];

function UserFollowEvents() {
  const [tableParams, setTableParams] = useState<TableParams>({
    pagination: {
      current: 1,
      pageSize: 10,
      responsive: true,
      showLessItems: true,
    },
  });

  const auth = useAuth();
  const { data, error, isLoading } = useUserFollowEvents({
    id: auth.user?.id ?? '',
    page: tableParams.pagination?.current,
    limit: tableParams.pagination?.pageSize,
  });

  const handleTableChange: TableProps<ResponseEventWithUserReference>['onChange'] =
    (pagination) => {
      setTableParams({
        pagination,
      });
    };

  // @ts-expect-error we have to fix the API Spec generation
  if (data && data.pagination.totalItems !== tableParams.pagination?.total) {
    setTableParams((prevState) => ({
      ...prevState,
      pagination: {
        ...prevState.pagination,
        // @ts-expect-error we have to fix the API Spec generation
        total: data.pagination.totalItems,
      },
    }));
  }

  return (
    <ErrorBoundary error={error}>
      <Table<ResponseEventWithUserReference>
        rowKey="id"
        columns={columns}
        dataSource={data?.data}
        pagination={tableParams.pagination}
        loading={isLoading}
        onChange={handleTableChange}
        scroll={{ x: true }}
      />
    </ErrorBoundary>
  );
}

export { UserFollowEvents };
