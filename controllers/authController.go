package controllers

import (
	"BackEnd_21BCE5685/models"
	"BackEnd_21BCE5685/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	Token string `json:"token"`
}

// Register new user
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save user to DB
	err = user.Create()
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login user
func Login(w http.ResponseWriter, r *http.Request) {
	var reqUser models.User
	json.NewDecoder(r.Body).Decode(&reqUser)

	// Find user by email
	user, err := models.GetUserByEmail(reqUser.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{Token: token})
}
