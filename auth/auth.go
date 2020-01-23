package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/deliriumproducts/aumo"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/google/uuid"
)

type authKey string

const (
	// CookieKey is the key used for cookies
	CookieKey = "aumo"
	// UserContextKey is the key used for contexts
	UserContextKey authKey = "aumo_user"
)

var (
	// ErrBadTypeAssertion is an error for when an assertion failed
	ErrBadTypeAssertion = errors.New("auth: failed to assert type")
)

// Authenticator holds the methods and config used for authentication
type Authenticator struct {
	redis      redis.Conn
	us         aumo.UserStore
	domain     string
	path       string
	expiryTime int
}

// New returns new Auth instance
func New(r redis.Conn, us aumo.UserStore, domain, path string, expiryTime int) *Authenticator {
	return &Authenticator{
		r,
		us,
		domain,
		path,
		expiryTime,
	}
}

// NewSession creates a session and returns the session ID
func (a *Authenticator) NewSession(u *aumo.User) (string, error) {
	sID := uuid.New().String()

	_, err := a.redis.Do("SETEX", sID, a.expiryTime, u.ID)
	if err != nil {
		return "", err
	}

	return sID, err
}

// Get gets a session from Redis based on the session ID
func (a *Authenticator) Get(sID string) (*aumo.User, error) {
	uID, err := redis.Uint64(a.redis.Do("GET", sID))
	if err != nil {
		return nil, err
	}

	return a.us.FindByID(nil, uint(uID), true)
}

func (a *Authenticator) Del(sID string) error {
	_, err := a.redis.Do("DEL", sID)
	return err
}

// GetFromRequest gets a session from Redis based on the Cookie value from the request
func (a *Authenticator) GetFromRequest(r *http.Request) (*aumo.User, error) {
	cookie, err := r.Cookie(CookieKey)
	if err != nil {
		return nil, err
	}

	return a.Get(cookie.Value)
}

// SetCookieHeader sets the cookie to the response
func (a *Authenticator) SetCookieHeader(w http.ResponseWriter, sID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieKey,
		Value:    sID,
		HttpOnly: true,
		Path:     a.path,
		Domain:   a.domain,
		Expires: time.Now().Add(
			time.Duration(a.expiryTime) * time.Second,
		),
	})
}

// WithUser sets a user to a context
func WithUser(ctx context.Context, user *aumo.User) context.Context {
	return context.WithValue(ctx, UserContextKey, *user)
}

// CurrentUser gets a user from a context
func CurrentUser(ctx context.Context) (aumo.User, error) {
	if user, ok := ctx.Value(UserContextKey).(aumo.User); ok {
		return user, nil
	}

	return aumo.User{}, ErrBadTypeAssertion
}
