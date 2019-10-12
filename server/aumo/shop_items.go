package aumo

import "github.com/fr3fou/aumo/server/aumo/models"

// CreateShopItem creates a shop item
func (a *Aumo) CreateShopItem(name string, price float64, desc string, stock uint, image string) (ShopItem, error) {
	shopItem := &models.ShopItem{
		Name:        name,
		Price:       price,
		Image:       image,
		Description: desc,
		Stock:       stock,
	}

	if err := a.db.Create(shopItem).Error; err != nil {
		return models.ShopItem{}, err
	}

	return *shopItem, nil
}

// GetShopItemByID returns a user that has a matching email
func (a *Aumo) GetShopItemByID(id uint) (models.ShopItem, error) {
	var si models.ShopItem
	err := a.firstX(&si, "id = ?", id)
	return si, err
}

// GetShopItems returns a user that has a matching email
func (a *Aumo) GetShopItems() ([]models.ShopItem, error) {
	var si []models.ShopItem
	err := a.findX(&si)
	return si, err
}

// UpdateShopItem updates an item
func (a *Aumo) UpdateShopItem(si models.ShopItem) (models.ShopItem, error) {
	return si, a.updateX(&si)
}

// DeleteShopItem deletes an item
func (a *Aumo) DeleteShopItem(i models.ShopItem) error {
	return a.deleteX(i)
}
