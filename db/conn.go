package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB = nil

func NewConnection() {
	db, err := gorm.Open(sqlite.Open(".locals/local.db"), &gorm.Config{})
	if err != nil {
		log.Panicln("failed to connect to db:", err)
	}
	log.Println("Connected to the database")
	dbConn = db

	log.Println("Performing migrations...")
	// migrate all schemas
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Session{})
	log.Println("Migrations complete")
}

func GetDB() *gorm.DB {
	if dbConn == nil {
		log.Panicln("No database connection found. Please establish a db connection first.")
	}
	return dbConn
}
