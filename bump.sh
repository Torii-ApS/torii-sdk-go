#!/usr/bin/env bash
# Go has no version manifest: a module's version IS its git tag, resolved by
# proxy.golang.org. So there is nothing to edit here. This script exists only
# so the release train can call ./bump.sh uniformly across all 7 SDKs.
set -euo pipefail

VERSION="${1:?usage: ./bump.sh <version>  (e.g. 0.0.5)}"
VERSION="${VERSION#v}"
if ! [[ "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+([.-][0-9A-Za-z.]+)?$ ]]; then
	echo "✗ invalid version: '$VERSION'" >&2
	exit 1
fi

echo "✓ torii-sdk-go -> $VERSION (no manifest; the git tag is the release)"
