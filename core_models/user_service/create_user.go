package user_service

type CreateUserEntry struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type CreateUserEntryResponse struct {
	ID int `json:"id"`
}
