import * as d3 from 'd3';

import { ChartConfiguration, DataPoint } from '../types';
import { DEFAULTS } from '../constants';
import { SVG } from '../utils/svg';

function addLine(
  svg: SVG,
  config: ChartConfiguration,
  data: DataPoint[],
  xScale: d3.ScaleTime<number, number>,
  yScale: d3.ScaleLinear<number, number>
) {
  svg
    .append('path')
    .datum(data)
    .attr('fill', 'none')
    .attr('stroke', config.lineColor ?? DEFAULTS.lineColor)
    .attr('stroke-width', 1.5)
    .attr(
      'd',
      d3
        .line<DataPoint>()
        .x((d) => xScale(d.date))
        .y((d) => yScale(d.value))
    );
}

export { addLine };
