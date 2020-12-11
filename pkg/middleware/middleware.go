package middleware

import (
	"net/http"
)

type Middleware func(h http.HandlerFunc) http.HandlerFunc


func ApplyMiddleware(h http.HandlerFunc, mdw ...Middleware) http.HandlerFunc {
	// If there are now middlewares, use the handler.
	if len(mdw) == 0 {
		return h
	}

	// Otherwise nest the middlewares.
	return mdw[0](ApplyMiddleware(h, mdw[1:cap(mdw)]...))
}
