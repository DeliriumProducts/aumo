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

// DecrementStock decreases the stock of a `Product`
func (p *Product) DecrementStock(n uint) {
	p.Stock = p.Stock - n
}

// IncrementStock increases the stock of a `Product`
func (p *Product) IncrementStock(n uint) {
	p.Stock = p.Stock + n
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

// ProductService contains all `Product`
// related business logic
type ProductService interface {
	Product(id uint) *Product
	Products() []Product
	Save(*Product)
	Update(*Product)
	Delete(*Product)
}
