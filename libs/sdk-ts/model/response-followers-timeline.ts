/* tslint:disable */
/* eslint-disable */
/**
 * Follytics API
 * Follytics API service
 *
 * The version of the OpenAPI document: 0.1.0
 * Contact: support@follytics.localhost
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


// May contain unused imports in some cases
// @ts-ignore
import type { ResponseFollowersTimelineItem } from './response-followers-timeline-item';
// May contain unused imports in some cases
// @ts-ignore
import type { ResponseUserForTimeline } from './response-user-for-timeline';

/**
 * 
 * @export
 * @interface ResponseFollowersTimeline
 */
export interface ResponseFollowersTimeline {
    /**
     * 
     * @type {Array<ResponseFollowersTimelineItem>}
     * @memberof ResponseFollowersTimeline
     */
    'timeline': Array<ResponseFollowersTimelineItem>;
    /**
     * 
     * @type {ResponseUserForTimeline}
     * @memberof ResponseFollowersTimeline
     */
    'user': ResponseUserForTimeline;
}

