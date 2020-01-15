package rest

import (
	"github.com/deliriumproducts/aumo"
	"github.com/go-chi/chi"
)

func (rest *Rest) routes() {
	rest.router.Post("/register", rest.registerHandler)
	rest.router.Post("/login", rest.loginHandler)
	rest.router.Route("/me", func(r chi.Router) {
		r.Use(rest.WithAuth())
		r.Get("/", rest.meHandler)
	})

	// r.Route("/receipts", func(r chi.Router) {
	// 	r.Post("/", rest.NewReceiptHandler)
	// 	r.Get("/{id}", w.ReceiptHandler)
	// })

	rest.router.Route("/shop", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.productCreateHandler)
		r.Get("/", rest.productsHandler)
		// 	r.Get("/{id}", w.ShopItemHandler)
	})

}
