package main

import (
	hardcodeauth "github.com/sifatulrabbi/hardcode-auth"
)

func main() {
	if err := hardcodeauth.StartAPI(); err != nil {
		panic(err)
	}
}
