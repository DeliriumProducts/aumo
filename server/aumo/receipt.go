package aumo

import (
	"database/sql"
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	ErrUserIDAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

type Receipt struct {
	gorm.Model
	Content string
	UserID  sql.NullInt64
}

func (a *Aumo) CreateReceipt(content string) (Receipt, error) {
	receipt := &Receipt{
		Content: content,
	}

	if err := a.DB.Create(receipt).Error; err != nil {
		return Receipt{}, err
	}

	return *receipt, nil
}

func (r *Receipt) SetUserID(userID int64) error {
	if !r.UserID.Valid {
		return ErrUserIDAlreadySet
	}

	r.UserID.Value = userID

	return nil
}
