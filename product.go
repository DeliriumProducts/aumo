package aumo

// Product is a product in the shop of aumo
type Product struct {
	Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Stock       uint    `json:"stock"`
}

// NewProduct is a constructor for `Product`
func NewProduct(name string, price float64, image string, description string, stock uint) *Product {
	return &Product{
		Name:        name,
		Price:       price,
		Image:       image,
		Description: description,
		Stock:       stock,
	}
}

// Order is an order of a `User`
type Order struct {
	User    *User    `json:"-"`
	Product *Product `json:"product"`
	Amount  uint     `json:"amount"`
}

// NewOrder is a constructor for `Order`
func NewOrder(u *User, p *Product, n uint) *Order {
	return &Order{
		User:    u,
		Product: p,
		Amount:  n,
	}
}

// DecrementStock decreases the stock of a `Product`
func (p *Product) DecrementStock(n uint) {
	p.Stock = p.Stock - n
}

// IncrementStock increases the stock of a `Product`
func (p *Product) IncrementStock(n uint) {
	p.Stock = p.Stock + n
}
