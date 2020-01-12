package auth

import (
	"errors"

	"github.com/deliriumproducts/aumo"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrBadTypeAssertion = errors.New("auth: failed to assert type")
)

// Config struct holds auth related configs
type Config struct {
	Redis       redis.Conn
	UserService aumo.UserService
	ExpiryTime  string
}

// Auth struct holds the methods and config used for authentication
type Auth struct {
	Config
}

// New returns new Auth instance
func New(c Config) *Auth {
	return &Auth{
		Config: c,
	}
}

// SetSess sets a sesion in redis then returns it
func (a *Auth) Set(u aumo.User) (string, error) {
	sess := uuid.NewV4().String()

	_, err := a.Redis.Do("SETEX", sess, a.ExpiryTime, u.ID)
	return sess, err
}

func (a *Auth) Get(sess string) (*aumo.User, error) {
	val, err := a.Redis.Do("GET", sess)
	if err != nil {
		return nil, err
	}

	uID, ok := val.(uint)
	if !ok {
		return nil, ErrBadTypeAssertion
	}

	return a.UserService.User(uID, false)
}
