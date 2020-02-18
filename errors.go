package aumo

import "errors"

var (
	// ErrNotSufficientPoints is an error for when the user doens't have enough points
	ErrNotSufficientPoints = errors.New("aumo: user doesn't have enough points to buy this item")
	// ErrNotInStock is an error for when an item isn't in stock
	ErrNotInStock = errors.New("aumo: shop doesn't have enough stock of the item")
	// ErrInvalidPassword is an error for when the user provided an invalid password
	ErrInvalidPassword = errors.New("aumo: wrong password")
	// ErrUserNotFound is an error for when a user hasn't been found
	ErrUserNotFound = errors.New("aumo: user not found")
	// ErrDuplicateEmail is an error for when a user tries to register with an already existing email
	ErrDuplicateEmail = errors.New("aumo: duplicate email")
	// ErrNotVerified is an error for when a user isn't verified
	ErrNotVerified = errors.New("aumo: user is not verified")
	// ErrUserAlreadySet is an error for when a user has already claimed a receipt
	ErrUserAlreadySet = errors.New("aumo: this receipt has already been claimed")
	// ErrReceiptUserNotExist is an error for when a user doesn't exist when trying to claim a receipt
	ErrReceiptUserNotExist = errors.New("aumo: can't claim a receipt for a user that doesn't exist")
	// ErrReceiptNotFound when a receipt isn't found
	ErrReceiptNotFound = errors.New("aumo: receipt not found")
	// ErrProductNotFound when a receipt isn't found
	ErrProductNotFound = errors.New("aumo: product not found")
	// ErrOrderProductNotFound is an error for when a user places an order on a product that doesn't exist
	ErrOrderProductNotFound = errors.New("aumo: can't place an order for a non existing product")
	// ErrOrderUserNotFound is an error for when a user doesn't exist when placing an order
	ErrOrderUserNotFound = errors.New("aumo: can't place an order for a user that doesn't exist")
	// ErrOrderNotFound when a receipt isn't found
	ErrOrderNotFound = errors.New("aumo: order not found")
	// ErrShopOwnerUserNotFound is an error for when a user doesn't exist when using him as a shop owner
	ErrShopOwnerUserNotFound = errors.New("aumo: can't use a non-existing user as a shop owner")
	// ErrShopNotFound is an error for when a shop wasn't found
	ErrShopNotFound = errors.New("aumo: shop not found")
)
