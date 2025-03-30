import { HttpStatusCode } from 'axios';
import { NextRequest } from 'next/server';

import {
  type DataPoint,
  createFollowersTimelineChart,
  createText,
} from '@follytics/chart';
import {
  type ResponseFollowersTimelineItem,
  type ResponseHTTPResponseResponseFollowersTimeline,
} from '@follytics/sdk';

import { fetchServerData } from '@self/data/server';
import fetchUserFollowersTimeline from '@self/data/user/user-followers-timeline/fetcher';
import getServerApiFactory from '@self/lib/api/server-api-factory';

async function handler(request: NextRequest) {
  const chartConfig = Object.fromEntries(
    request.nextUrl.searchParams.entries()
  );

  try {
    const userId = extractUserIdFromPath(request);
    if (!userId) {
      return createErrorResponse(chartConfig, 'Invalid path');
    }

    const response = await fetchTimelineData(userId);
    if (!response) {
      return createErrorResponse(chartConfig, 'Error fetching data');
    }

    const dataPoints = mapTimelineData(response.data.timeline);
    const svg = createFollowersTimelineChart(
      dataPoints,
      response.data.user,
      chartConfig
    );

    if (!svg) {
      return createErrorResponse(chartConfig, 'Error creating chart');
    }

    return createChartResponse(svg.outerHTML);
  } catch (error) {
    console.error('Unexpected error in widget generation:', error);
    return createErrorResponse(chartConfig, 'Unexpected error occurred');
  }
}

function extractUserIdFromPath(request: NextRequest): string | null {
  const path = request.nextUrl.pathname.split('/');
  if (path.length !== 3 || path[2] === '') {
    return null;
  }

  return path[2];
}

async function fetchTimelineData(
  userId: string
): Promise<ResponseHTTPResponseResponseFollowersTimeline | null> {
  try {
    const apiFactory = await getServerApiFactory();
    const { data, error } = await fetchServerData(
      fetchUserFollowersTimeline,
      apiFactory,
      { id: userId }
    );

    if (error || !data?.data) {
      console.error('Error fetching timeline data:', error);

      return null;
    }

    return data;
  } catch (err) {
    console.error('Exception in fetchTimelineData:', err);

    return null;
  }
}

function mapTimelineData(
  timeline: ResponseFollowersTimelineItem[]
): DataPoint[] {
  return timeline.map((item) => ({
    date: new Date(item.date),
    value: item.total,
  }));
}

function createErrorResponse(config: object, message: string): Response {
  const svg = createText(config, message);

  if (!svg) {
    return new Response(`Error: ${message}`, {
      status: HttpStatusCode.InternalServerError,
      headers: {
        'Content-Type': 'text/plain',
      },
    });
  }

  return new Response(svg.outerHTML, {
    status: HttpStatusCode.Ok,
    headers: {
      'Content-Type': 'image/svg+xml',
      'Cache-Control': 'no-store, no-cache, must-revalidate, proxy-revalidate',
    },
  });
}

function createChartResponse(svgContent: string): Response {
  return new Response(svgContent, {
    status: HttpStatusCode.Ok,
    headers: {
      'Content-Type': 'image/svg+xml',
      'Cache-Control': 'public, max-age=3600',
    },
  });
}

export { handler as GET };
