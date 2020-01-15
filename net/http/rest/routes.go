package rest

import (
	"github.com/deliriumproducts/aumo"
	"github.com/go-chi/chi"
)

func (rest *Rest) routes() {
	rest.router.Post("/register", rest.RegisterHandler)
	rest.router.Post("/login", rest.LoginHandler)
	rest.router.Route("/me", func(r chi.Router) {
		r.Use(rest.WithAuth())
		r.Get("/", rest.MeHandler)
	})

	// r.Route("/receipts", func(r chi.Router) {
	// 	r.Post("/", rest.NewReceiptHandler)
	// 	r.Get("/{id}", w.ReceiptHandler)
	// })

	rest.router.Route("/shop", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.NewProductHandler)
		r.Get("/", rest.ProductsHandler)
		// 	r.Get("/{id}", w.ShopItemHandler)
	})

}
