import { useCallback } from 'react';

import { type ChartConfiguration } from '@follytics/chart';

import { Button, ColorPicker, Form, InputNumber } from 'antd';
import { AggregationColor } from 'antd/lib/color-picker/color';

type ColorConfigProps = {
  backgroundColor: AggregationColor;
  lineColor: AggregationColor;
  axisColor: AggregationColor;
  textColor: AggregationColor;
};

type FormValues = ChartConfiguration & ColorConfigProps;

type WidgetSettingsProps = {
  onChange: (values: Partial<ChartConfiguration>) => void;
  defaultValues?: Partial<FormValues>;
};

const DIMENSIONS = {
  height: { min: 200, max: 1000, default: 500 },
  width: { min: 500, max: 1000, default: 960 },
};

const DEFAULT_COLORS = {
  backgroundColor: '',
  lineColor: '#0d91ce',
  axisColor: '#4B5C6B',
  textColor: '#4B5C6B',
};

const DEFAULT_VALUES = {
  height: DIMENSIONS.height.default,
  width: DIMENSIONS.width.default,
  ...DEFAULT_COLORS,
};

function WidgetSettings({
  onChange,
  defaultValues = DEFAULT_VALUES as FormValues,
}: WidgetSettingsProps) {
  const [form] = Form.useForm<FormValues>();

  const handleReset = useCallback(() => {
    form.resetFields();
    onChange(defaultValues);
  }, [form, onChange, defaultValues]);

  const dimensionControls = [
    {
      name: 'height',
      label: 'Height',
      min: DIMENSIONS.height.min,
      max: DIMENSIONS.height.max,
    },
    {
      name: 'width',
      label: 'Width',
      min: DIMENSIONS.width.min,
      max: DIMENSIONS.width.max,
    },
  ];

  const colorControls = [
    { name: 'backgroundColor', label: 'Background Color', allowClear: true },
    { name: 'lineColor', label: 'Line Color', allowClear: false },
    { name: 'axisColor', label: 'Axis Color', allowClear: false },
    { name: 'textColor', label: 'Text Color', allowClear: false },
  ];

  return (
    <Form
      form={form}
      initialValues={defaultValues}
      onValuesChange={onChange}
      layout="vertical"
    >
      {dimensionControls.map(({ name, label, min, max }) => (
        <Form.Item key={name} label={label} name={name}>
          <InputNumber min={min} max={max} />
        </Form.Item>
      ))}

      {colorControls.map(({ name, label, allowClear }) => (
        <Form.Item key={name} label={label} name={name}>
          <ColorPicker showText allowClear={allowClear} />
        </Form.Item>
      ))}

      <Form.Item>
        <Button htmlType="button" onClick={handleReset}>
          Reset to Default
        </Button>
      </Form.Item>
    </Form>
  );
}

export { WidgetSettings };
