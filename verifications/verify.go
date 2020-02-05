package verifications

import (
	"time"

	"github.com/deliriumproducts/aumo/mail"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
)

// Verifier is a struct for verifications
//
// First part (Sending)
//
// Generate a unique token (UUID), which will be used as a key
// in a Redis store. A value is also provided to set in the redis store.
// This is followed by sending an email which contains the UUID + a link.
// A user can check their email and click the link to proceed to the second part.
//
// Second part (Verifying)
//
// After a user has clicked the link, the Verifier can check if the UUID exists
// and if it does, the value will be returned to the caller, meaning they can
// carry on with their actual logic.
type Verifier struct {
	mailer mail.Mailer
	r      *redis.Client
}

// New returns an instance of a Verifier
func New(m mail.Mailer, r *redis.Client) *Verifier {
	return &Verifier{
		mailer: m,
		r:      r,
	}
}

// Send starts the sirst part of the verification process
func (v *Verifier) Send(to string, value, subject, body, link string, expiry time.Duration) (string, error) {
	token := uuid.New().String()

	err := v.r.Set(token, value, expiry).Err()
	if err != nil {
		return "", err
	}

	err = v.mailer.SendMail(to,
		"To: "+to+
			"\r\n"+
			"Subject: "+subject+
			"\r\n"+
			"\r\n"+
			body+
			"\r\n"+
			"Your link is "+link+"/"+token+
			"\r\n",
	)

	if err != nil {
		return "", err
	}

	return token, err
}

// Verify is the second part of the verification process
func (v *Verifier) Verify(token string) (string, error) {
	return v.r.Get(token).Result()
}
