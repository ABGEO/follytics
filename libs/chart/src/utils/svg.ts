import * as d3 from 'd3';
import { JSDOM } from 'jsdom';

import { DEFAULTS, FONT_FAMILY } from '../constants';
import { ChartConfiguration } from '../types';

type SVG = d3.Selection<SVGSVGElement, unknown, null, undefined>;

function createSVGElement(config: ChartConfiguration): SVG {
  const {
    width = DEFAULTS.width,
    height = DEFAULTS.height,
    backgroundColor,
  } = config;
  const dom = new JSDOM(`<!DOCTYPE html><body></body>`);
  const body = d3.select(dom.window.document.querySelector('body'));

  const svg = body
    .append('svg')
    .attr('xmlns', 'http://www.w3.org/2000/svg')
    .attr(
      'style',
      `max-width: 100%; height: auto; height: intrinsic; background-color: ${
        backgroundColor ?? 'transparent'
      }; font-family: ${FONT_FAMILY};`,
    )
    .attr('width', width)
    .attr('height', height)
    .attr('viewBox', [0, 0, width, height]);

  svg
    .append('style')
    .attr('type', 'text/css')
    .text(
      `
      path.domain,
      .tick line {
        stroke: ${config.axisColor ?? DEFAULTS.textColor};
      }
      .tick text {
        fill: ${config.textColor ?? DEFAULTS.textColor};
        font-family: ${FONT_FAMILY};
      }
      `,
    );

  return svg;
}

export { type SVG, createSVGElement };
