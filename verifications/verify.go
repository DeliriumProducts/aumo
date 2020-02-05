package verifications

import (
	"time"

	"github.com/deliriumproducts/aumo/mail"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type Verifier struct {
	mailer mail.Mailer
	r      *redis.Client
}

func New(m mail.Mailer, r *redis.Client) *Verifier {
	return &Verifier{
		mailer: m,
		r:      r,
	}
}

func (v *Verifier) Send(to string, value interface{}, subject, body, link string, expiry time.Duration) error {
	token := uuid.New().String()

	err := v.r.Set(token, value, expiry).Err()
	if err != nil {
		return err
	}

	return v.mailer.SendMail(to,
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
}
