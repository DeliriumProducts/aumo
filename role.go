package aumo

// Role is a user role
type Role string

const (
	// Admin is the owner
	Admin Role = "Admin"
	// Customer is the default role for any user
	Customer Role = "Customer"
)
