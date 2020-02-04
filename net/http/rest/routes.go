package rest

import (
	"github.com/deliriumproducts/aumo"
	"github.com/go-chi/chi"
)

func (rest *Rest) routes() {
	rest.router.Post("/register", rest.userRegister)
	rest.router.Post("/login", rest.userLogin)
	rest.router.Group(func(r chi.Router) {
		r.Use(rest.WithAuth())
		r.Get("/logout", rest.userLogout)
		r.Get("/me", rest.userGetCurrent)
	})

	rest.router.Route("/users", func(r chi.Router) {
		r.Use(rest.WithAuth(aumo.Admin))
		r.Get("/", rest.userGetAll)
		r.Get("/{id}", rest.userGet)
		r.Put("/{id}/set-role", rest.userEditRole)
		r.Put("/{id}/add-points", rest.userAddPoints)
		r.Put("/{id}/sub-points", rest.userSubPoints)
		r.Delete("/{id}", rest.userDelete)
	})

	rest.router.Route("/receipts", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.receiptCreate)
		r.With(rest.WithAuth(aumo.Customer)).Get("/{id}", rest.receiptClaim)
	})

	rest.router.Route("/products", func(r chi.Router) {
		r.Get("/", rest.productGetAll)
		r.Get("/{id}", rest.productGet)
		r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.productCreate)
		r.With(rest.WithAuth(aumo.Admin)).Put("/{id}", rest.productEdit)
		r.With(rest.WithAuth(aumo.Admin)).Delete("/{id}", rest.productDelete)
	})

	rest.router.Route("/orders", func(r chi.Router) {
		r.With(rest.WithAuth(aumo.Customer)).Post("/", rest.orderCreate)
		r.With(rest.WithAuth(aumo.Admin)).Get("/", rest.orderGetAll)
		r.With(rest.WithAuth(aumo.Admin)).Get("/{id}", rest.orderGet)
	})
}
