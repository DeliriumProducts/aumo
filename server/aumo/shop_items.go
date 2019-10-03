package aumo

import "github.com/jinzhu/gorm"

type ShopItem struct {
	gorm.Model
	Name        string
	Price       float64
	Description string
	Quantity    uint
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

	return *shopItem, nil
}
