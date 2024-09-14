package controllers

import (
	"BackEnd_21BCE5685/db"
	"BackEnd_21BCE5685/models"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// deleteExpiredFiles deletes files older than 1 minute from S3 and removes their metadata from PostgreSQL
func deleteExpiredFiles() {
	for {
		startTime := time.Now()
		log.Printf("[%s] Starting expired file cleanup job...\n", startTime.Format(time.RFC3339))

		// Define the total sleep duration
		sleepDuration := 1 * time.Minute
		interval := 1 * time.Second

		// Loop to log remaining time every second
		for remainingTime := sleepDuration; remainingTime > 0; remainingTime -= interval {
			time.Sleep(interval)
			log.Printf("[%s] Time remaining until next execution: %s\n", time.Now().Format(time.RFC3339), remainingTime)
		}

		// Calculate and log the actual time spent before execution
		timeElapsed := time.Since(startTime)
		timeRemaining := sleepDuration - timeElapsed
		if timeRemaining < 0 {
			timeRemaining = 0
		}
		log.Printf("[%s] Time remaining until next execution after processing: %s\n", time.Now().Format(time.RFC3339), timeRemaining)

		// Create a new AWS S3 session
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(region),
		})
		if err != nil {
			log.Printf("[%s] Failed to create AWS session: %v\n", time.Now().Format(time.RFC3339), err)
			continue
		}
		svc := s3.New(sess)

		// Get the current time
		now := time.Now()
		log.Printf("[%s] Current time: %s\n", now.Format(time.RFC3339), now.Format(time.RFC3339))

		// Get the list of expired files from the database
		expiredFiles, err := models.GetExpiredFiles(db.DB, now.Add(-1*time.Minute))
		if err != nil {
			log.Printf("[%s] Failed to get expired files: %v\n", time.Now().Format(time.RFC3339), err)
			continue
		}

		if len(expiredFiles) == 0 {
			log.Printf("[%s] No expired files found to delete.\n", time.Now().Format(time.RFC3339))
		}

		// Delete files from S3 and remove metadata from the database
		for _, file := range expiredFiles {
			log.Printf("[%s] Processing file: %s\n", time.Now().Format(time.RFC3339), file.Filename)

			// Delete the file from S3
			_, err := svc.DeleteObject(&s3.DeleteObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String("uploads/" + file.Filename),
			})
			if err != nil {
				log.Printf("[%s] Failed to delete file from S3: %v\n", time.Now().Format(time.RFC3339), err)
				continue
			}
			log.Printf("[%s] Successfully deleted file from S3: %s\n", time.Now().Format(time.RFC3339), file.Filename)

			// Remove the file metadata from the database
			if err := models.DeleteFileMetadata(db.DB, file.ID); err != nil {
				log.Printf("[%s] Failed to delete file metadata: %v\n", time.Now().Format(time.RFC3339), err)
			} else {
				log.Printf("[%s] Successfully deleted file metadata for ID: %d\n", time.Now().Format(time.RFC3339), file.ID)
			}
		}

		log.Printf("[%s] Expired file cleanup job completed.\n", time.Now().Format(time.RFC3339))
	}
}

// StartBackgroundJob starts the background job for deleting expired files
func StartBackgroundJob() {
	go deleteExpiredFiles()
}
