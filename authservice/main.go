package main

import (
	"log"
	"net/http"
)

func authServiceHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("AuthService is healthy!"))
}

func main() {
	http.HandleFunc("/health", authServiceHealth)
	log.Println("AuthService running on :8083")
	http.ListenAndServe(":8083", nil)
}
