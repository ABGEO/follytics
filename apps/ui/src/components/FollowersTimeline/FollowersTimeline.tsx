'use client';

import { Line } from '@ant-design/charts';

import useUserFollowersTimeline from '@self/data/user/user-followers-timeline/swr';

type FollowersTimelineProps = {
  userId: string;
};

function FollowersTimeline({ userId }: FollowersTimelineProps) {
  const { data: timeline } = useUserFollowersTimeline({
    id: userId,
  });

  const chartData = [];
  for (const element of timeline?.data?.timeline ?? []) {
    chartData.push({
      date: new Date(element.date),
      value: element.total,
      category: 'Total',
    });
    chartData.push({
      date: new Date(element.date),
      value: element.follows,
      category: 'Follows',
    });
    chartData.push({
      date: new Date(element.date),
      value: element.unfollows,
      category: 'Unfollows',
    });
  }

  const config = {
    data: chartData,
    xField: 'date',
    yField: 'value',
    legend: { size: false },
    colorField: 'category',
    scale: { color: { range: ['#FAAD14', '#30BF78', '#F4664A'] } },
  };

  return <Line {...config} />;
}

export { FollowersTimeline };
