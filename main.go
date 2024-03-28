package main

import (
    "github.com/piyushkumar/authenticationmayursir/handlers"
	"github.com/piyushkumar/authenticationmayursir/middleware"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
    r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
    r.Handle("/protected", middleware.Authenticate(http.HandlerFunc(ProtectedEndpoint))).Methods("GET")

    log.Println("Server starting on :8080...")
    http.ListenAndServe(":8080", r)
}

// ProtectedEndpoint is a demo endpoint to demonstrate protected route
func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Access granted."))
}
