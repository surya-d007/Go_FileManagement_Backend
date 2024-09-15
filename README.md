# 21BCE5685 SURYA D - Backend Task - Trademarkia®

## File Management System - Go , AWS S3 , Postgres , Docker , EC2 hosting

## Features

1. User registration and login with password hashing.
2. File upload to AWS S3 with metadata storage in PostgreSQL.
3. Search functionality with caching for file metadata.
4. Automatic cleanup of expired files from S3 and PostgreSQL.
5. Shareable public links for files.
6. DB queries cache in server

## Operations

1. Post /register

Description: Registers a new user with email and hashed password.

```bash
Request :
http://52.66.239.215/register
body :

{
    "email":"sample@gmail.com",
    "password":"samplepass"
}
```

2. Post /login
3. POST /upload: Upload a file to S3 and store metadata in PostgreSQL.
4. GET /searchFiles: Search for files with optional query parameters (filename, upload_date, file_type).
5. GET /files/{email}: Retrieve metadata for files uploaded by a specific email.
6. GET /share/{file_id}: Generate a public link for a file.
