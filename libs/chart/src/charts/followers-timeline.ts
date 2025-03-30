import * as d3 from 'd3';

import { DEFAULTS } from '../constants';

import { ChartConfiguration, DataPoint, UserDetails } from '../types';
import { addAxis } from '../components/axis';
import { addLine } from '../components/line';
import { addUserDetails } from '../components/user-info';
import { addWatermark } from '../components/watermark';
import { createSVGElement } from '../utils/svg';

function createFollowersTimelineChart(
  data: DataPoint[],
  user: UserDetails,
  config: ChartConfiguration
): SVGSVGElement | null {
  const margins = { top: 20, right: 30, bottom: 30, left: 40 };
  const xScale = d3
    .scaleUtc()
    .domain(d3.extent(data, (d) => d.date) as [Date, Date])
    .range([margins.left, (config.width ?? DEFAULTS.width) - margins.right]);
  const yScale = d3
    .scaleLinear()
    .domain(d3.extent(data, (d) => d.value) as [number, number])
    .range([(config.height ?? DEFAULTS.height) - margins.bottom, margins.top]);

  const svg = createSVGElement(config);
  addAxis(svg, config, xScale, yScale, margins);
  addLine(svg, config, data, xScale, yScale);
  addUserDetails(
    svg,
    config,
    { left: margins.left + 20, top: margins.top + 10 },
    user
  );
  addWatermark(svg, config, {
    right: margins.right,
    bottom: margins.bottom + 20,
  });

  return svg.node();
}

export { createFollowersTimelineChart };
