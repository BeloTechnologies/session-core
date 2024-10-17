package user_models

type CreateUserRow struct {
	Username  string `json:"username"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// May be pointless, could remove?
type CreateUserRowResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
