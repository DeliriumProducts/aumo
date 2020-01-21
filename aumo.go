// Package aumo provides three things:
//    type User struct // domain types
//    type UserService interface // business related logic - contains validation and extra logic
//    type UserStore interface // low level dumb db related logic - it has no validation
// Implemenentations of different stores belong in their respective packages, such as MySQL.
// Implemenentations of actual services belong in vertical slices (grouping by feature)
//    ├── auth # auth logic
//    │   └── auth.go
//    ├── cmd # entrypoints
//    │   └── aumo
//    │       └── main.go
//    ├── data.sql
//    ├── go.mod
//    ├── go.sum
//    ├── mysql # mysql implementation of stores
//    │   ├── mysql.go
//    │   ├── order.go
//    │   ├── product.go
//    │   ├── receipt.go
//    │   ├── schema.go
//    │   └── user.go
//    ├── net
//    │   └── http
//    │       └── rest # rest api
//    │           ├── form.go
//    │           ├── json.go
//    │           ├── middlewares.go
//    │           ├── order_handlers.go
//    │           ├── products_handlers.go
//    │           ├── rest.go
//    │           ├── routes.go
//    │           └── user_handlers.go
//    ├── order.go
//    ├── ordering # ordering layer
//    │   └── service.go
//    ├── product.go
//    ├── products # products layer
//    │   └── service.go
//    ├── receipt # receipts layer
//    │   └── service.go
//    ├── receipt.go
//    ├── role.go
//    ├── store.go
//    ├── tests
//    │   ├── mysql.go
//    │   ├── order_service_test.go
//    │   ├── product_service_test.go
//    │   ├── receipt_service_test.go
//    │   ├── user_service_test.go
//    │   └── user_test.go
//    ├── user.go
//    └── users # users layer
//        └── service.go
package aumo
