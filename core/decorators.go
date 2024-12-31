package core

import (
	"fmt"
	"net/http"
)

// RouteDecorator is a function that can decorate HTTP handlers
type RouteDecorator func(http.Handler) http.Handler

// WithAuth adds authentication to a route
func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add authentication logic here
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Guard decorator for route guards
func Guard(guardFunc func(*http.Request) bool) RouteDecorator {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !guardFunc(r) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
