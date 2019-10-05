package aumo

type ShopItem struct {
	Model
	Name        string
	Price       float64
	Description string
	Stock       uint
}

// DecrementStock decreases the stock of an item
func (si *ShopItem) DecrementStock(i uint) {
	si.Stock = si.Stock - i
}

// IncrementStock increases the stock of an item
func (si *ShopItem) IncrementStock(i uint) {
	si.Stock = si.Stock + i
}

// CreateShopItem creates a shop item
func (a *Aumo) CreateShopItem(name string, price float64, desc string, stock uint) (ShopItem, error) {
	shopItem := &ShopItem{
		Name:        name,
		Price:       price,
		Description: desc,
		Stock:       stock,
	}

	if err := a.db.Create(shopItem).Error; err != nil {
		return ShopItem{}, err
	}

	return *shopItem, nil
}

// GetShopItemByID returns a user that has a matching email
func (a *Aumo) GetShopItemByID(id uint) (ShopItem, error) {
	var si ShopItem
	err := a.firstX(&si, "id = ?", id)
	return si, err
}

// GetShopItems returns a user that has a matching email
func (a *Aumo) GetShopItems() ([]ShopItem, error) {
	var si []ShopItem
	err := a.findX(&si)
	return si, err
}

// UpdateShopItem updates an item
func (a *Aumo) UpdateShopItem(si ShopItem) (ShopItem, error) {
	return si, a.updateX(&si)
}

// DeleteShopItem deletes an item
func (a *Aumo) DeleteShopItem(i ShopItem) error {
	return a.deleteX(i)
}
