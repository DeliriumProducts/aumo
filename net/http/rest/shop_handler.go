package rest

import (
	"net/http"
)

func (rest *Rest) shopGetAll(w http.ResponseWriter, r *http.Request) {
	shops, err := rest.shopService.Shops()
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, shops, http.StatusOK)
}

func (rest *Rest) shopGet(w http.ResponseWriter, r *http.Request) {
	sID := rest.ParamNumber(w, r, "id")

	shop, err := rest.shopService.Shop(sID)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, shop, http.StatusOK)
}
