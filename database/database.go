package database

import (
	"github.com/joho/godotenv"

	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func InitDb() {
	godotenv.Load()

	connStr := "user=" + os.Getenv("DB_USERNAME") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" host=127.0.0.1 port=5432 sslmode=disable"

	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func GetDB() *sql.DB {
	return db
}