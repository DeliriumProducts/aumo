package aumo

import "github.com/jinzhu/gorm"

type ShopItem struct {
	gorm.Model
	Name        string
	Price       float64
	Description string
	Quantity    uint
	a           *Aumo
}

func (a *Aumo) AddShopItem(name string, price float64, desc string, quantity uint) (ShopItem, error) {
	shopItem := &ShopItem{
		Name:        name,
		Price:       price,
		Description: desc,
		Quantity:    quantity,
	}

	if err := a.DB.Create(shopItem).Error; err != nil {
		return ShopItem{}, err
	}

	shopItem.a = a

	return *shopItem, nil
}

// GetShopItemById returns a user that has a matching email
func (a *Aumo) GetShopItemById(id uint) (ShopItem, error) {
	return a.getShopItem(&ShopItem{}, "id = ?", id)
}

// getShopItem is an internal helper function to quickly get a shop item
func (a *Aumo) getShopItem(out interface{}, where ...interface{}) (ShopItem, error) {
	var si ShopItem

	err := a.DB.First(out, where...).Error

	if err != nil {
		return ShopItem{}, nil
	}

	si.a = a

	return si, nil
}

// func (si *ShopItem)
