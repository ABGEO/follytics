'use client';

import { signOut, useSession } from 'next-auth/react';

import { Avatar, Dropdown, Space, Spin } from 'antd';
import { LogoutOutlined, UserOutlined } from '@ant-design/icons';
import type { ItemType } from 'antd/lib/menu/interface';

import classes from './AccountDropdown.module.css';

const items: ItemType[] = [
  {
    key: 'logout',
    label: 'Logout',
    icon: <LogoutOutlined />,
    onClick: () => signOut({ callbackUrl: '/' }),
    className: classes.logoutItem,
  },
];

function AccountDropdown() {
  const { data: session, status } = useSession();

  if (status === 'loading') {
    return <Spin />;
  }

  return (
    <Dropdown menu={{ items }} trigger={['click']} placement="bottomRight">
      <Space className={classes.userWrapper}>
        <Avatar
          icon={<UserOutlined />}
          src={session?.user?.image}
          alt={session?.user?.name ?? 'Default User Avatar'}
        />
        <span className={classes.userName}>{session?.user?.name}</span>
      </Space>
    </Dropdown>
  );
}

export { AccountDropdown };
