package main

import (
	"log"
	"net/http"

	"github.com/piyushkumar/authenticationmayursir/handlers"
	"github.com/piyushkumar/authenticationmayursir/middleware"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)


var (
	// Key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	store = sessions.NewCookieStore([]byte("your-very-secure-key"))
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


	//! search api
	r.HandleFunc("/books/search", handlers.SearchBooks).Methods("GET")

	//! http://localhost:8085/books/search?bookname=Employee


	//! session based authentication

	r.HandleFunc("/sessionlogin", handlers.LoginHandler(store)).Methods("POST")
	r.HandleFunc("/sessionlogout", handlers.LogoutHandler(store)).Methods("POST")
    



	log.Println("Server starting on :8085...")
	http.ListenAndServe(":8085", r)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Access granted."))
}
