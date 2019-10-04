package web

import (
	"github.com/fr3fou/aumo/server/aumo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Config struct {
	*aumo.Aumo
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

	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store := sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

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
	})

	r.Route("/receipts", func(r chi.Router) {
		r.Post("/", w.NewReceiptHandler)
		r.Get("/{id}", w.ReceiptHandler)
	})

	r.Route("/shop", func(r chi.Router) {
		r.Post("/", w.NewShopItemHandler)
		// r.Get("/{id}", w.ReceiptHandler)
	})

	return w
}
