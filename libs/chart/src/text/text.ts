import { ChartConfiguration } from '../types';
import { DEFAULTS } from '../constants';
import { addWatermark } from '../components/watermark';
import { createSVGElement } from '../utils/svg';

function createText(
  config: ChartConfiguration,
  text: string
): SVGSVGElement | null {
  const { width = DEFAULTS.width, height = DEFAULTS.height } = config;

  const svg = createSVGElement(config);

  svg
    .append('text')
    .attr('x', (config.width ?? DEFAULTS.width) / 2)
    .attr('y', (config.height ?? DEFAULTS.height) / 2)
    .attr('fill', config.textColor ?? DEFAULTS.textColor)
    .attr('text-anchor', 'middle')
    .text(text)
    .style('font-family', 'Open Sans, sans-serif')
    .style('font-size', '18px');

  addWatermark(svg, config, { right: width / 2 - 50, bottom: height / 2 - 50 });

  return svg.node();
}

export { createText };
