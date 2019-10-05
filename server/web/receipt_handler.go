package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type NewReceiptForm struct {
	Content string `json:"content"`
}

func (wb *Web) NewReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var nrw NewReceiptForm
	if err := json.NewDecoder(r.Body).Decode(&nrw); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	receipt, err := wb.CreateReceipt(nrw.Content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(receipt); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (wb *Web) ReceiptHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rec, err := wb.GetReceiptByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
