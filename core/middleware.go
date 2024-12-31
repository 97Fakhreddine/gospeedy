package core

import "net/http"

// Middleware defines a structure for middleware
type Middleware struct {
	Handler func(http.Handler) http.Handler
}

// NewMiddleware creates new middleware
func NewMiddleware(handler func(http.Handler) http.Handler) *Middleware {
	return &Middleware{
		Handler: handler,
	}
}
