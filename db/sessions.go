package db

import "gorm.io/gorm"

type Session struct {
	gorm.Model

	ID        string `json:"id"`
	Exp       string `json:"exp"`
	CreatedAt string `json:"createdAt"`
}
