#!/usr/bin/env bash
# Regenerate the generated REST client under internal/generated/ from
# spec/server-v1.json. Clears the dir first so stale models from a previous
# spec can't linger, then strips generator scaffolding the SDK doesn't ship
# (test/docs/README/git_push.sh/…). The generated tests in particular import a
# `github.com/GIT_USER_ID/GIT_REPO_ID` placeholder that breaks `go mod tidy`.
set -euo pipefail
cd "$(dirname "$0")"

rm -rf internal/generated
npx -y @openapitools/openapi-generator-cli generate \
  -i spec/server-v1.json -g go -o internal/generated \
  --additional-properties=packageName=generated,withGoMod=false,isGoSubmodule=false

# Strip generator scaffolding the SDK doesn't ship (and that would otherwise
# break the build / go.mod tidy).
( cd internal/generated && rm -rf \
  test docs api .openapi-generator .openapi-generator-ignore .travis.yml \
  git_push.sh README.md .gitignore )

echo "✓ regenerated internal/generated/ from spec/server-v1.json"
