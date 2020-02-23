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
			r.Use(rest.Authentication())

			r.Get("/logout", rest.userLogout)
			r.Get("/me", rest.userGetCurrent)
			r.Put("/me", rest.userEdit)
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(rest.Authentication(aumo.Admin))

			r.Get("/", rest.userGetAll)
			r.Get("/{id}", rest.userGet)
			r.Put("/{id}/set-role", rest.userEditRole)
			r.Put("/{id}/add-points", rest.userAddPoints)
			r.Put("/{id}/sub-points", rest.userSubPoints)
			r.Delete("/{id}", rest.userDelete)
		})

		r.Route("/shops", func(r chi.Router) {
			r.Get("/", rest.shopGetAll)
			r.With(rest.Authentication(aumo.Admin, aumo.ShopOwner)).Post("/", rest.shopCreate)

			r.Route("/{shop_id}", func(r chi.Router) {
				r.With(rest.Authentication()).Get("/", rest.shopGet)

				r.Group(func(r chi.Router) {
					r.Use(rest.Authentication(aumo.Admin, aumo.ShopOwner))
					r.Use(rest.OwnsShop)

					r.Put("/", rest.shopEdit)
					r.Delete("/", rest.shopDelete)
					r.Post("/add-owner", rest.shopAddOwner)
					r.Post("/remove-owner", rest.shopRemoveOwner)
				})

				r.Route("/products", func(r chi.Router) {
					r.Get("/", rest.productGetAllByShop)
					r.Get("/{product_id}", rest.productGet)

					r.Group(func(r chi.Router) {
						r.Use(rest.Authentication(aumo.Admin, aumo.ShopOwner))
						r.Use(rest.OwnsShop)

						r.Post("/", rest.productCreate)
						r.Put("/{product_id}", rest.productEdit)
						r.Delete("/{product_id}", rest.productDelete)
					})
				})
			})
		})

		r.Route("/receipts", func(r chi.Router) {
			r.With(rest.Authentication(aumo.Admin)).Post("/", rest.receiptCreate)
			r.With(rest.Authentication()).Get("/{id}", rest.receiptClaim)
		})

		r.Route("/orders", func(r chi.Router) {
			r.With(rest.Authentication()).Post("/", rest.orderCreate)

			r.Group(func(r chi.Router) {
				r.Use(rest.Authentication(aumo.Admin))

				r.Get("/", rest.orderGetAll)
				r.Get("/{id}", rest.orderGet)
			})
		})
	})
}
