package rest

import (
	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/net/http/rest/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const (
	CookieStoreKey = "aumo"
	UserContextKey = "user"
	UserSessionKey = "user"
)

type Config struct {
	UserService    aumo.UserService
	ReceiptService aumo.ReceiptService
	OrderService   aumo.OrderService
	ProductService aumo.ProductService
	Auth           *auth.Auth
	CookieSecret   []byte
}

type Rest struct {
	Config
	Router *chi.Mux
}

func New(c Config) *Rest {
	if c.CookieSecret == nil {
		panic("http: CookieSecret not passed to rest.New()")
	}

	r := chi.NewRouter()

	rest := &Rest{
		Config: c,
		Router: r,
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
		r.Post("/register", rest.RegisterHandler)
		// 	r.Post("/login", rest.LoginHandler)
		// 	r.Group(func(r chi.Router) {
		// 		r.Use(rest.WithAuth)
		// 		r.Post("/claim-receipt/{id}", rest.ClaimReceiptHandler)
		// 		r.Post("/buy/{id}", rest.BuyHandler)
		// 	})
	})

	// r.Group(func(r chi.Router) {
	// 	r.Use(rest.WithAuth)
	// 	r.Get("/me", rest.MeHandler)
	// })

	// r.Route("/receipts", func(r chi.Router) {
	// 	r.Post("/", rest.NewReceiptHandler)
	// 	r.Get("/{id}", w.ReceiptHandler)
	// })

	// r.Route("/shop", func(r chi.Router) {
	// 	r.Post("/", w.NewShopItemHandler)
	// 	r.Get("/", w.ShopItemsHandler)
	// 	r.Get("/{id}", w.ShopItemHandler)
	// })

	return rest
}
