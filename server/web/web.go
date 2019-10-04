package web

import (
	"github.com/fr3fou/aumo/server/aumo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Config struct {
	*aumo.Aumo
}

type Web struct {
	Config
	Router *chi.Mux
}

func New(c Config) *Web {

	if c.Aumo == nil {
		panic("web: aumo instance not passed to web.New()")
	}

	r := chi.NewRouter()

	w := &Web{
		Config: c,
		Router: r,
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(ContentTypeJSON)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", w.RegisterHandler)
	})

	return w
}
