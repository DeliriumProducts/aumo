package main

import (
	"fmt"

	"github.com/deliriumproducts/aumo/auth"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/net/http/rest"
	"github.com/deliriumproducts/aumo/ordering"
	"github.com/deliriumproducts/aumo/products"
	"github.com/deliriumproducts/aumo/receipt"
	"github.com/deliriumproducts/aumo/users"
	"github.com/go-chi/docgen"
)

func main() {
	ps := mysql.NewProductStore(nil)
	os := mysql.NewOrderStore(nil)
	rs := mysql.NewReceiptStore(nil)
	us := mysql.NewUserStore(nil)
	auth := auth.New(nil, us, 0)

	r := rest.New(rest.Config{
		UserService:    users.New(us),
		ReceiptService: receipt.New(rs, us),
		OrderService:   ordering.New(os, ps, us),
		ProductService: products.New(ps),
		Auth:           auth,
		MountRoute:     "/api/v1",
	})

	fmt.Println(r.MarkdownRoutesDoc(docgen.MarkdownOpts{
		ProjectPath: "github.com/deliriumproducts/aumo",
		Intro:       "Welcome to aumo's REST API documentation",
	}))
}
