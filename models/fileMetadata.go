package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// FileMetadata represents the metadata for uploaded files
type FileMetadata struct {
	ID         int
	Filename   string
	URL        string
	Size       int
	UploadDate time.Time
	Email      string
}

// SaveFileMetadata saves the file metadata to the PostgreSQL database
func SaveFileMetadata(db *sql.DB, filename, fileURL string, size int, email string) error {
	query := `INSERT INTO file_metadata (filename, url, size, upload_date, email) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, filename, fileURL, size, time.Now(), email)
	return err
}


func GetFileMetadataByID(db *sql.DB, id int) (*FileMetadata, error) {
	query := `SELECT id, filename, url, size, upload_date, email FROM file_metadata WHERE id = $1`
	row := db.QueryRow(query, id)
	
	var metadata FileMetadata
	if err := row.Scan(&metadata.ID, &metadata.Filename, &metadata.URL, &metadata.Size, &metadata.UploadDate, &metadata.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("file metadata not found")
		}
		return nil, err
	}
	return &metadata, nil
}

func GetFileMetadataByEmail(db *sql.DB, email string) ([]*FileMetadata, error) {
    query := `SELECT id, filename, url, size, upload_date, email FROM file_metadata WHERE email = $1`
    rows, err := db.Query(query, email)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var metadataList []*FileMetadata
    for rows.Next() {
        var metadata FileMetadata
        if err := rows.Scan(&metadata.ID, &metadata.Filename, &metadata.URL, &metadata.Size, &metadata.UploadDate, &metadata.Email); err != nil {
            return nil, err
        }
        metadataList = append(metadataList, &metadata)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    if len(metadataList) == 0 {
        return nil, fmt.Errorf("no file metadata found for email %s", email)
    }

    return metadataList, nil
}

func SearchFiles(db *sql.DB, filename string, uploadDate *time.Time, fileType string) ([]*FileMetadata, error) {
	var conditions []string
	var params []interface{}
	query := "SELECT id, filename, url, size, upload_date, email FROM file_metadata WHERE 1=1"

	paramIndex := 1 // PostgreSQL uses $1, $2, etc. for placeholders

	// Add conditions based on the provided filters
	if filename != "" {
		conditions = append(conditions, fmt.Sprintf("LOWER(filename) LIKE $%d", paramIndex))
		params = append(params, "%"+strings.ToLower(filename)+"%")
		paramIndex++
	}
	if uploadDate != nil {
		conditions = append(conditions, fmt.Sprintf("DATE(upload_date) = $%d", paramIndex))
		params = append(params, *uploadDate)
		paramIndex++
	}
	if fileType != "" {
		conditions = append(conditions, fmt.Sprintf("LOWER(filename) LIKE $%d", paramIndex))
		params = append(params, "%."+strings.ToLower(fileType))
		paramIndex++
	}

	// Append conditions to the query if any
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	// Execute the query
	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan the results into a slice of FileMetadata
	var metadataList []*FileMetadata
	for rows.Next() {
		var metadata FileMetadata
		if err := rows.Scan(&metadata.ID, &metadata.Filename, &metadata.URL, &metadata.Size, &metadata.UploadDate, &metadata.Email); err != nil {
			return nil, err
		}
		metadataList = append(metadataList, &metadata)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// If no records are found, return an error
	if len(metadataList) == 0 {
		return nil, fmt.Errorf("no files found for the given criteria")
	}

	return metadataList, nil
}




// GetExpiredFiles retrieves files that are older than the given cutoff time
func GetExpiredFiles(db *sql.DB, cutoffTime time.Time) ([]*FileMetadata, error) {
	query := "SELECT id, filename, url, size, upload_date, email FROM file_metadata WHERE upload_date <= $1"
	rows, err := db.Query(query, cutoffTime)
	if err != nil {
		return nil, fmt.Errorf("failed to query expired files: %w", err)
	}
	defer rows.Close()

	var expiredFiles []*FileMetadata
	for rows.Next() {
		var file FileMetadata
		if err := rows.Scan(&file.ID, &file.Filename, &file.URL, &file.Size, &file.UploadDate, &file.Email); err != nil {
			return nil, fmt.Errorf("failed to scan file metadata: %w", err)
		}
		expiredFiles = append(expiredFiles, &file)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return expiredFiles, nil
}

// DeleteFileMetadata removes file metadata from the database using file ID
func DeleteFileMetadata(db *sql.DB, fileID int) error {
	query := "DELETE FROM file_metadata WHERE id = $1"
	_, err := db.Exec(query, fileID)
	if err != nil {
		return fmt.Errorf("failed to delete file metadata: %w", err)
	}
	return nil
}
