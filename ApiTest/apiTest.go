package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"BackEnd_21BCE5685/controllers"
)

func TestRegister(t *testing.T) {
	// Define a sample registration payload
	registrationPayload := map[string]string{
		"email":    "user@example.com",
		"password": "password123",
	}

	// Convert the payload to JSON
	registrationPayloadJSON, _ := json.Marshal(registrationPayload)

	// Create a new HTTP request for /register
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registrationPayloadJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler for the /register route
	handler := http.HandlerFunc(controllers.Register)

	// Call the handler with the request and the response recorder
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	// Check if the response contains the message "Registration successful"
	expected := `"message":"Registration successful"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("Expected response to contain %v, got %v", expected, rr.Body.String())
	}
}

func TestLogin(t *testing.T) {
	// Define a sample login payload
	loginPayload := map[string]string{
		"email":    "user@example.com",
		"password": "password123",
	}

	// Convert the payload to JSON
	loginPayloadJSON, _ := json.Marshal(loginPayload)

	// Create a new HTTP request for /login
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginPayloadJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler for the /login route
	handler := http.HandlerFunc(controllers.Login)

	// Call the handler with the request and the response recorder
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	// Check if the response contains the message "Login successful"
	expected := `"message":"Login successful"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("Expected response to contain %v, got %v", expected, rr.Body.String())
	}
}

func TestLoginInvalid(t *testing.T) {
	// Define an invalid login payload
	loginPayload := map[string]string{
		"email":    "user@example.com",
		"password": "wrongpassword",
	}

	// Convert the payload to JSON
	loginPayloadJSON, _ := json.Marshal(loginPayload)

	// Create a new HTTP request for /login
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginPayloadJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler for the /login route
	handler := http.HandlerFunc(controllers.Login)

	// Call the handler with the request and the response recorder
	handler.ServeHTTP(rr, req)

	// Check if the status code is 401 Unauthorized
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Expected status code %v, got %v", http.StatusUnauthorized, status)
	}

	// Check if the response contains the message "Invalid credentials"
	expected := `"Invalid credentials"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("Expected response to contain %v, got %v", expected, rr.Body.String())
	}
}

func TestSearchFiles(t *testing.T) {
	// Define a sample search query
	queryParams := "filename=testfile&upload_date=2023-09-01&file_type=pdf"

	// Create a new HTTP request for /searchFiles
	req, err := http.NewRequest("GET", "/searchFiles?"+queryParams, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and initialize routes

	// Create a handler for the /searchFiles route
	handler := http.HandlerFunc(controllers.SearchFiles)

	// Call the handler with the request and the response recorder
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	// Check if the response contains the expected search results
	expectedSubstring := `"filename":"testfile"`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expectedSubstring)) {
		t.Errorf("Expected response to contain %v, got %v", expectedSubstring, rr.Body.String())
	}
}
