package handlers

import (
	"net/http"
	"github.com/gorilla/sessions"
)

// LoginHandler authenticates the user and creates a session
func LoginHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// For demonstration, we skip actual user authentication
		session, _ := store.Get(r, "session-name")
		session.Values["authenticated"] = true
		session.Save(r, w)
		w.Write([]byte("Logged in successfully"))
	}
}
