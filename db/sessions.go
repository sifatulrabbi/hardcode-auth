package db

type Session struct {
	ID        string `json:"id"`
	Exp       string `json:"exp"`
	CreatedAt string `json:"createdAt"`
}
