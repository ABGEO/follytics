FROM golang:1.23-alpine AS builder

WORKDIR /app

ARG UID=1001

RUN apk update \
  && apk add --no-cache \
  tzdata \
  && update-ca-certificates \
  && adduser \
  --disabled-password \
  --gecos "" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "follytics"

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=bind,source=apps/api/go.mod,target=go.mod \
  --mount=type=bind,source=apps/api/go.sum,target=go.sum \
  go mod download

COPY apps/api .
COPY follytics.example.yaml follytics.yaml

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux go build -v -o follytics cmd/follytics/follytics.go

FROM scratch

COPY --from=arigaio/atlas:0.32.0-alpine /atlas /usr/bin/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder --chown=follytics:follytics /tmp /tmp/

COPY --from=builder /app/follytics /usr/bin/
COPY --from=builder /app/follytics.yaml ./
COPY --from=builder /app/db/migrations /var/migrations/

EXPOSE 8000

USER follytics:follytics

ENTRYPOINT ["follytics"]
CMD ["serve"]
