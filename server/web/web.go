package web

import (
	"encoding/gob"

	"github.com/fr3fou/aumo/server/aumo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	}

	gob.Register(aumo.User{})

	w := &Web{
		Config: c,
		Router: r,
		store:  store,
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(ContentTypeJSON)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", w.RegisterHandler)
		r.Post("/login", w.LoginHandler)
		r.Get("/secret", w.SecretHandler)
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
