package controllers

import (
	"BackEnd_21BCE5685/db"
	"BackEnd_21BCE5685/models"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// CacheItem structure for the search cache
type CacheItem struct {
	Data      []*models.FileMetadata
	ExpiresAt time.Time
}

// Global cache map for SearchFiles
var (
	Cache_search      = make(map[string]*CacheItem)
	cacheSearchMutex  = sync.RWMutex{}
	cacheSearchTTL    = 3 * time.Minute // Time-to-live for cache (3 minutes)
)

// generateCacheKey creates a unique cache key based on the search parameters
func generateCacheKey(filename string, uploadDate *time.Time, fileType string) string {
	var uploadDateStr string
	if uploadDate != nil {
		uploadDateStr = uploadDate.Format("2006-01-02")
	}
	return fmt.Sprintf("filename:%s|uploadDate:%s|fileType:%s", filename, uploadDateStr, fileType)
}

// isCacheValid checks if the cache is still valid
func isCacheValid(item *CacheItem) bool {
	return time.Now().Before(item.ExpiresAt)
}

// SearchFiles handles searching files based on filename, upload date, or file type with caching
func SearchFiles(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	filename := r.URL.Query().Get("filename")
	uploadDateStr := r.URL.Query().Get("upload_date")
	fileType := r.URL.Query().Get("file_type")

	// Parse upload date if provided
	var uploadDate *time.Time
	if uploadDateStr != "" {
		date, err := time.Parse("2006-01-02", uploadDateStr) // expecting yyyy-mm-dd format
		if err != nil {
			http.Error(w, "Invalid upload date format, expected yyyy-mm-dd", http.StatusBadRequest)
			return
		}
		uploadDate = &date
	}

	// Check if any search parameter is provided
	if filename == "" && uploadDate == nil && fileType == "" {
		http.Error(w, "At least one search parameter (filename, upload_date, or file_type) is required", http.StatusBadRequest)
		return
	}

	// Generate cache key based on search parameters
	cacheKey := generateCacheKey(filename, uploadDate, fileType)

	// Try to read from the search cache
	cacheSearchMutex.RLock()
	cachedItem, found := Cache_search[cacheKey]
	cacheSearchMutex.RUnlock()

	// If cache exists and is valid, return cached data
	if found && isCacheValid(cachedItem) {
		fmt.Printf("\nCache hit - Serving from cache for key: %s\n", cacheKey)
		writeJSONResponse(w, cachedItem.Data)
		return
	}

	// If not found in cache or cache is expired, fetch from the database
	
	dbConn := db.DB
	fmt.Printf("\nCache miss - Executing database query with parameters: filename=%s , uploadDate=%v, fileType=%s\n", filename, uploadDate, fileType)
	metadataList, err := models.SearchFiles(dbConn, filename, uploadDate, fileType)
	if err != nil {
		http.Error(w, "Failed to search files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Update cache with new data
	cacheSearchMutex.Lock()
	Cache_search[cacheKey] = &CacheItem{
		Data:      metadataList,
		ExpiresAt: time.Now().Add(cacheSearchTTL),
	}
	cacheSearchMutex.Unlock()

	// Return the metadata as JSON
	writeJSONResponse(w, metadataList)
}

