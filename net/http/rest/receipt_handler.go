package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/auth"
)

func (rest *Rest) receiptCreate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Content string `form:"content" validate:"required"`
	}

	var re request
	if ok := rest.Form(w, r, &re); !ok {
		return
	}

	receipt := aumo.NewReceipt(re.Content)

	err := rest.receiptService.Create(receipt)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, receipt, http.StatusOK)
}

func (rest *Rest) receiptClaim(w http.ResponseWriter, r *http.Request) {
	rID := rest.ParamNumber(w, r, "id")

	user, err := auth.CurrentUser(r.Context())
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	receipt, err := rest.receiptService.ClaimReceipt(user.ID, rID)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, receipt, http.StatusOK)
}
