package models

import (
	"database/sql"
	"errors"
)

var (
	ErrUserIDAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

type Receipt struct {
	Model
	Content string        `json:"content"`
	UserID  sql.NullInt64 `json:"-"`
}

// SetUserID claims a receipt with the provided ID
func (r *Receipt) SetUserID(userID int64) error {
	if !r.UserID.Valid {
		return ErrUserIDAlreadySet
	}

	r.UserID.Int64 = userID

	return nil
}
