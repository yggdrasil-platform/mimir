package middleware

import (
	"github.com/kieranroneill/mimir/pkg/logger"
  "net/http"
)

func LogRequest() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.Info.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
			next.ServeHTTP(w, r)
		}
	}
}
