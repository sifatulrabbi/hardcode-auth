package db

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserPrivate struct {
	User
	Passowrd string `json:"passowrd"`
}
