package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/net/http/rest/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
)

// Config is the configuration for the REST API
type Config struct {
	UserService    aumo.UserService
	ReceiptService aumo.ReceiptService
	OrderService   aumo.OrderService
	ProductService aumo.ProductService
	Auth           *auth.Authenticator
	CookieSecret   []byte
	MountRoute     string
}

// Rest is a REST API for Aumo
type Rest struct {
	router         *chi.Mux
	userService    aumo.UserService
	receiptService aumo.ReceiptService
	orderService   aumo.OrderService
	productService aumo.ProductService
	auth           *auth.Authenticator
	validator      *validator.Validate
	decoder        *form.Decoder
	cookieSecret   []byte
}

func New(c Config) *Rest {
	switch {
	case c.UserService == nil:
		panic("rest: UserService not provided")
	case c.ReceiptService == nil:
		panic("rest: ReceiptService not provided")
	case c.OrderService == nil:
		panic("rest: OrderService not provided")
	case c.ProductService == nil:
		panic("rest: ProductService not provided")
	case c.Auth == nil:
		panic("rest: Authenticator not provided")
	case c.CookieSecret == nil:
		panic("rest: CookieSecret not provided")
	}

	r := chi.NewRouter()
	validator := validator.New()
	decoder := form.NewDecoder()

	en := en.New()
	uni := ut.New(en, en)
	trans, found := uni.GetTranslator("en")
	if !found {
		panic("rest: translator couldn't be found")
	}

	enTrans.RegisterDefaultTranslations(validator, trans)

	rest := &Rest{
		router:         r,
		userService:    c.UserService,
		receiptService: c.ReceiptService,
		orderService:   c.OrderService,
		productService: c.ProductService,
		auth:           c.Auth,
		validator:      validator,
		translator:     trans,
		decoder:        decoder,
		cookieSecret:   c.CookieSecret,
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
	r.Use(Security)
	r.Mount(c.MountRoute, r)

	r.Route("/users", func(r chi.Router) {
		r.Use(ParseForm)
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

func (rest *Rest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rest.router.ServeHTTP(w, r)
}

// Error is a json type for error handling
type Error struct {
	Error string `json:"error"`
}

var errMarshaling, _ = json.Marshal(Error{
	Error: "Failed to Marshal Error",
})

// JSONError is a convenience function for handling errors
func JSONError(w http.ResponseWriter, err error, statusCode int) {
	json, err := json.Marshal(Error{
		Error: err.Error(),
	})

	// We fallback to a default error if we encountered one
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errMarshaling)
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(json)
}

func JSON(w http.ResponseWriter, v interface{}, statusCode int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}
