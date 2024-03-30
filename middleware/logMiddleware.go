package middleware

import (
	"net/http"
	"github.com/gorilla/sessions"
)

// AuthMiddleware checks if the user's session is authenticated
func AuthMiddleware(next http.Handler, store *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
