package aumo

import "github.com/jinzhu/gorm"

type ShopItem struct {
	gorm.Model
	Name        string
	Price       float64
	Description string
	Quantity    uint
}

// CreateShopItem creates a shop item
func (a *Aumo) CreateShopItem(name string, price float64, desc string, quantity uint) (ShopItem, error) {
	shopItem := &ShopItem{
		Name:        name,
		Price:       price,
		Description: desc,
		Quantity:    quantity,
	}

	if err := a.DB.Create(shopItem).Error; err != nil {
		return ShopItem{}, err
	}

	return *shopItem, nil
}

// GetShopItemByID returns a user that has a matching email
func (a *Aumo) GetShopItemByID(id uint) (ShopItem, error) {
	return a.getShopItem("id = ?", id)
}

// getShopItem is an internal helper function to quickly get a shop item
func (a *Aumo) getShopItem(where ...interface{}) (ShopItem, error) {
	var si ShopItem

	err := a.DB.First(&si, where...).Error

	if err != nil {
		return ShopItem{}, err
	}

	return si, nil
}

// UpdateShopItem updates an item
func (a *Aumo) UpdateShopItem(old, new ShopItem) (ShopItem, error) {
	return old, a.updateX(&old, new)
}

// DeleteShopItem deletes an item
func (a *Aumo) DeleteShopItem(i ShopItem) error {
	return a.deleteX(i)
}
