package torii

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Torii-ApS/torii-sdk-go/internal/generated"
)

// defaultAPIURL is the production base URL for torii's backend SDK API.
const defaultAPIURL = "https://api.torii.so"

// userAgent advertises the SDK in outgoing requests. Bumped manually with releases.
const userAgent = "torii-sdk-go/0.0.2"

// Client is the top-level entrypoint to the torii backend API.
// Construct with New and reuse a single instance for the lifetime of your
// process — Client is safe for concurrent use.
type Client struct {
	cfg   Options
	api   *generated.APIClient
	users *usersClient
	sess  *sessionsClient
}

// Users returns the resource client for /users endpoints.
func (c *Client) Users() Users { return c.users }

// Sessions returns the resource client for /sessions endpoints.
func (c *Client) Sessions() Sessions { return c.sess }

// New constructs a *Client from the given Options.
// Returns an error if SecretKey is empty or APIURL is invalid.
func New(opts Options) (*Client, error) {
	if opts.SecretKey == "" {
		return nil, newError("torii.New: Options.SecretKey is required", nil)
	}
	apiURL := opts.APIURL
	if apiURL == "" {
		apiURL = defaultAPIURL
	}
	u, err := url.Parse(apiURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, newError(fmt.Sprintf("torii.New: invalid APIURL %q", apiURL), err)
	}

	cfg := generated.NewConfiguration()
	cfg.Servers = generated.ServerConfigurations{{URL: apiURL, Description: "torii API"}}
	cfg.UserAgent = userAgent
	cfg.AddDefaultHeader("Authorization", "Bearer "+opts.SecretKey)
	if opts.HTTPClient != nil {
		cfg.HTTPClient = opts.HTTPClient
	}

	api := generated.NewAPIClient(cfg)
	c := &Client{cfg: opts, api: api}
	c.users = &usersClient{api: api}
	c.sess = &sessionsClient{api: api}
	return c, nil
}

// CursorPage is a single page of cursor-paginated results.
type CursorPage[T any] struct {
	Items      []T
	NextCursor *string
	HasMore    bool
}

// UserStatus enumerates server-side user lifecycle states.
type UserStatus string

const (
	UserStatusPendingVerification UserStatus = "pending_verification"
	UserStatusActive              UserStatus = "active"
	UserStatusBanned              UserStatus = "banned"
	UserStatusDeleted             UserStatus = "deleted"
)

// Locale is the end-user's preferred display language.
type Locale string

const (
	LocaleEN Locale = "en"
	LocaleDA Locale = "da"
)

// User represents a torii end-user as returned by the backend API.
// Nullable fields use pointer types so callers can distinguish "not present"
// (nil) from "present and empty" (*string == "").
type User struct {
	ID            string
	EnvironmentID string
	Name          *string
	Email         *string
	Phone         *string
	Locale        *Locale
	Address       *string
	DateOfBirth   *string // ISO-8601 date (YYYY-MM-DD)
	Status        UserStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

// Session represents an active end-user session for a given user.
type Session struct {
	ID            string
	UserID        string
	EnvironmentID string
	UserAgent     *string
	IPAddress     *string
	CreatedAt     time.Time
	ExpiresAt     time.Time
	LastUsedAt    time.Time
}

// CreateUserInput is the request body for Users.Create.
type CreateUserInput struct {
	Email       *string
	Name        *string
	Phone       *string
	Password    *string
	Address     *string
	DateOfBirth *string // ISO-8601 date (YYYY-MM-DD)
	// EmailVerified, Locale, CustomClaims are accepted by the surface contract
	// but not yet wired in the OpenAPI spec we generate from. They're declared
	// here so a future spec update can populate them without breaking callers.
	EmailVerified *bool
	Locale        *Locale
	CustomClaims  map[string]any
}

// ListUsersOptions controls the search payload for Users.List.
type ListUsersOptions struct {
	Limit         *int32
	Cursor        *string
	Name          *string
	Email         *string
	Statuses      []UserStatus
	CreatedAfter  *time.Time
	CreatedBefore *time.Time
}

// Users is the resource interface for /users endpoints.
type Users interface {
	List(ctx context.Context, opts ListUsersOptions) (CursorPage[User], error)
	Get(ctx context.Context, userID string) (*User, error)
	Create(ctx context.Context, in CreateUserInput) (*User, error)
	Update(ctx context.Context, userID string, in UpdateUserInput) (*User, error)
	Delete(ctx context.Context, userID string) error
	Ban(ctx context.Context, userID string) (*User, error)
	Unban(ctx context.Context, userID string) (*User, error)
}

// Sessions is the resource interface for /sessions endpoints.
type Sessions interface {
	ListForUser(ctx context.Context, userID string) ([]Session, error)
	RevokeAllForUser(ctx context.Context, userID string) error
	Revoke(ctx context.Context, userID, sessionID string) error
}

// --- usersClient -------------------------------------------------------------

type usersClient struct{ api *generated.APIClient }

func (c *usersClient) List(ctx context.Context, opts ListUsersOptions) (CursorPage[User], error) {
	req := c.api.ServerUsersAPI.SearchUsers(ctx)
	if opts.Limit != nil {
		req = req.Limit(*opts.Limit)
	}
	if opts.Cursor != nil {
		req = req.Cursor(*opts.Cursor)
	}
	body := generated.NewServerUserSearchRequest()
	if opts.Name != nil {
		body.SetName(*opts.Name)
	}
	if opts.Email != nil {
		body.SetEmail(*opts.Email)
	}
	if len(opts.Statuses) > 0 {
		strs := make([]string, len(opts.Statuses))
		for i, s := range opts.Statuses {
			strs[i] = string(s)
		}
		body.Statuses = strs
	}
	if opts.CreatedAfter != nil {
		body.SetCreatedAfter(*opts.CreatedAfter)
	}
	if opts.CreatedBefore != nil {
		body.SetCreatedBefore(*opts.CreatedBefore)
	}
	req = req.ServerUserSearchRequest(*body)
	res, httpRes, err := req.Execute()
	if err := wrapAPIError(httpRes, err); err != nil {
		return CursorPage[User]{}, err
	}
	page := CursorPage[User]{
		Items:   make([]User, 0, len(res.Items)),
		HasMore: res.HasMore,
	}
	for i := range res.Items {
		page.Items = append(page.Items, userFromGenerated(&res.Items[i]))
	}
	if res.NextCursor.IsSet() && res.NextCursor.Get() != nil {
		s := *res.NextCursor.Get()
		page.NextCursor = &s
	}
	return page, nil
}

func (c *usersClient) Get(ctx context.Context, userID string) (*User, error) {
	res, httpRes, err := c.api.ServerUsersAPI.GetUser(ctx, userID).Execute()
	if err := wrapAPIError(httpRes, err); err != nil {
		return nil, err
	}
	u := userFromGenerated(res)
	return &u, nil
}

func (c *usersClient) Create(ctx context.Context, in CreateUserInput) (*User, error) {
	body := generated.NewCreateUserRequest()
	if in.Email != nil {
		body.SetEmail(*in.Email)
	}
	if in.Password != nil {
		body.SetPassword(*in.Password)
	}
	if in.Name != nil {
		body.SetName(*in.Name)
	}
	if in.Phone != nil {
		body.SetPhone(*in.Phone)
	}
	if in.Address != nil {
		body.SetAddress(*in.Address)
	}
	if in.DateOfBirth != nil {
		body.SetDateOfBirth(*in.DateOfBirth)
	}
	res, httpRes, err := c.api.ServerUsersAPI.CreateUser(ctx).CreateUserRequest(*body).Execute()
	if err := wrapAPIError(httpRes, err); err != nil {
		return nil, err
	}
	u := userFromGenerated(res)
	return &u, nil
}

func (c *usersClient) Update(ctx context.Context, userID string, in UpdateUserInput) (*User, error) {
	// The generated UpdateUserRequest uses `*string` fields (omitempty), which
	// cannot express "send JSON null to clear this field". We sidestep the
	// generated body marshalling by serialising the tri-state Patch[T]
	// wrappers ourselves and PATCHing the bytes directly.
	body, err := in.asJSONBody()
	if err != nil {
		return nil, newError("torii.Users.Update: encode body", err)
	}
	var out generated.UserResponse
	if err := c.doJSON(ctx, http.MethodPatch,
		"/api/server/v1/users/"+url.PathEscape(userID), body, &out); err != nil {
		return nil, err
	}
	u := userFromGenerated(&out)
	return &u, nil
}

// doJSON sends a JSON request body via the generated APIClient's
// configured *http.Client and decodes the response into out (when non-nil
// and the response is 2xx). Non-2xx responses are returned as *APIError.
func (c *usersClient) doJSON(ctx context.Context, method, path string, body []byte, out any) error {
	cfg := c.api.GetConfig()
	if len(cfg.Servers) == 0 {
		return newError("torii: no server URL configured", nil)
	}
	base := cfg.Servers[0].URL
	req, err := http.NewRequestWithContext(ctx, method, base+path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	for k, v := range cfg.DefaultHeader {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if cfg.UserAgent != "" {
		req.Header.Set("User-Agent", cfg.UserAgent)
	}
	resp, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, rerr := io.ReadAll(resp.Body)
	if rerr != nil {
		return rerr
	}
	if resp.StatusCode >= 300 {
		apiErr := &APIError{Status: resp.StatusCode, Message: resp.Status, Body: respBody}
		var parsed struct {
			Code      string `json:"code"`
			SupportID string `json:"supportId"`
			Message   string `json:"message"`
		}
		if jerr := json.Unmarshal(respBody, &parsed); jerr == nil {
			if parsed.Code != "" {
				apiErr.Code = parsed.Code
			}
			if parsed.SupportID != "" {
				apiErr.SupportID = parsed.SupportID
			}
			if parsed.Message != "" {
				apiErr.Message = parsed.Message
			}
		}
		return apiErr
	}
	if out != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, out); err != nil {
			return newError("torii: decode response", err)
		}
	}
	return nil
}

func (c *usersClient) Delete(ctx context.Context, userID string) error {
	httpRes, err := c.api.ServerUsersAPI.DeleteUser(ctx, userID).Execute()
	return wrapAPIError(httpRes, err)
}

func (c *usersClient) Ban(ctx context.Context, userID string) (*User, error) {
	res, httpRes, err := c.api.ServerUsersAPI.BanUser(ctx, userID).Execute()
	if err := wrapAPIError(httpRes, err); err != nil {
		return nil, err
	}
	u := userFromGenerated(res)
	return &u, nil
}

func (c *usersClient) Unban(ctx context.Context, userID string) (*User, error) {
	res, httpRes, err := c.api.ServerUsersAPI.UnbanUser(ctx, userID).Execute()
	if err := wrapAPIError(httpRes, err); err != nil {
		return nil, err
	}
	u := userFromGenerated(res)
	return &u, nil
}

// --- sessionsClient ----------------------------------------------------------

type sessionsClient struct{ api *generated.APIClient }

func (c *sessionsClient) ListForUser(ctx context.Context, userID string) ([]Session, error) {
	res, httpRes, err := c.api.ServerSessionsAPI.ListSessions(ctx, userID).Execute()
	if err := wrapAPIError(httpRes, err); err != nil {
		return nil, err
	}
	out := make([]Session, 0, len(res))
	for i := range res {
		out = append(out, sessionFromGenerated(&res[i]))
	}
	return out, nil
}

func (c *sessionsClient) RevokeAllForUser(ctx context.Context, userID string) error {
	httpRes, err := c.api.ServerSessionsAPI.RevokeAllSessions(ctx, userID).Execute()
	return wrapAPIError(httpRes, err)
}

func (c *sessionsClient) Revoke(ctx context.Context, userID, sessionID string) error {
	httpRes, err := c.api.ServerSessionsAPI.RevokeSession(ctx, userID, sessionID).Execute()
	return wrapAPIError(httpRes, err)
}

// --- mapping helpers ---------------------------------------------------------

func userFromGenerated(g *generated.UserResponse) User {
	u := User{
		ID:            g.Id,
		EnvironmentID: g.EnvironmentId,
		Status:        UserStatus(g.Status),
		CreatedAt:     g.CreatedAt,
		UpdatedAt:     g.UpdatedAt,
		Name:          nullableStringToPtr(g.Name),
		Email:         nullableStringToPtr(g.Email),
		Phone:         nullableStringToPtr(g.Phone),
		Address:       nullableStringToPtr(g.Address),
	}
	if g.Locale.IsSet() && g.Locale.Get() != nil {
		l := Locale(*g.Locale.Get())
		u.Locale = &l
	}
	if g.DateOfBirth.IsSet() && g.DateOfBirth.Get() != nil {
		s := *g.DateOfBirth.Get()
		u.DateOfBirth = &s
	}
	if g.DeletedAt.IsSet() && g.DeletedAt.Get() != nil {
		t := *g.DeletedAt.Get()
		u.DeletedAt = &t
	}
	return u
}

func sessionFromGenerated(g *generated.UserSessionResponse) Session {
	return Session{
		ID:            g.Id,
		UserID:        g.UserId,
		EnvironmentID: g.EnvironmentId,
		UserAgent:     nullableStringToPtr(g.UserAgent),
		IPAddress:     nullableStringToPtr(g.IpAddress),
		CreatedAt:     g.CreatedAt,
		ExpiresAt:     g.ExpiresAt,
		LastUsedAt:    g.LastUsedAt,
	}
}

func nullableStringToPtr(n generated.NullableString) *string {
	if !n.IsSet() {
		return nil
	}
	v := n.Get()
	if v == nil {
		return nil
	}
	s := *v
	return &s
}

// wrapAPIError converts a generated *http.Response + error pair into a
// torii.APIError when the HTTP layer reported a non-2xx, or returns the
// original transport error otherwise. nil/2xx returns nil.
func wrapAPIError(httpRes *http.Response, err error) error {
	if err == nil {
		return nil
	}
	if httpRes == nil {
		return err
	}
	if httpRes.StatusCode >= 200 && httpRes.StatusCode < 300 {
		// Generated client returned an error during decode but the server
		// said 2xx — surface as-is.
		return err
	}
	apiErr := &APIError{Status: httpRes.StatusCode, Message: err.Error()}
	// Best-effort body decode: jwx's generated client closes httpRes.Body
	// before returning, but it also stashes a copy via GenericOpenAPIError.
	var generic generated.GenericOpenAPIError
	if errors.As(err, &generic) {
		apiErr.Body = generic.Body()
		var parsed struct {
			Code      string `json:"code"`
			SupportID string `json:"supportId"`
			Message   string `json:"message"`
		}
		if jerr := json.Unmarshal(generic.Body(), &parsed); jerr == nil {
			if parsed.Code != "" {
				apiErr.Code = parsed.Code
			}
			if parsed.SupportID != "" {
				apiErr.SupportID = parsed.SupportID
			}
			if parsed.Message != "" {
				apiErr.Message = parsed.Message
			}
		}
	} else if httpRes.Body != nil {
		if b, rerr := io.ReadAll(httpRes.Body); rerr == nil {
			apiErr.Body = b
		}
	}
	return apiErr
}
