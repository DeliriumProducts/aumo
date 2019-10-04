package web

import (
	"encoding/json"
	"log"
	"net/http"
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
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(receipt); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
