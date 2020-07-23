package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/auth"
	"github.com/deliriumproducts/aumo/verifications"
	"github.com/didip/tollbooth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

// Config is the configuration for the REST API
type Config struct {
	UserService    aumo.UserService
	ReceiptService aumo.ReceiptService
	OrderService   aumo.OrderService
	ProductService aumo.ProductService
	ShopService    aumo.ShopService
	Auth           *auth.Authenticator
	Verifier       *verifications.Verifier
	MountRoute     string
	BackendURL     string
}

// Rest is a REST API for Aumo
type Rest struct {
	router         *chi.Mux
	userService    aumo.UserService
	receiptService aumo.ReceiptService
	orderService   aumo.OrderService
	productService aumo.ProductService
	shopService    aumo.ShopService
	auth           *auth.Authenticator
	validator      *validator.Validate
	verifier       *verifications.Verifier
	decoder        *form.Decoder
	translator     ut.Translator
	backendURL     string
}

// New returns a new instance of aumo's REST API
func New(c *Config) *Rest {
	switch {
	case c.UserService == nil:
		panic("rest: UserService not provided")
	case c.ReceiptService == nil:
		panic("rest: ReceiptService not provided")
	case c.OrderService == nil:
		panic("rest: OrderService not provided")
	case c.ProductService == nil:
		panic("rest: ProductService not provided")
	case c.ShopService == nil:
		panic("rest: ShopService not provided")
	case c.Auth == nil:
		panic("rest: Authenticator not provided")
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

	err := enTrans.RegisterDefaultTranslations(validator, trans)
	if err != nil {
		panic(err)
	}

	rest := &Rest{
		router:         r,
		userService:    c.UserService,
		receiptService: c.ReceiptService,
		orderService:   c.OrderService,
		productService: c.ProductService,
		shopService:    c.ShopService,
		auth:           c.Auth,
		verifier:       c.Verifier,
		backendURL:     c.BackendURL,
		validator:      validator,
		translator:     trans,
		decoder:        decoder,
	}

	r.Use(
		middleware.RequestID,
		middleware.RedirectSlashes,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		RateLimit(tollbooth.NewLimiter(5, nil).SetOnLimitReached(rest.onRateLimit)),
		Security,
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
		}).Handler,
		middleware.Heartbeat(c.MountRoute+"/ping"),
	)

	rest.mount(c.MountRoute)

	return rest
}

func (rest *Rest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rest.router.ServeHTTP(w, r)
}
