import { registerOTel } from '@vercel/otel';

export function register() {
  registerOTel({
    serviceName: 'ui',
    propagators: ['auto'],
    traceSampler: 'auto',
    spanProcessors: ['auto'],
  });
}
