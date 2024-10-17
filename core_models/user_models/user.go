package user_models

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first-name"`
	LastName       string `json:"last-name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	CreatedAt      string `json:"created-at"`
	FollowersCount int    `json:"followers-count"`
	FollowingCount int    `json:"following-count"`
}
