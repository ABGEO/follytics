'use client';

import Image from 'next/image';
import { useState } from 'react';

import { Avatar, Table, Tag } from 'antd';
import type { GetProp, TableProps } from 'antd';
import Link from 'antd/lib/typography/Link';
import { UserOutlined } from '@ant-design/icons';

import type { ResponseEventWithUserReference } from '@follytics/sdk';

import useUserFollowEvents from '@self/data/user/user-follow-events/swr';

type TablePaginationConfig = Exclude<
  GetProp<TableProps, 'pagination'>,
  boolean
>;

type TableParams = {
  pagination?: TablePaginationConfig;
};

type UserFollowEventsProps = {
  userId: string;
};

const columns: TableProps<ResponseEventWithUserReference>['columns'] = [
  {
    title: 'Date',
    dataIndex: 'createdAt',
    key: 'createdAt',
    render: (createdAt) => {
      const date = new Date(createdAt);
      return (
        <>
          {date.toLocaleDateString('en-US', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false,
          })}
        </>
      );
    },
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
    dataIndex: 'user',
    render: (_, { user }) => {
      return (
        <div>
          <Link href={`https://github.com/${user?.username}`} target="_blank">
            <Avatar
              icon={
                !user?.avatar ? (
                  <UserOutlined />
                ) : (
                  <Image
                    src={user?.avatar}
                    alt={user?.name ?? 'Default User Avatar'}
                    sizes="100%"
                    fill
                  />
                )
              }
            />
            <span style={{ marginLeft: 10 }}>@{user?.username}</span>
          </Link>
        </div>
      );
    },
  },
];

function UserFollowEvents({ userId }: UserFollowEventsProps) {
  const [tableParams, setTableParams] = useState<TableParams>({
    pagination: {
      current: 1,
      pageSize: 10,
      responsive: true,
      showLessItems: true,
    },
  });

  const { data, isLoading } = useUserFollowEvents({
    id: userId,
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
    <Table<ResponseEventWithUserReference>
      rowKey="id"
      size="small"
      columns={columns}
      dataSource={data?.data}
      pagination={tableParams.pagination}
      loading={isLoading}
      onChange={handleTableChange}
      scroll={{ x: true }}
    />
  );
}

export { UserFollowEvents };
