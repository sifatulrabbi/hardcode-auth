package main

import (
	hardcodeauth "github.com/sifatulrabbi/hardcode-auth"
	"github.com/sifatulrabbi/hardcode-auth/db"
)

func main() {
	db.NewConnection()
	api := hardcodeauth.New(db.GetDB())
	if err := api.StartAPI(); err != nil {
		panic(err)
	}
}
