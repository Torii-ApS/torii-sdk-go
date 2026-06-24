#!/usr/bin/env bash
# Regenerate the generated REST client under internal/generated/ from
# spec/server-v1.json. Idempotent; safe to re-run after a spec bump.
set -euo pipefail
cd "$(dirname "$0")"

npx -y @openapitools/openapi-generator-cli generate \
  -i spec/server-v1.json -g go -o internal/generated \
  --additional-properties=packageName=generated,withGoMod=false,isGoSubmodule=false

echo "✓ regenerated internal/generated/ from spec/server-v1.json"
