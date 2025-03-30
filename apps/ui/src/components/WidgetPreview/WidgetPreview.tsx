import { useCallback, useMemo } from 'react';

import { Image, Typography } from 'antd';
import { AggregationColor } from 'antd/lib/color-picker/color';
import Title from 'antd/lib/typography/Title';

const { Paragraph } = Typography;

type WidgetPreviewProps = {
  url: string;
  config?: object;
  altText?: string;
};

function WidgetPreview({
  url,
  config,
  altText = 'Check my followers using Follytics.app',
}: WidgetPreviewProps) {
  const formatConfigValue = useCallback((value: unknown): string => {
    if (
      value &&
      typeof value === 'object' &&
      'toHexString' in value &&
      typeof (value as AggregationColor).toHexString === 'function'
    ) {
      return (value as AggregationColor).toHexString();
    }

    return String(value);
  }, []);

  const queryString = useMemo(() => {
    if (!config || Object.keys(config).length === 0) return '';

    const serializableConfig = Object.entries(config).reduce(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = formatConfigValue(value);
        }

        return acc;
      },
      {} as Record<string, string>
    );

    const params = new URLSearchParams(serializableConfig);

    return `?${params}`;
  }, [config, formatConfigValue]);

  const imageUrl = useMemo(() => `${url}${queryString}`, [url, queryString]);
  const embedOptions = useMemo(
    () => [
      { title: 'URL', content: imageUrl },
      { title: 'Markdown', content: `![${altText}](${imageUrl})` },
      { title: 'HTML', content: `<img alt="${altText}" src="${imageUrl}" />` },
    ],
    [imageUrl, altText]
  );

  return (
    <>
      <div>
        <Image src={imageUrl} alt={altText} preview={false} />
      </div>

      <div>
        <Title level={3}>Embedding</Title>

        {embedOptions.map((option) => (
          <div key={option.title}>
            <Title level={4}>{option.title}</Title>
            <Paragraph copyable>{option.content}</Paragraph>
          </div>
        ))}
      </div>
    </>
  );
}

export { WidgetPreview };
