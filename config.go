package hardcodeauth

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

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
	ENVConfig.GO_ENV = getENV("GO_ENV", false, "development")

	if ENVConfig.GO_ENV != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Panicln(err)
		}
	}

	ENVConfig.JWT_SECRET = getENV("JWT_SECRET", true)
	ENVConfig.DB_NAME = getENV("DB_NAME", true)
	ENVConfig.DB_HOST = getENV("DB_HOST", true)
	ENVConfig.DB_HOST = getENV("DB_HOST", true)
	ENVConfig.DB_PASS = getENV("DB_PASS", true)
	ENVConfig.DB_USER = getENV("DB_USER", true)
	ENVConfig.PORT = getENV("PORT", true, "8000")
}

func getENV(name string, required bool, defaultVal ...string) string {
	v := os.Getenv(name)
	if required && v == "" {
		log.Panicf("%q env is required but not found\n", name)
	}
	if v == "" && len(defaultVal) > 0 {
		return strings.Join(defaultVal, "")
	}
	return v
}
