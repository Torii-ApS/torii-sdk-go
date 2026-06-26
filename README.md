# torii-sdk-go

The official Go backend SDK for [torii](https://torii.so) — verify end-user JWTs without a per-request round trip and manage users from your Go server.

> **v0.x — API may still change.**

## Setup

1. Sign in to [app.torii.so](https://app.torii.so) and from your dashboard copy:
   - your **issuer URL** (e.g. `https://acme.torii.so`)
   - a **secret key** (`sk_test_…` for development, `sk_live_…` for production)

2. Install the SDK:

   ```sh
   go get github.com/Torii-ApS/torii-sdk-go
   ```

   Go 1.22+.

3. Verify an end-user JWT:

   ```go
   import (
       "context"
       torii "github.com/Torii-ApS/torii-sdk-go"
   )

   auth, err := torii.VerifyToken(ctx, token, torii.VerifyOptions{
       Issuer: "https://acme.torii.so",
   })
   if err != nil {
       // handle invalid / expired token
   }
   fmt.Println(auth.UserID, auth.EnvironmentID, auth.EmailVerified)
   ```

   The first call fetches the issuer's JWKS; subsequent calls reuse the cache and rotate keys automatically (handled by [`lestrrat-go/jwx`](https://github.com/lestrrat-go/jwx)). No round trip per request.

4. Call the backend REST API:

   ```go
   client, err := torii.New(torii.Options{
       SecretKey: os.Getenv("TORII_SECRET_KEY"),
   })
   if err != nil {
       log.Fatal(err)
   }

   user, err := client.Users().Get(ctx, userID)
   ```

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

### Tri-state PATCH semantics

`UpdateUserInput` fields use the generic `Patch[T any]` wrapper. This lets callers express three distinct intents:

```go
client.Users().Update(ctx, userID, torii.UpdateUserInput{
    Name:    torii.SetPatch("Acme Inc"),     // change to "Acme Inc"
    Address: torii.ClearPatch[string](),     // clear (JSON null)
    // Phone is the zero value — leave unchanged (field omitted from request)
})
```

## License

MIT
