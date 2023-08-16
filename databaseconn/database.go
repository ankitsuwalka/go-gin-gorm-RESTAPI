package databaseconn

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dbstring := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"

	if err1 := godotenv.Load(".env"); err1 != nil {
		log.Fatal(err1)
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER1")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")
	timezone := os.Getenv("TIMEZONE")

	dsn := fmt.Sprintf(dbstring, host, user, password, dbname, port, sslmode, timezone)

	db, err1 := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err1 != nil {
		log.Fatal(err1)
	}

	return db

}
