package middleware

import (
    "github.com/piyushkumar/authenticationmayursir/handlers"
    "net/http"

    "github.com/dgrijalva/jwt-go"
)

// Authenticate is a middleware that checks for a valid JWT token
func Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        c, err := r.Cookie("token")
        if err != nil {
            if err == http.ErrNoCookie {
                w.WriteHeader(http.StatusUnauthorized)
                return
            }
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        tknStr := c.Value
claims := &handlers.Claims{}

token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
	return handlers.JwtKey, nil   //! updated the error of jwtkey
})

if err != nil {
	if err == jwt.ErrSignatureInvalid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

        if !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
