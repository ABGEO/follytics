import * as d3 from 'd3';

import { ChartConfiguration } from '../types';
import { DEFAULTS } from '../constants';
import { SVG } from '../utils/svg';

function addAxis(
  svg: SVG,
  config: ChartConfiguration,
  xScale: d3.ScaleTime<number, number>,
  yScale: d3.ScaleLinear<number, number>,
  margins: { left: number; bottom: number }
) {
  const { width = DEFAULTS.width, height = DEFAULTS.height } = config;

  svg
    .append('g')
    .attr('transform', `translate(0,${height - margins.bottom})`)
    .call(
      d3
        .axisBottom<Date>(xScale)
        .ticks(width / 50)
        .tickFormat(d3.timeFormat('%d %b'))
        .tickSizeOuter(0)
    );

  svg
    .append('g')
    .attr('transform', `translate(${margins.left},0)`)
    .call(d3.axisLeft(yScale).ticks(height / 50));
}

export { addAxis };
