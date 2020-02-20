package rest

import (
	"github.com/deliriumproducts/aumo"
	"github.com/go-chi/chi"
)

func (rest *Rest) mount(mnt string) {
	rest.router.Route(mnt, func(r chi.Router) {
		r.Post("/register", rest.userRegister)
		r.Post("/login", rest.userLogin)
		r.Get("/confirm-email/{token}", rest.userConfirmEmail)
		r.Group(func(r chi.Router) {
			r.Use(rest.WithAuth())
			r.Get("/logout", rest.userLogout)
			r.Get("/me", rest.userGetCurrent)
			r.Put("/me", rest.userEdit)
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(rest.WithAuth(aumo.Admin))
			r.Get("/", rest.userGetAll)
			r.Get("/{id}", rest.userGet)
			r.Put("/{id}/set-role", rest.userEditRole)
			r.Put("/{id}/add-points", rest.userAddPoints)
			r.Put("/{id}/sub-points", rest.userSubPoints)
			r.Delete("/{id}", rest.userDelete)
		})

		r.Route("/shops", func(r chi.Router) {
			r.With(rest.WithAuth()).Get("/", rest.shopGetAll)
			r.With(rest.WithAuth()).Get("/{id}", rest.shopGet)
			r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.shopCreate)
			r.With(rest.WithAuth(aumo.Admin), rest.WithShopOwnersAndAdmins).Put("/{id}", rest.shopEdit)
			r.With(rest.WithAuth(aumo.Admin), rest.WithShopOwnersAndAdmins).Delete("/{id}", rest.shopDelete)
			r.With(rest.WithAuth(aumo.Admin), rest.WithShopOwnersAndAdmins).Post("/{id}/add-owner", rest.shopAddOwner)
			r.With(rest.WithAuth(aumo.Admin), rest.WithShopOwnersAndAdmins).Post("/{id}/remove-owner", rest.shopRemoveOwner)
		})

		r.Route("/receipts", func(r chi.Router) {
			r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.receiptCreate)
			r.With(rest.WithAuth(aumo.Customer)).Get("/{id}", rest.receiptClaim)
		})

		r.Route("/products", func(r chi.Router) {
			r.With(rest.WithAuth()).Get("/", rest.productGetAll)
			r.With(rest.WithAuth()).Get("/{id}", rest.productGet)
			r.With(rest.WithAuth(aumo.Admin), rest.WithShopOwnersAndAdmins).Post("/", rest.productCreate)
			r.With(rest.WithAuth(aumo.Admin), rest.WithShopOwnersAndAdmins).Put("/{id}", rest.productEdit)
			r.With(rest.WithAuth(aumo.Admin), rest.WithShopOwnersAndAdmins).Delete("/{id}", rest.productDelete)
		})

		r.Route("/orders", func(r chi.Router) {
			r.With(rest.WithAuth(aumo.Customer)).Post("/", rest.orderCreate)
			r.With(rest.WithAuth(aumo.Admin)).Get("/", rest.orderGetAll)
			r.With(rest.WithAuth(aumo.Admin)).Get("/{id}", rest.orderGet)
		})
	})
}
