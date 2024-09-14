package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	//connStr := "host=database-1.cfaiiyw4ghic.ap-south-1.rds.amazonaws.com port=5432 user=postgres password=Surya420 dbname=database-1"
	connStr := "host=database-1.cfaiiyw4ghic.ap-south-1.rds.amazonaws.com port=5432 user=postgres password=Surya420 dbname=postgres"

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
