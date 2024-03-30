package main

import (
	"log"
	"net/http"

	"github.com/piyushkumar/authenticationmayursir/handlers"
	"github.com/piyushkumar/authenticationmayursir/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	r.Handle("/protected", middleware.Authenticate(http.HandlerFunc(ProtectedEndpoint))).Methods("GET")
	r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{username}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{username}", handlers.UpdateUser).Methods("PUT") //! method for updating user

	//! employee data apis

	r.HandleFunc("/reports", handlers.CreateReport).Methods("POST")
	r.HandleFunc("/reports/{id}", handlers.GetReport).Methods("GET")
	r.HandleFunc("/reports/{id}", handlers.UpdateReport).Methods("PUT")
	r.HandleFunc("/reports/{id}", handlers.DeleteReport).Methods("DELETE")

	//! retriving all reports

	r.HandleFunc("/reports", handlers.GetAllReports).Methods("GET")

	//! serach api
	r.HandleFunc("/reports/search", handlers.SearchReports).Methods("GET")

	//! library management system apis

	r.HandleFunc("/addbook", handlers.Addbook).Methods("POST")
	r.HandleFunc("/book/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/book/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", handlers.DeleteBook).Methods("DELETE")
	
	//! retriving all books
	r.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")

	

	log.Println("Server starting on :8085...")
	http.ListenAndServe(":8085", r)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Access granted."))
}
