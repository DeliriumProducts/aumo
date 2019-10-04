package web

import (
	"encoding/json"
	"net/http"
)

type NewShopItemForm struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       uint    `json:"stock"`
}

func (wb *Web) NewShopItemHandler(w http.ResponseWriter, r *http.Request) {
	var nsi NewShopItemForm
	if err := json.NewDecoder(r.Body).Decode(&nsi); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shopItem, err := wb.CreateShopItem(nsi.Name, nsi.Price, nsi.Description, nsi.Stock)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(shopItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
