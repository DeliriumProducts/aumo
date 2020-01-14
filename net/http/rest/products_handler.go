package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
)

type ProductForm struct {
	Name        string  `form:"name" validate:"required"`
	Image       string  `form:"image" validate:"required,url"`
	Price       float64 `form:"price" validate:"required,numeric"`
	Description string  `form:"description" validate:"required"`
	Stock       uint    `form:"stock" validate:"required,numeric"`
}

func (rest *Rest) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := rest.productService.Products()
	if err != nil {
		JSONError(w, err, http.StatusNotFound)
		return
	}

	JSON(w, products, http.StatusOK)
}

func (rest *Rest) NewProductHandler(w http.ResponseWriter, r *http.Request) {
	var npf ProductForm
	if ok := rest.Form(w, r, &npf); !ok {
		return
	}

	product := aumo.NewProduct(npf.Name, npf.Price, npf.Image, npf.Description, npf.Stock)

	err := rest.productService.Create(product)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	JSON(w, product, http.StatusOK)
}
