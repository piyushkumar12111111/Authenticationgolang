package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/piyushkumar/authenticationmayursir/models"
)

var JwtKey = []byte("your_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.CreateUser(creds.Username, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !models.AuthenticateUser(creds.Username, creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//! set the token in the cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// ! for fetching all the users data
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersData := make(map[string]string)
	for username, user := range models.Users {
		usersData[username] = string(user.Password)
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(usersData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//! delete 

func DeleteUser(w http.ResponseWriter, r *http.Request) {
   
    vars := mux.Vars(r)
    username := vars["username"]

    
    if _, ok := models.Users[username]; !ok {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    //! Delete the user from the map
    delete(models.Users, username)

  
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User deleted successfully"))
}