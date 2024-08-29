package hardcodeauth

import "fmt"

var ENVConfig = struct {
	PORT       string
	GO_ENV     string
	DB_NAME    string
	DB_HOST    string
	DB_PASS    string
	DB_USER    string
	JWT_SECRET string
}{}

func prepareConfigs() {
	fmt.Println(ENVConfig)
}
