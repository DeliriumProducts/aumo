package aumo

// Product is a product in the shop of aumo
type Product struct {
	ID          uint    `json:"id" db:"id,omitempty" validate:"-"`
	Name        string  `json:"name" db:"name" validate:"required"`
	Price       float64 `json:"price" db:"price" validate:"required"`
	Image       string  `json:"image" db:"image" validate:"required,url"`
	Description string  `json:"description" db:"description" validate:"required"`
	Stock       uint    `json:"stock" db:"stock" validate:"required,gte=1"`
}

// DecrementStock decreases the stock of a `Product`
func (p *Product) DecrementStock() {
	p.Stock--
}

// IncrementStock increases the stock of a `Product`
func (p *Product) IncrementStock() {
	p.Stock++
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

// ProductStore contains all `Product`
// related persistance logic
type ProductStore interface {
	FindByID(id uint) (*Product, error)
	FindAll() ([]Product, error)
	Save(*Product) error
	Update(id uint, p *Product) error
	Delete(id uint) error
}
