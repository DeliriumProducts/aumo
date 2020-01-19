package rest

import (
	"github.com/deliriumproducts/aumo"
	"github.com/go-chi/chi"
)

func (rest *Rest) routes() {
	rest.router.Post("/register", rest.userHandlerRegister)
	rest.router.Post("/login", rest.userHandlerLogin)
	rest.router.Route("/me", func(r chi.Router) {
		r.Use(rest.WithAuth())
		r.Get("/", rest.userHandlerGet)
	})

	// r.Route("/receipts", func(r chi.Router) {
	// 	r.Post("/", rest.NewReceiptHandler)
	// 	r.Get("/{id}", w.ReceiptHandler)
	// })

	rest.router.Route("/products", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.productHandlerCreate)
		r.Get("/", rest.productHandlerGetAll)
		// 	r.Get("/{id}", w.ShopItemHandler)
	})

	rest.router.Route("/orders", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Customer)).Post("/", rest.orderHandlerCreate)
	})
}
