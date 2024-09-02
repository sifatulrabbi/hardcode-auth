package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       string `gorm:"primaryKey;not null"`
	Email    string `gorm:"uniqueIndex;size:300;not null"`
	Name     string `gorm:"type:varchar(100);not null"`
	Passowrd string `gorm:"not null" json:"omitempty"`
}

// removes sensitive fields like "password" from the user
func (u User) Truncate() User {
	tu := u
	tu.Passowrd = ""
	return tu
}

func (u *User) Update() error {
	db := GetDB()
	tx := db.Save(u)
	return tx.Error
}
