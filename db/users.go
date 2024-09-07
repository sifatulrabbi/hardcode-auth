package db

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       string `gorm:"primaryKey;not null"`
	Email    string `gorm:"uniqueIndex;size:300;not null"`
	Name     string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"not null" json:"omitempty"`
}

// removes sensitive fields like "password" from the user
func (u User) Truncate() User {
	tu := u
	tu.Password = ""
	return tu
}

func (u *User) Update() error {
	db := GetDB()
	tx := db.Save(u)
	return tx.Error
}

func (u *User) Create() error {
	u.ID = uuid.New().String()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return errors.New("unable to hash the password")
	}
	u.Password = string(hashedPass)
	return u.Update()
}

func (u *User) FindById(id string) *User {
	db := GetDB()
	user := User{}
	if err := db.Find(&user, "id = ?").Error; err != nil {
		log.Println("Unable to get the user:", err)
	}
	return &user
}
