package routes

import (
	"BackEnd_21BCE5685/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/upload", controllers.UploadFile).Methods("POST") // Added /upload route

	router.HandleFunc("/files/{email}", controllers.RetrieveFileMetadata).Methods("GET") // Retrieve file metadata
	router.HandleFunc("/share/{file_id:[0-9]+}", controllers.ShareFile).Methods("GET") // Share file URL
	router.HandleFunc("/searchFiles", controllers.SearchFiles).Methods("GET")


	return router
}
