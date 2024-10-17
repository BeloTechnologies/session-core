package user_models

type CreateUserRow struct {
	Username  string `json:"username"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type CreateUserRowResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
