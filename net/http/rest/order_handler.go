package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo/auth"
)

func (rest *Rest) orderCreate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		ProductID uint `form:"product_id" validate:"required,numeric" json:"product_id"`
	}

	var of request
	if ok := rest.Form(w, r, &of); !ok {
		return
	}

	user, err := auth.CurrentUser(r.Context())
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	order, err := rest.orderService.PlaceOrder(user.ID, of.ProductID)
	if err != nil {
		rest.JSONError(w, err, http.StatusBadRequest)
		return
	}

	rest.JSON(w, order, http.StatusOK)
}

func (rest *Rest) orderGetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := rest.orderService.Orders()
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, orders, http.StatusOK)
}

func (rest *Rest) orderGet(w http.ResponseWriter, r *http.Request) {
	orderID := rest.ParamNumber(w, r, "id")

	order, err := rest.orderService.Order(orderID)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, order, http.StatusOK)
}
