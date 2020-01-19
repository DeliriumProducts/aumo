package products

import "github.com/deliriumproducts/aumo"

type service struct {
	store aumo.ProductStore
}

// New returns an instance of `aumo.ProductService`
func New(store aumo.ProductStore) aumo.ProductService {
	return &service{
		store: store,
	}
}

func (ps *service) Product(id uint) (*aumo.Product, error) {
	return ps.store.FindByID(nil, id)
}

func (ps *service) Products() ([]aumo.Product, error) {
	return ps.store.FindAll(nil)
}

func (ps *service) Create(p *aumo.Product) error {
	return ps.store.Save(nil, p)
}

func (ps *service) Update(id uint, p *aumo.Product) error {
	return ps.store.Update(nil, id, p)
}

func (ps *service) Delete(id uint) error {
	return ps.store.Delete(nil, id)
}
