package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rahat7288/simple-go-api/handlers"
	"github.com/Rahat7288/simple-go-api/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var users []models.User

func TestCRUDAPI(t *testing.T) {
	// Create a new router for testing
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Test Create User endpoint
	user := models.User{ID: "1", Username: "Rahat islam akash", Email: "example@mail.com", FirstName: "Rahat", LastName: "Akash", Address: "Rajbari"}
	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Test Get User endpoint
	req, _ = http.NewRequest("GET", "/users/1", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var retrievedUser models.User
	json.Unmarshal(rec.Body.Bytes(), &retrievedUser)
	assert.Equal(t, user.ID, retrievedUser.ID)
	assert.Equal(t, user.Username, retrievedUser.Username)
	assert.Equal(t, user.Email, retrievedUser.Email)

	// Test Update User endpoint
	updatedUser := models.User{ID: "1", Username: "change my name", Email: "amar@mail.com", FirstName: "Akash", LastName: "Islam", Address: "Gazipur"}
	updatedUserJSON, _ := json.Marshal(updatedUser)
	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(updatedUserJSON))
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var updatedUserResponse models.User
	json.Unmarshal(rec.Body.Bytes(), &updatedUserResponse)
	assert.Equal(t, updatedUser.Username, updatedUserResponse.Username)
	assert.Equal(t, updatedUser.Email, updatedUserResponse.Email)

	// Test Delete User endpoint
	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify user deletion
	req, _ = http.NewRequest("GET", "/users", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var emptyUsers []models.User
	json.Unmarshal(rec.Body.Bytes(), &emptyUsers)
	assert.Equal(t, 0, len(emptyUsers))
}
