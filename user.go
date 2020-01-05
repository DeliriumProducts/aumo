package aumo

// User represents a user of aumo
type User struct {
	Model
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"-" gorm:"not null"`
	Avatar   string    `json:"avatar" gorm:"not null"`
	Points   float64   `json:"points" gorm:"not null"`
	Orders   []Order   `json:"orders"`
	Receipts []Receipt `json:"receipts"`
}

// UserService contains all `User`
// related business logic
type UserService interface {
	User(id uint) *User
	Users() []User
	Save(*User)
	Update(*User)
	Delete(*User)
}
