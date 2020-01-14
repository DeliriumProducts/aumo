package rest

import (
	"net/http"
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
	return
}
