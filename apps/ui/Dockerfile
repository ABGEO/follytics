FROM node:22.16.0-alpine AS base

FROM base AS base-pnpm

# https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine
RUN apk update \
  && apk add --no-cache \
  libc6-compat \
  && corepack enable pnpm

# Install all dependencies for the monorepo.
FROM base-pnpm AS nx-deps

WORKDIR /app

RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
  --mount=type=bind,source=package.json,target=package.json,rw \
  --mount=type=bind,source=pnpm-lock.yaml,target=pnpm-lock.yaml \
  pnpm install --frozen-lockfile --ignore-scripts

# Build the UI application using NX.
FROM base-pnpm AS nx-builder

WORKDIR /app

ENV NEXT_TELEMETRY_DISABLED=1

COPY --from=nx-deps /app/node_modules ./node_modules

COPY nx.json package.json pnpm-lock.yaml tsconfig.base.json ./

COPY apps/ui ./apps/ui/
COPY apps/ui/.docker/.env ./apps/ui/.env

COPY libs/chart/ ./libs/chart/
COPY libs/sdk-ts/ ./libs/sdk-ts/

RUN npx nx build ui

# Prepare the final production image.
FROM base

WORKDIR /app

ARG UID=1001

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV HOSTNAME="0.0.0.0"
ENV PORT=3000

RUN adduser \
  --disabled-password \
  --gecos "" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "follytics"

COPY --from=nx-builder --chown=follytics:follytics /app/dist/apps/ui/.next/standalone ./
COPY --from=nx-builder --chown=follytics:follytics /app/dist/apps/ui/public ./apps/ui/public
COPY --from=nx-builder --chown=follytics:follytics /app/dist/apps/ui/.next/static ./dist/apps/ui/.next/static

COPY apps/ui/.docker/entrypoint.sh /docker-entrypoint.sh

USER follytics:follytics

EXPOSE 3000

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["node", "apps/ui/server.js"]
