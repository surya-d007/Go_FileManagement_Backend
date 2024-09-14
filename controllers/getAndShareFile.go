package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	lru "github.com/hashicorp/golang-lru"

	"BackEnd_21BCE5685/db"
	"BackEnd_21BCE5685/models"
)

const (

	cacheSize  = 100                // Number of items to keep in cache
)

// In-memory cache for file metadata
var cache, _ = lru.New(cacheSize)

// RetrieveFileMetadata handles retrieval of file metadata




func RetrieveFileMetadata(w http.ResponseWriter, r *http.Request) {
	//email := r.URL.Query().Get("email")
	vars := mux.Vars(r)
	email := vars["email"]
    if email == "" {
        http.Error(w, "Email is required", http.StatusBadRequest)
        return
    }

    // Retrieve file metadata by email from the database
    dbConn := db.DB
    metadataList, err := models.GetFileMetadataByEmail(dbConn, email)
    if err != nil {
        http.Error(w, "Failed to retrieve file metadata: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the metadata as JSON
    writeJSONResponse(w, metadataList)
}



// ShareFile handles generating a public link for the file
func ShareFile(w http.ResponseWriter, r *http.Request) {
	// Parse the file ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["file_id"]
	if idStr == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	// Retrieve file metadata from database
	dbConn := db.DB
	metadata, err := models.GetFileMetadataByID(dbConn, id)
	if err != nil {
		http.Error(w, "Failed to retrieve file metadata", http.StatusInternalServerError)
		return
	}

	// Generate public link (assuming the file is publicly accessible)
	fileURL := metadata.URL

	// Optional: Implement URL expiration logic here

	// Return the file URL
	writeJSONResponse(w, map[string]string{"fileURL": fileURL})
}

// Utility function to write JSON response
func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
