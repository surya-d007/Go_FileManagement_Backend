# 21BCE5685 SURYA D - Backend Task - Trademarkia

> ## File Management System - Go , AWS S3 , AWS RDS Postgres , Docker , EC2 hosting , JWT , Cache , AWS CLI , Cron Job

> ### Docker hub - https://hub.docker.com/r/suryad007/backendgoapp1/tags

## Server base URL - http://52.66.239.215

> ## Note : It is http

## Features

1. User registration and login with password hashing & JWT.
2. File upload to AWS S3 with metadata storage in PostgreSQL.
3. Search functionality with caching for file metadata.
4. Automatic cleanup of expired files from S3 and PostgreSQL.
5. Shareable public links for files.
6. DB queries cache
7. Corn job which runs once in 24 hrs
8. TestApi Scripts
9. Dockerization

## Prerequisites

Before running the Go server, ensure you have the following tools installed and configured:

> #### 1. Install Go
>
> #### 2. Install AWS CLI
>
> #### 3. Configure AWS CLI
>
> ```bash
> aws configure
> AWS Access Key ID [None]: YOUR_ACCESS_KEY
> AWS Secret Access Key [None]: YOUR_SECRET_KEY
> Default region name [None]: YOUR_REGION (e.g., ap-south-1)
> ```
>
> #### 4. PostgresSQL in cloud or Local

## Background Cron Job

> #### The application includes a background job that runs once in 24hrs to delete expired files from S3 and PostgreSQL. The job is started automatically when the application starts.

## Caching

> #### A simple in-memory cache stores PostgreSQL query search results, enhancing performance with a 5-minute Time-To-Live (TTL).

#

#

# API info:

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
> #### sample search - http://52.66.239.215/searchFiles?filename=ab&file_type=pdf&upload_date=2024-09-15
>
> #### Only type http://52.66.239.215/searchFiles?file_type=pdf
>
> #### Onnly date - http://52.66.239.215/searchFiles?upload_date=2024-09-15
>
> #### Only Filename - http://52.66.239.215/searchFiles?filename=ab
>
> #### File name and type - http://52.66.239.215/searchFiles?filename=ab&file_type=pdf

> Description: Search for files with optional query parameters (filename, upload_date, file_type).
>
> **Query Params :**
>
> ```bash
>
>    "filename":"ab"
>    "file_type":"pdf"
>    "upload_date":"2024-09-15"
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
> Response
>
> ```bash
> {
>    "fileURL": "https://file-upload-bucket-surya-aws.s3.ap-south-1.amazonaws.com/uploads/aaaa.pdf"
> }
> ```

To add a description of your PostgreSQL database schema and details to your `README.md` file in GitHub, you can structure it like this:

--

# Database Schema

### Tables

The PostgreSQL database contains two main tables: `file_metadata` and `users`.

#### 1. `file_metadata`

This table stores metadata for uploaded files, including the file name, URL, size, upload date, and the email of the user who uploaded the file.

| Column Name   | Data Type   | Description                                        |
| ------------- | ----------- | -------------------------------------------------- |
| `id`          | `INT`       | Unique identifier for each file record             |
| `filename`    | `TEXT`      | The name of the uploaded file                      |
| `url`         | `TEXT`      | The URL where the file is stored (e.g., S3 bucket) |
| `size`        | `INT`       | Size of the file in bytes                          |
| `upload_date` | `TIMESTAMP` | The timestamp when the file was uploaded           |
| `email`       | `TEXT`      | Email of the user who uploaded the file            |

**Sample Data:**

```
 id |                   filename                    |                         url                          |  size   |        upload_date         |      email
----+-----------------------------------------------+------------------------------------------------------|---------+----------------------------+------------------
  6 | abcd (1).pdf                                  | https://file-upload-bucket-surya-aws.s3.ap-south-1...|   69249 | 2024-09-15 08:00:41.729401 | sample@gmail.com
 10 | Database Details - RDS Management Console.pdf | https://file-upload-bucket-surya-aws.s3.ap-south-1...|  519341 | 2024-09-15 08:05:59.113392 | sample@gmail.com
```

#### 2. `users`

This table stores user authentication information, including email and hashed password.

| Column Name | Data Type | Description                             |
| ----------- | --------- | --------------------------------------- |
| `id`        | `INT`     | Unique identifier for each user         |
| `email`     | `TEXT`    | User's email address                    |
| `password`  | `TEXT`    | Hashed password for user authentication |

**Sample Data:**

```
 id |       email       |                           password
----+-------------------+--------------------------------------------------------------
  1 | sample@gmail.com | $2a$10$4zO.Iie2Ta.KrvpC9kxZcuezfTeRyVmQbJldo5BZ.etCP6hbs6nm6
  2 | sample@gmail.com | $2a$10$faDjlMv1XgATo020ExsCWOx0SgrcrMILTN7mqKkyVvIexV3FBEAmy
```

### Go Structs

The Go code interacting with this database uses the following structs for the tables:

```go
type FileMetadata struct {
    ID         int
    Filename   string
    URL        string
    Size       int
    UploadDate time.Time
    Email      string
}

type User struct {
    ID       int    `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
}
```
