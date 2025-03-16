'use client';

import { useEffect, useState } from 'react';
import { usePathname } from 'next/navigation';

import { Breadcrumb, Skeleton } from 'antd';
import type { ItemType } from 'antd/lib/breadcrumb/Breadcrumb';

import classes from './AppBreadcrumb.module.css';

function AppBreadcrumb() {
  const [items, setItems] = useState<ItemType[]>([]);

  const pathname = usePathname();

  const startCase = (string: string) => {
    return string
      .toLowerCase()
      .replace(/[_-]+/g, ' ')
      .replace(/\s+/g, ' ')
      .replace(/([a-z])([A-Z])/g, '$1 $2')
      .trim()
      .split(' ')
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');
  };

  useEffect(() => {
    const defaultItem: ItemType = { title: 'Follytics' };

    const pathSegments = pathname
      .split('/')
      .filter((segment) => segment !== '')
      .map((segment) => ({
        title: startCase(segment),
      }));

    setItems([defaultItem, ...pathSegments]);
  }, [pathname]);

  if (!items || !items.length) {
    return (
      <Skeleton
        className={classes.skeleton}
        paragraph={false}
        title={{ width: 250 }}
      />
    );
  }

  return <Breadcrumb items={items} className={classes.breadcrumb} />;
}

export { AppBreadcrumb };
