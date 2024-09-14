package controllers

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"BackEnd_21BCE5685/db"
	"BackEnd_21BCE5685/models"
)

const (
	bucketName = "file-upload-bucket-surya-aws" // Replace with your S3 bucket name
	region     = "ap-south-1"      // Replace with your AWS region
)

// UploadFile handles file uploads to S3 and saves metadata in PostgreSQL
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the form data to get the file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Extract user email from form
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	// Create a new AWS S3 session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		http.Error(w, "Failed to create AWS session", http.StatusInternalServerError)
		return
	}

	svc := s3.New(sess)

	// Upload the file to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("uploads/" + fileHeader.Filename),
		Body:   file,
	})
	if err != nil {
		http.Error(w, "Failed to upload file to S3", http.StatusInternalServerError)
		return
	}

	// Generate the file URL
	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/uploads/%s", bucketName, region, fileHeader.Filename)

	// Save file metadata to the database
	go saveFileMetadata(fileHeader.Filename, fileURL, int(fileHeader.Size), email)

	// Return the file URL as response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"fileURL": "%s"}`, fileURL)))
}

// saveFileMetadata saves the file metadata to the PostgreSQL database
func saveFileMetadata(filename, fileURL string, size int, email string) {
	dbConn := db.DB
	if err := models.SaveFileMetadata(dbConn, filename, fileURL, size, email); err != nil {
		fmt.Printf("Failed to save file metadata: %v\n", err)
	}
}
