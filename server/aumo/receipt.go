package aumo

import "github.com/fr3fou/aumo/server/aumo/models"

// CreateReceipt creates a receipt
func (a *Aumo) CreateReceipt(content string) (models.Receipt, error) {
	receipt := &models.Receipt{
		Content: content,
	}

	if err := a.db.Create(receipt).Error; err != nil {
		return models.Receipt{}, err
	}
	return *receipt, nil
}

// GetReceiptByID gets a receipt by ID
func (a *Aumo) GetReceiptByID(id uint) (models.Receipt, error) {
	var r models.Receipt
	err := a.firstX(&r, "id = ?", id)
	return r, err
}

// SetReceiptUserID claims the receipt by calling the ClaimReceipt(r) (adds receipt to the receipt list of the user)
// Sets the user id in the receipt (receipt is claimed by the user)
func (a *Aumo) SetReceiptUserID(u models.User, r models.Receipt) error {
	u.ClaimReceipt(r)
	r.SetUserID(int64(u.ID))

	err := a.updateX(u)
	if err != nil {
		return err
	}

	err = a.updateX(r)
	if err != nil {
		return err
	}

	return nil
}
