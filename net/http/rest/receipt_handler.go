package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
)

func (rest *Rest) receiptsCreate(w http.ResponseWriter, r *http.Request) {
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
