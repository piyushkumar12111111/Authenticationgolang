package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/piyushkumar/authenticationmayursir/models"
)

var JwtKey = []byte("your_secret_key")

// Credentials used for login and registration
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Claims struct will add username as a claim to the token
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// RegisterUser creates a new user
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

// LoginUser logs in a user and returns a token
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Verify the username and password
    if !models.AuthenticateUser(creds.Username, creds.Password) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Create a new token for the user
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

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })
}
