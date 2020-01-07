package aumo

type Order struct {
	ID        uint `json:"id" db:"id,omitempty"`
	UserID    uint `json:"userID" db:"user_id"`
	ProductID uint `json:"productID" db:"product_id"`
}
