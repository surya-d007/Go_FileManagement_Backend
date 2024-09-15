package main

import (
	"BackEnd_21BCE5685/controllers"
	"BackEnd_21BCE5685/db"
	"BackEnd_21BCE5685/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	db.InitDB()

	controllers.StartBackgroundJob()
	// Initialize router
	router := routes.InitRoutes()

	// Start the server
	log.Println("Server started at port 80")
	log.Fatal(http.ListenAndServe(":80", router))
}
