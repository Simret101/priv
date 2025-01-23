package main

import (
	"log"
	"net/http"
)

func userServiceHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserService is healthy!"))
}

func main() {
	http.HandleFunc("/health", userServiceHealth)
	log.Println("UserService running on :8082")
	http.ListenAndServe(":8082", nil)
}
