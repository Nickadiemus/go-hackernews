package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	err = godotenv.Load(pwd + "/.env")
	if err != nil {
		log.Panic("Failed to load enironment variable file:", err)
	}
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// "password=%s dbname=%s sslmode=disable",
	// host, port, user, password, dbname)
	postgresConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println("postgresConfig", postgresConfig)
	db, err := sql.Open("postgres", postgresConfig)
	if err != nil {
		log.Panic(err)
	}
	Db = db
}
