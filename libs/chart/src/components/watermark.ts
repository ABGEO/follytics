import { ChartConfiguration } from '../types';
import { DEFAULTS } from '../constants';
import { SVG } from '../utils/svg';

const watermarkImage =
  'data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPHN2ZyBpZD0iTGF5ZXJfMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB2ZXJzaW9uPSIxLjEiIHZpZXdCb3g9IjAgMCAyMDAgMjAwIj4KICA8IS0tIEdlbmVyYXRvcjogQWRvYmUgSWxsdXN0cmF0b3IgMjkuMy4xLCBTVkcgRXhwb3J0IFBsdWctSW4gLiBTVkcgVmVyc2lvbjogMi4xLjAgQnVpbGQgMTUxKSAgLS0+CiAgPGRlZnM+CiAgICA8c3R5bGU+CiAgICAgIC5zdDAgewogICAgICAgIGZpbGw6ICMwMjMwNDc7CiAgICAgIH0KCiAgICAgIC5zdDEgewogICAgICAgIGZpbGw6ICM4ZWNhZTY7CiAgICAgIH0KCiAgICAgIC5zdDIgewogICAgICAgIGZpbGw6ICNmNWY3ZmE7CiAgICAgIH0KCiAgICAgIC5zdDMgewogICAgICAgIGZpbGw6ICMyMTllYmM7CiAgICAgIH0KICAgIDwvc3R5bGU+CiAgPC9kZWZzPgogIDxnIGlkPSJCYWNrZ3JvdW5kIj4KICAgIDxjaXJjbGUgY2xhc3M9InN0MiIgY3g9IjEwMCIgY3k9IjEwMCIgcj0iOTAiLz4KICA8L2c+CiAgPGcgaWQ9IkxvZ28iPgogICAgPHBhdGggY2xhc3M9InN0MyIgZD0iTTUwLDg1aDcxYzIuMiwwLDQsMS44LDQsNHYyMmMwLDIuMi0xLjgsNC00LDRINTB2LTMwaDBaIi8+CiAgICA8cGF0aCBjbGFzcz0ic3QxIiBkPSJNNTAsNTBoOTZjMi4yLDAsNCwxLjgsNCw0djIyYzAsMi4yLTEuOCw0LTQsNEg1MHYtMzBoMFoiLz4KICAgIDxwYXRoIGNsYXNzPSJzdDAiIGQ9Ik01MCwxMjBoMzFjMi4yLDAsNCwxLjgsNCw0djIyYzAsMi4yLTEuOCw0LTQsNGgtMzF2LTMwaDBaIi8+CiAgPC9nPgo8L3N2Zz4=';

function addWatermark(
  svg: SVG,
  config: ChartConfiguration,
  margins: { right: number; bottom: number }
) {
  const { width = DEFAULTS.width, height = DEFAULTS.height } = config;

  svg
    .append('image')
    .attr('xlink:href', watermarkImage)
    .attr('x', width - margins.right - 115)
    .attr('y', height - margins.bottom - 20)
    .attr('width', 30)
    .attr('height', 30);

  svg
    .append('text')
    .attr('x', width - margins.right)
    .attr('y', height - margins.bottom)
    .attr('fill', config.textColor ?? DEFAULTS.textColor)
    .attr('text-anchor', 'end')
    .text('follytics.app')
    .style('font-family', 'Open Sans, sans-serif')
    .style('font-size', '14px');
}

export { watermarkImage, addWatermark };
