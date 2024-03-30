package handlers

import (
	"net/http"
	"github.com/gorilla/sessions"
)

// LogoutHandler clears the session
func LogoutHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		delete(session.Values, "authenticated")
		session.Save(r, w)
		w.Write([]byte("Logged out successfully"))
	}
}
