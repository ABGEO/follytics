import { DotChartOutlined } from '@ant-design/icons';
import Node from 'antd/lib/skeleton/Node';
import { Skeleton } from 'antd';

type WidgetPreviewSkeletonProps = {
  width?: number;
  height?: number;
};

function WidgetPreviewSkeleton({
  width = 960,
  height = 500,
}: WidgetPreviewSkeletonProps) {
  return (
    <div>
      <Node active style={{ width, height }}>
        <DotChartOutlined
          style={{
            fontSize: 128,
            color: '#bfbfbf',
          }}
        />
      </Node>

      <Skeleton
        active
        paragraph={false}
        title={{ width: 125 }}
        style={{ marginTop: 20 }}
      />

      <Skeleton
        active
        paragraph={false}
        title={{ width: 35 }}
        style={{ marginTop: 8, marginBottom: 0 }}
      />

      <Skeleton
        active
        paragraph={false}
        title={{ width: 500 }}
        style={{ marginTop: -15 }}
      />
    </div>
  );
}

export { WidgetPreviewSkeleton };
