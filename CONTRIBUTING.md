# Contributing

Thanks for your interest in `torii-sdk-go`!

## Reporting bugs

Open an issue with:

- The version of `torii-sdk-go` you're using (`go list -m github.com/Torii-ApS/torii-sdk-go`).
- A minimal reproduction — a few lines that exhibit the bug.
- What you expected to happen vs. what actually happened.

For security-sensitive issues (anything that could let an attacker forge or bypass token verification), please email **security@torii.so** instead of filing a public issue.

## Development

```sh
git clone https://github.com/Torii-ApS/torii-sdk-go
cd torii-sdk-go
go test ./...
```

The REST client under `internal/generated/` is produced by [`openapi-generator`](https://openapi-generator.tech/) from `spec/server-v1.json`. Don't hand-edit it. Run `./regen.sh` to regenerate it (it encapsulates the steps below). To regenerate after a spec update:

```sh
npx -y @openapitools/openapi-generator-cli generate \
  -i spec/server-v1.json -g go -o internal/generated \
  --additional-properties=packageName=generated,withGoMod=false,isGoSubmodule=false
```

The hand-written surface (`torii.go`, `verify.go`, `authenticate.go`, `middleware/http.go`, `webhook.go`, `types.go`) is where bug reports and PRs typically land.

## Pull requests

1. Open an issue first for non-trivial changes so we can discuss the shape.
2. Branch off `main`, name it `fix/<short>` or `feat/<short>`.
3. Run `gofmt -l .`, `go vet ./...`, `go test ./...` before pushing — CI checks all three.
4. Keep PRs small and focused. One concern per PR.
5. Update `README.md` if you change the public surface.

## Releases

Tagged off `main`. Bump the constant in `torii.go` (`userAgent`) and any references in `README.md`, then:

```sh
git tag v0.0.2
git push origin v0.0.2
```

Consumers pick up the new version via `go get github.com/Torii-ApS/torii-sdk-go@v0.0.2`.

## Code of Conduct

Be kind. Disagreements happen; argue the position, not the person.
