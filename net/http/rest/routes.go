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
			r.With(rest.WithAuth()).Get("/{shop_id}", rest.shopGet)

			r.Group(func(r chi.Router) {
				r.Use(
					rest.WithAuth(aumo.Admin, aumo.ShopOwner),
					rest.WithShopOwnersAndAdmins,
				)

				r.Post("/", rest.shopCreate)
				r.Route("/{shop_id}", func(r chi.Router) {
					r.Put("/", rest.shopEdit)
					r.Delete("/", rest.shopDelete)
					r.Post("/add-owner", rest.shopAddOwner)
					r.Post("/remove-owner", rest.shopRemoveOwner)

					r.Route("/products", func(r chi.Router) {
						r.Get("/", rest.productGetAllByShop)
						r.Post("/", rest.productCreate)
						r.Put("/{product_id}", rest.productEdit)
						r.Delete("/{product_id}", rest.productDelete)
					})
				})
			})
		})

		r.Route("/products", func(r chi.Router) {
			r.Get("/", rest.productGetAll)
			r.Get("/{product_id}", rest.productGet)
		})

		r.Route("/receipts", func(r chi.Router) {
			r.With(rest.WithAuth(aumo.Admin)).Post("/", rest.receiptCreate)
			r.With(rest.WithAuth(aumo.Customer)).Get("/{id}", rest.receiptClaim)
		})

		r.Route("/orders", func(r chi.Router) {
			r.With(rest.WithAuth(aumo.Customer)).Post("/", rest.orderCreate)

			r.Group(func(r chi.Router) {
				r.Use(rest.WithAuth(aumo.Admin))
				r.Get("/", rest.orderGetAll)
				r.Get("/{id}", rest.orderGet)
			})
		})
	})
}
