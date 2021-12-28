package types

// User encapsulate a user in the system.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}
