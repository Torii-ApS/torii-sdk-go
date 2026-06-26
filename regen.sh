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

# The generator's raw Go output isn't gofmt-clean (e.g. `a,b :=` without the
# space, stray blank lines). Format it so the committed tree is canonical and
# re-running regen is a true no-op (the release gate asserts this).
gofmt -w internal/generated/

echo "✓ regenerated internal/generated/ from spec/server-v1.json"
