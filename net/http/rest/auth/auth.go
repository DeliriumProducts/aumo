package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/deliriumproducts/aumo"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/google/uuid"
)

var (
	// ErrBadTypeAssertion is an error for when an assertion failed
	ErrBadTypeAssertion = errors.New("auth: failed to assert type")
)

// Authenticator holds the methods and config used for authentication
type Authenticator struct {
	redis      redis.Conn
	us         aumo.UserService
	expiryTime string
}

// New returns new Auth instance
func New(r redis.Conn, us aumo.UserService, expiryTime string) *Authenticator {
	return &Authenticator{
		redis:      r,
		us:         us,
		expiryTime: expiryTime,
	}
}

// NewSession creates a session and returns the session ID
func (a *Authenticator) NewSession(u *aumo.User) (string, error) {
	sess := uuid.New().String()

	_, err := a.redis.Do("SETEX", sess, a.expiryTime, u.ID)
	if err != nil {
		return "", err
	}

	return sess, err
}

func (a *Authenticator) Get(sess string) (*aumo.User, error) {
	val, err := a.redis.Do("GET", sess)
	if err != nil {
		return nil, err
	}

	uID, ok := val.(uint)
	if !ok {
		return nil, ErrBadTypeAssertion
	}

	return a.us.User(uID, false)
}

func (a *Authenticator) GetFromReq(r *http.Request) (*aumo.User, error) {
	at := r.Header.Get("Authorization")
	tok := strings.Split(at, " ")[1]

	return a.Get(tok)
}
