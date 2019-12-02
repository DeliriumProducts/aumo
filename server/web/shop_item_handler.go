package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type NewShopItemForm struct {
	Name        string  `json:"name"`
	Image       string  `json:"image"`
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

	shopItem, err := wb.CreateShopItem(nsi.Name, nsi.Price, nsi.Description, nsi.Stock, nsi.Image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(shopItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (wb *Web) ShopItemHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	si, err := wb.GetShopItemByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(si); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (wb *Web) ShopItemsHandler(w http.ResponseWriter, r *http.Request) {
	si, err := wb.GetShopItems()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(si); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
