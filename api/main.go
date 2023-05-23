package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// defining the handlers directory
	"github.com/Rahat7288/simple-go-api/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
