package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
)

func (rest *Rest) productGetAll(w http.ResponseWriter, r *http.Request) {
	products, err := rest.productService.Products()
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, products, http.StatusOK)
}

func (rest *Rest) productCreate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name        string  `form:"name" validate:"required" json:"name"`
		Image       string  `form:"image" validate:"required,url" json:"image"`
		Price       float64 `form:"price" validate:"required,numeric" json:"price"`
		Description string  `form:"description" validate:"required" json:"description"`
		Stock       uint    `form:"stock" validate:"required,numeric" json:"stock"`
		ShopID      uint    `form:"shop_id" validate:"required,numeric" json:"shop_id"`
	}

	var npf request
	if ok := rest.Form(w, r, &npf); !ok {
		return
	}

	product := aumo.NewProduct(npf.Name, npf.Price, npf.Image, npf.Description, npf.Stock, npf.ShopID)

	err := rest.productService.Create(product)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, product, http.StatusOK)
}

func (rest *Rest) productGet(w http.ResponseWriter, r *http.Request) {
	pID := rest.ParamNumber(w, r, "id")

	product, err := rest.productService.Product(pID)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, product, http.StatusOK)
}

func (rest *Rest) productEdit(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name        string  `form:"name" validate:"required" json:"name"`
		Image       string  `form:"image" validate:"required,url" json:"image"`
		Price       float64 `form:"price" validate:"required,numeric" json:"price"`
		Description string  `form:"description" validate:"required" json:"description"`
		Stock       uint    `form:"stock" validate:"required,numeric" json:"stock"`
		ShopID      uint    `form:"shop_id" validate:"required,numeric" json:"shop_id"`
	}

	pID := rest.ParamNumber(w, r, "id")

	var npf request
	if ok := rest.Form(w, r, &npf); !ok {
		return
	}

	product := aumo.NewProduct(npf.Name, npf.Price, npf.Image, npf.Description, npf.Stock, npf.ShopID)

	err := rest.productService.Update(pID, product)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, product, http.StatusOK)
}

func (rest *Rest) productDelete(w http.ResponseWriter, r *http.Request) {
	pID := rest.ParamNumber(w, r, "id")

	err := rest.productService.Delete(pID)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, Message{"Product successfully deleted!"}, http.StatusOK)
}
