package models

type ShopItem struct {
	Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Stock       uint    `json:"stock"`
}

// DecrementStock decreases the stock of an item
func (si *ShopItem) DecrementStock(i uint) {
	si.Stock = si.Stock - i
}

// IncrementStock increases the stock of an item
func (si *ShopItem) IncrementStock(i uint) {
	si.Stock = si.Stock + i
}
