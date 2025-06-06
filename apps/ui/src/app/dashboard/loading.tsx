import { Skeleton } from 'antd';

function Loading() {
  return (
    <>
      <Skeleton
        active
        paragraph={false}
        title={{ width: 100 }}
        style={{ marginTop: 20 }}
      />

      <Skeleton active paragraph={{ rows: 16 }} />
    </>
  );
}

export default Loading;
