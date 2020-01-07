package aumo

// Product is a product in the shop of aumo
type Product struct {
	ID          uint    `json:"id" db:"id,omitempty"`
	Name        string  `json:"name" db:"name"`
	Price       float64 `json:"price" db:"price"`
	Image       string  `json:"image" db:"image"`
	Description string  `json:"description" db:"description"`
	Stock       uint    `json:"stock" db:"stock"`
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
	Product(id uint) (*Product, error)
	Products() ([]Product, error)
	Create(*Product) error
	Update(id uint, p *Product) error
	Delete(id uint) error
}
