package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo/auth"
)

func (rest *Rest) orderHandlerCreate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		ProductID uint `form:"product_id" validate:"required,numeric"`
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
