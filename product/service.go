package product

import "github.com/deliriumproducts/aumo"

type productService struct {
	store aumo.ProductStore
}

// New returns an instance of `aumo.ProductService`
func New(store aumo.ProductStore) aumo.ProductService {
	return &productService{
		store: store,
	}
}

func (ps *productService) Product(id uint) (*aumo.Product, error) {
	return ps.store.FindByID(id)
}

func (ps *productService) Products() ([]aumo.Product, error) {
	return ps.store.FindAll()
}

func (ps *productService) Create(p *aumo.Product) error {
	return ps.store.Save(p)
}

func (ps *productService) Update(id uint, p *aumo.Product) error {
	return ps.store.Update(id, p)
}

func (ps *productService) Delete(id uint) error {
	return ps.store.Delete(id)
}
