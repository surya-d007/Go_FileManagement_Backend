package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	

	errr := godotenv.Load()
    if errr != nil {
        log.Fatal("Error loading .env file")
    }

    connStr := os.Getenv("DB_CONNECTION_STRING")
    if connStr == "" {
        log.Fatal("DB_CONNECTION_STRING environment variable not set")
    }
	
	
	//connStr := "CONNECTTION STRING TO UR POSTGRES SQL"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Could not ping the database:", err)
	}

	log.Println("Database connected successfully")
}
