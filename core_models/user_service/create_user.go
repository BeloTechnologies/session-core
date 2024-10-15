package user_service

type CreateUserEntry struct {
	Username  string `json:"username"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type CreateUserEntryResponse struct {
	ID int `json:"id"`
}
