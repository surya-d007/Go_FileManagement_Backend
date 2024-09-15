# 21BCE5685 SURYA D - Backend Task - Trademarkia®

## File Management System - Go , AWS S3 , AWS RDS Postgres , Docker , EC2 hosting , JWT , Cache , AWS CLI

## Features

1. User registration and login with password hashing.
2. File upload to AWS S3 with metadata storage in PostgreSQL.
3. Search functionality with caching for file metadata.
4. Automatic cleanup of expired files from S3 and PostgreSQL.
5. Shareable public links for files.
6. DB queries cache

## Operations

### 1. POST - /register - http://52.66.239.215/register

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

### 2. POST /login - http://52.66.239.215/login

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

### 3. POST /upload: - http://52.66.239.215/upload

> Description: DesUpload a file to S3 and store metadata in PostgreSQL and get filE url as response
>
> **body : Content-Type: multipart/form-data **
>
> ```bash
> {
>
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

### 4. GET /files/{email}: - http://52.66.239.215/files/{userEmail}

> Description: Retrieve metadata for files uploaded by a specific email.
> Sample : http://52.66.239.215/files/sample@gmail.com > **Response**
>
> ```bash
> [
>    {
>        "ID": 26,
>        "Filename": "free-parking.png",
>        "URL": "https://file-upload-bucket-surya-aws.s3.ap-south-1.amazonaws.com/uploads/free-parking.png",
>        "Size": 12241,
>        "UploadDate": "2024-09-15T07:07:28.706456Z",
>        "Email": "sample@gmail.com"
>    },
>    {
>        "ID": 27,
>        "Filename": "aaaa.pdf",
>        "URL": "https://file-upload-bucket-surya-aws.s3.ap-south-1.amazonaws.com/uploads/aaaa.pdf",
>        "Size": 198520,
>        "UploadDate": "2024-09-15T07:07:38.317639Z",
>        "Email": "sample@gmail.com"
>    }
> ]
>
> ```

### 5. GET /searchFiles: - http://52.66.239.215/searchFiles?filename={FileName}&file_type={FileType}&upload_date={yyyy-mm-dd}

> #### like this u can combine any combination in the Query params
>
> ### sample search - http://52.66.239.215/searchFiles?filename=ab&file_type=pdf&upload_date=2024-09-15
>
> ### Only type http://52.66.239.215/searchFiles?file_type=pdf
>
> ### Onnly date - http://52.66.239.215/searchFiles?upload_date=2024-09-15
>
> ### Only Filename - http://52.66.239.215/searchFiles?filename=ab
>
> ### File name and type - http://52.66.239.215/searchFiles?filename=ab&file_type=pdf

> Description: Search for files with optional query parameters (filename, upload_date, file_type).
>
> **Query Params :**
>
> ```bash
>
>    "filename":"abcd"
>    "file_type":"pdf"
>    "upload_date":"YYYY-MM-DD"
>
> ```
>
> **Response**
>
> ```bash
> [
>    {
>        "ID": 6,
>        "Filename": "abcd (1).pdf",
>        "URL": "https://file-upload-bucket-surya-aws.s3.ap-south-1.amazonaws.com/uploads/abcd (1).pdf",
>        "Size": 69249,
>        "UploadDate": "2024-09-15T08:00:41.729401Z",
>        "Email": "surya"
>    },
>    {
>        "ID": 10,
>        "Filename": "Database Details - RDS Management Console.pdf",
>        "URL": "https://file-upload-bucket-surya-aws.s3.ap-south-1.amazonaws.com/uploads/Database Details - RDS Management Console.pdf",
>        "Size": 519341,
>        "UploadDate": "2024-09-15T08:05:59.113392Z",
>        "Email": "surya"
>    }
> ]
> ```

### 6. GET /share/{file_id}: - http://52.66.239.215/share/27

> Description: Generate a public link for a file using the id assigned by Postgres
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
