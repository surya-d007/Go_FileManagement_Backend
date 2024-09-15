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

### 1. Post - /register - http://52.66.239.215/register

> Description: Registers a new user with email and hashed password.
>
> body :
>
> ```bash
> {
>    "email":"sample@gmail.com",
>    "password":"samplepass"
> }
> ```
>
> Response
>
> ```
> Response:
> Status - 201 Created
> ```

### 2. Post /login - http://52.66.239.215/login

> Description: Logs in an existing user and returns a JWT token.
>
> **body :**
>
> ```bash
> {
>    "email":"sample@gmail.com",
>    "password":"samplepass"
> }
> ```
>
> **Response**
>
> ```bash
> {
>    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXBsZUBnZG1haWwuY29tIiwiZXhwIjoxNzI2NjQxNTQyfQ.oLIUq3OqlKk4AFc3UXrnnIhQQj38XNydW6drLYf00OE"
> }
> ```

### 3. POST /upload:

> Description: DesUpload a file to S3 and store metadata in PostgreSQL and get filE url as response
>
> **body : form-data **
>
> ```bash
> {
>    Key   : Value
>    file  type file : [Your file]
>    email type text : "sample@gmail.com"
> }
>
> ```
>
> **Response**
>
> ```bash
> {
>    "fileURL": "https://file-upload-bucket-surya-aws.s3.ap-south-1.amazonaws.com/uploads/Stack.pdf"
> }
> ```

4.  GET /searchFiles: Search for files with optional query parameters (filename, upload_date, file_type).
5.  GET /files/{email}: Retrieve metadata for files uploaded by a specific email.
6.  GET /share/{file_id}: Generate a public link for a file.
