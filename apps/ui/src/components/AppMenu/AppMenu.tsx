'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

import { DashboardOutlined } from '@ant-design/icons';
import { Menu } from 'antd';
import type { MenuProps } from 'antd';

import classes from './AppMenu.module.css';

const createMenuItem = (
  path: string,
  icon: React.ReactNode,
  label: string
) => ({
  key: path,
  icon,
  label: <Link href={path}>{label}</Link>,
});

const menuItems: MenuProps['items'] = [
  createMenuItem('/dashboard', <DashboardOutlined />, 'Dashboard'),
  createMenuItem('/dashboard/widget', <DashboardOutlined />, 'Widget'),
  createMenuItem('/dashboard/profile', <DashboardOutlined />, 'Profile'),
];

function AppMenu() {
  const pathname = usePathname();

  return (
    <>
      <Menu
        theme="dark"
        className={classes.menu}
        mode="inline"
        selectedKeys={[pathname]}
        defaultSelectedKeys={['/dashboard']}
        items={menuItems}
      />
    </>
  );
}

export { AppMenu };
