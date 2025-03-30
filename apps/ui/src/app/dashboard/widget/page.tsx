'use client';

import { useCallback, useMemo, useState } from 'react';

import { Col, Row } from 'antd';
import Title from 'antd/lib/typography/Title';

import { type ChartConfiguration } from '@follytics/chart';

import { WidgetPreview } from '@self/components/WidgetPreview';
import { WidgetSettings } from '@self/components/WidgetSettings/WidgetSettings';
import { useAuth } from '@self/providers/AuthProvider';

function Widget() {
  const [widgetConfig, setWidgetConfig] = useState<ChartConfiguration>({});

  const auth = useAuth();

  const onConfigChange = useCallback((values: Partial<ChartConfiguration>) => {
    setWidgetConfig((prev) => ({
      ...prev,
      ...values,
    }));
  }, []);

  const widgetUrl = useMemo(() => {
    const baseUrl = process.env.NEXT_PUBLIC_APP_URL;
    const userId = auth.user?.id;

    if (!userId) return '';

    return `${baseUrl}/widget/${userId}`;
  }, [auth.user?.id]);

  return (
    <>
      <Title level={2}>Widget</Title>

      <Row gutter={[24, 24]}>
        <Col md={24} lg={6}>
          <Title level={3}>Configuration</Title>
          <WidgetSettings onChange={onConfigChange} />
        </Col>
        <Col md={24} lg={18}>
          <Title level={3}>Preview</Title>
          <WidgetPreview url={widgetUrl} config={widgetConfig} />
        </Col>
      </Row>
    </>
  );
}

export default Widget;
