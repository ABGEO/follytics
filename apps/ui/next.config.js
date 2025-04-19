//@ts-check

const { composePlugins, withNx } = require('@nx/next');

/**
 * @type {import('@nx/next/plugins/with-nx').WithNxOptions}
 **/
const nextConfig = {
  nx: {
    // Set this to true if you would like to use SVGR
    // See: https://github.com/gregberge/svgr
    svgr: false,
  },
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'avatars.githubusercontent.com',
      },
    ],
  },
  output: 'standalone',
  outputFileTracingExcludes: {
    '**/*': [
      // Build tools
      './node_modules/.pnpm/@esbuild+linux-x64@**/**',
      './node_modules/.pnpm/webpack@**/**',

      // SWC binaries
      './node_modules/.pnpm/@swc+core-linux-x64-gnu@**/**',
      './node_modules/.pnpm/@swc+core-linux-x64-musl@**/**',

      // Rspack binaries
      './node_modules/.pnpm/@rspack+binding-linux-x64-gnu@**/**',
      './node_modules/.pnpm/@rspack+binding-linux-x64-musl@**/**',
    ],
  },
  allowedDevOrigins: [
    'localhost:3000',
    '127.0.0.1:3000',
    'follytics.localhost',
  ],
};

const plugins = [
  // Add more Next.js plugins to this list if needed.
  withNx,
];

module.exports = composePlugins(...plugins)(nextConfig);
