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

	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Test for create user
	user := models.User{ID: "1", Username: "Rahat islam akash", Email: "example@mail.com", FirstName: "Rahat", LastName: "Akash", Address: "Rajbari"}
	userJSON, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Test for get users
	req, _ = http.NewRequest("GET", "/users", nil)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var users []models.User
	json.Unmarshal(rec.Body.Bytes(), &users)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, user.ID, users[0].ID)
	assert.Equal(t, user.Username, users[0].Username)
	assert.Equal(t, user.Email, users[0].Email)

	// Test for update user
	updatedUser := models.User{ID: "1", Username: "Rahat islam akash", Email: "example@mail.com", FirstName: "Rahat", LastName: "Akash", Address: "Rajbari"}
	updatedUserJSON, _ := json.Marshal(updatedUser)
	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(updatedUserJSON))
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var updatedUserResponse models.User
	json.Unmarshal(rec.Body.Bytes(), &updatedUserResponse)
	assert.Equal(t, updatedUser.Username, updatedUserResponse.Username)
	assert.Equal(t, updatedUser.Email, updatedUserResponse.Email)

	// Testing Delete User
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
