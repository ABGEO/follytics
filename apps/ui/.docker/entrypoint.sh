#!/usr/bin/env sh

set -e

printenv | grep NEXT_PUBLIC_ | while read -r ENV_LINE ; do
  ENV_KEY=$(echo $ENV_LINE | cut -d "=" -f1)
  ENV_VALUE=$(echo $ENV_LINE | cut -d "=" -f2)

  find /app/dist/apps/ui/.next -type f -exec sed -i "s|_${ENV_KEY}_|${ENV_VALUE}|g" {} \;
done

exec "$@"
