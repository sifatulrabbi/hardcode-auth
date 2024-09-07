package tests

import (
	"fmt"
	"testing"

	"github.com/sifatulrabbi/hardcode-auth/db"
)

var connConfig = db.DBConnConfig{DBPath: "./test.db"}

func TestUserCRUD(t *testing.T) {
	db.NewConnection(&connConfig)

	tu1 := db.User{
		Name:     "Test User 1",
		Email:    "user1@test.test",
		Password: "password",
	}
	tu2 := db.User{
		Name:     "Test User 2",
		Email:    "user2@test.test",
		Password: "password",
	}

	fmt.Println("testing create user...")
	if err := tu1.Create(); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	if err := tu2.Create(); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	fmt.Println("create user success")
	fmt.Println("testing get user")
}
