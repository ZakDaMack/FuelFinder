package api

import (
	"log/slog"
	"net/http"
	"time"
)

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("incoming request", "timestamp", time.Now().UTC(), "method", r.Method, "url", r.URL.String())
		next.ServeHTTP(w, r)
	})
}
