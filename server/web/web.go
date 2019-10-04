package web

import (
	"github.com/fr3fou/aumo/server/aumo"
	"github.com/go-chi/chi"
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

	r.Route("/users", func(r chi.Router) {
		r.Post("/", w.RegisterHandler)
	})

	return w
}
