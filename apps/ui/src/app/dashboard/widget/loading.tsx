import { Col, Row, Skeleton } from 'antd';
import Button from 'antd/lib/skeleton/Button';
import { DotChartOutlined } from '@ant-design/icons';
import Input from 'antd/lib/skeleton/Input';
import Node from 'antd/lib/skeleton/Node';

export default function Loading() {
  return (
    <>
      <Skeleton
        active
        paragraph={false}
        title={{ width: 100 }}
        style={{ marginTop: 20 }}
      />

      <Row gutter={[24, 24]}>
        <Col md={24} lg={6}>
          <Skeleton
            active
            paragraph={false}
            title={{ width: 150 }}
            style={{ marginTop: 26 }}
          />

          {[
            { labelWidth: 50 },
            { labelWidth: 50 },
            { labelWidth: 150 },
            { labelWidth: 70 },
            { labelWidth: 70 },
          ].map((config, index) => (
            <div key={index} style={{ marginTop: index === 0 ? 0 : 20 }}>
              <Skeleton
                active
                paragraph={false}
                title={{ width: config.labelWidth }}
              />
              <Input active />
            </div>
          ))}

          <Button active block style={{ width: 100, marginTop: 20 }} />
        </Col>
        <Col md={24} lg={18}>
          <Skeleton
            active
            paragraph={false}
            title={{ width: 90 }}
            style={{ marginTop: 26 }}
          />

          <Node active style={{ width: 960, height: 500 }}>
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
        </Col>
      </Row>
    </>
  );
}
