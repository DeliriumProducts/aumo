package aumo

import "github.com/jinzhu/gorm"

type Receipt struct {
	gorm.Model
	Content string
	UserID  uint
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
