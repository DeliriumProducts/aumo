package aumo

// Role is a user role
type Role string

const (
	// Admin is the owner
	Admin Role = "Admin"
	// Customer is the default role for any user
	Customer Role = "Customer"
	// ShopOwner is a shop owner
	ShopOwner Role = "ShopOwner"
)
