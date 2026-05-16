# torii-sdk-go

The official Go backend SDK for [torii](https://torii.so) — verify end-user JWTs without a per-request round trip, manage users from your Go server, and (soon) react to outbound webhooks.

> **Status: 0.0.x preview.** Stable for verify + users + sessions. Outbound webhooks (`VerifyWebhook`) is a stub that returns an error until torii's webhook subsystem ships (tracked in [#424](https://github.com/Torii-ApS/torii/issues/424) Phase 0.5).

## Install

```sh
go get github.com/Torii-ApS/torii-sdk-go
```

Go 1.22+.

## Verify a JWT

```go
import (
    "context"
    torii "github.com/Torii-ApS/torii-sdk-go"
)

auth, err := torii.VerifyToken(ctx, token, torii.VerifyOptions{
    Issuer: "https://acme.torii.so", // or your verified custom domain
})
if err != nil {
    // handle invalid / expired token
}
fmt.Println(auth.UserID, auth.EnvironmentID, auth.EmailVerified)
```

The first call fetches the issuer's JWKS; subsequent calls reuse the cache and rotate keys automatically (handled by [`lestrrat-go/jwx`](https://github.com/lestrrat-go/jwx)). No network round trip per request.

## net/http middleware

```go
import (
    "net/http"
    torii "github.com/Torii-ApS/torii-sdk-go"
    "github.com/Torii-ApS/torii-sdk-go/middleware"
)

auth := middleware.Middleware(middleware.Options{
    Verify: torii.VerifyOptions{Issuer: "https://acme.torii.so"},
})

mux := http.NewServeMux()
mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
    a, _ := middleware.AuthFromContext(r.Context())
    fmt.Fprintf(w, "hello %s", a.UserID)
})

http.ListenAndServe(":8080", auth(mux))
```

`Middleware` writes a 401 by default; override via `Options.OnError`. Skip auth on specific routes (health, metrics, …) via `Options.Skip`.

## Authenticate a single request

```go
auth, err := torii.AuthenticateRequest(ctx, r.Header, torii.VerifyOptions{
    Issuer: "https://acme.torii.so",
})
```

Reads `Authorization: Bearer <token>` and forwards to `VerifyToken`.

## Backend REST API

```go
import (
    "context"
    "os"
    torii "github.com/Torii-ApS/torii-sdk-go"
)

client, err := torii.New(torii.Options{
    SecretKey: os.Getenv("TORII_SECRET_KEY"),
})
if err != nil {
    log.Fatal(err)
}

// List users
page, err := client.Users().List(ctx, torii.ListUsersOptions{
    Limit: ptr.Int32(50),
})

// Create a user
created, err := client.Users().Create(ctx, torii.CreateUserInput{
    Email: ptr.String("x@y.com"),
})

// Ban a user
banned, err := client.Users().Ban(ctx, created.ID)

// Manage sessions
sessions, err := client.Sessions().ListForUser(ctx, created.ID)
err = client.Sessions().RevokeAllForUser(ctx, created.ID)
```

Default base URL is `https://api.torii.so`. Override with `Options.APIURL` for staging or self-hosted.

### Tri-state PATCH semantics

`UpdateUserInput` fields are pointer-to-`PatchString`/`PatchLocale`. This lets callers express three distinct intents:

```go
client.Users().Update(ctx, userID, torii.UpdateUserInput{
    Name:    torii.SetString("Acme Inc"),  // change to "Acme Inc"
    Address: torii.ClearString(),          // clear (JSON null)
    Phone:   nil,                          // leave unchanged (field omitted)
})
```

## Verify outbound webhooks

```go
event, err := torii.VerifyWebhook(secret, r.Header, body) // currently returns an error stub
```

This signature is published so adopting it later won't be a breaking change. The real implementation lands with #424 Phase 0.5.

## License

MIT
