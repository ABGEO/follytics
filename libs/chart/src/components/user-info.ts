import { ChartConfiguration, UserDetails } from '../types';
import { DEFAULTS } from '../constants';
import { SVG } from '../utils/svg';

function addUserDetails(
  svg: SVG,
  config: ChartConfiguration,
  margins: { left: number; top: number },
  user: UserDetails
) {
  svg
    .append('text')
    .attr('x', margins.left)
    .attr('y', margins.top)
    .attr('fill', config.textColor ?? DEFAULTS.textColor)
    .attr('text-anchor', 'start')
    .text('Followers timeline for')
    .style('font-family', 'Open Sans, sans-serif')
    .style('font-size', '14px');

  svg
    .append('image')
    .attr('xlink:href', user.avatar)
    .attr('x', margins.left)
    .attr('y', margins.top + 15)
    .attr('width', 30)
    .attr('height', 30);

  svg
    .append('text')
    .attr('x', margins.left + 35)
    .attr('y', margins.top + 35)
    .attr('fill', config.textColor ?? DEFAULTS.textColor)
    .attr('text-anchor', 'start')
    .text(`@${user.username}`)
    .style('font-family', 'Open Sans, sans-serif')
    .style('font-size', '16px');
}

export { addUserDetails };
