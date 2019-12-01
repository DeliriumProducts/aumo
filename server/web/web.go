package web

import (
	"encoding/gob"

	"github.com/fr3fou/aumo/server/aumo"
	"github.com/fr3fou/aumo/server/aumo/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
)

type Config struct {
	*aumo.Aumo
	CookieSecret []byte
}

type Web struct {
	Config
	Router *chi.Mux
	store  *sessions.CookieStore
}

func New(c Config) *Web {
	if c.Aumo == nil {
		panic("web: aumo instance not passed to web.New()")
	}

	r := chi.NewRouter()

	store := sessions.NewCookieStore(c.CookieSecret)

	store.Options = &sessions.Options{
		MaxAge:   3600 * 24,
		HttpOnly: true,
		Path:     "/",
	}

	gob.Register(models.User{})

	w := &Web{
		Config: c,
		Router: r,
		store:  store,
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}).Handler,
	)
	r.Use(ContentTypeJSON)

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", w.RegisterHandler)
		r.Post("/login", w.LoginHandler)
		r.Group(func(r chi.Router) {
			r.Use(w.WithAuth)
			r.Post("/claim-receipt/{id}", w.ClaimReceiptHandler)
			r.Post("/buy/{id}", w.BuyHandler)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(w.WithAuth)
		r.Get("/me", w.MeHandler)
	})

	r.Route("/receipts", func(r chi.Router) {
		r.Post("/", w.NewReceiptHandler)
		r.Get("/{id}", w.ReceiptHandler)
	})

	r.Route("/shop", func(r chi.Router) {
		r.Post("/", w.NewShopItemHandler)
		r.Get("/", w.ShopItemsHandler)
		r.Get("/{id}", w.ShopItemHandler)
	})

	return w
}
