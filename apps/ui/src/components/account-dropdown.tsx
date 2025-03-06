import { useEffect, useState } from 'react';

import { signOut } from 'next-auth/react';
import { useSession } from 'next-auth/react';

import { Avatar, Dropdown, Space, Spin, theme } from 'antd';
import { LogoutOutlined, UserOutlined } from '@ant-design/icons';
import { red } from '@ant-design/colors';

const items = [
  {
    key: 'logout',
    label: 'Logout',
    icon: <LogoutOutlined />,
    onClick: () => signOut({ callbackUrl: '/' }),
    style: {
      color: red.primary,
    },
  },
];

export function AccountDropdown() {
  const [loading, setLoading] = useState<boolean>(true);

  const { data: session, status } = useSession();
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  useEffect(() => {
    if (status === 'loading') {
      setLoading(false);
    }
  }, [status]);

  if (loading) {
    return <Spin />;
  }

  return (
    <Dropdown menu={{ items }} trigger={['click']} placement="bottomRight">
      <Space style={{ cursor: 'pointer', marginBottom: '0' }}>
        <Avatar
          icon={<UserOutlined />}
          src={session?.user?.image as string}
          alt={session?.user?.name as string}
        />
        <span style={{ color: colorBgContainer }}>{session?.user?.name}</span>
      </Space>
    </Dropdown>
  );
}
