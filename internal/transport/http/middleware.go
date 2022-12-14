package http

import (
	"context"

	"net/http"
	"time"

	"github.com/CallumKerrEdwards/loggerrific"
)

type Middlewares struct {
	loggerrific.Logger
}

func NewMiddlewares(logger loggerrific.Logger) *Middlewares {
	return &Middlewares{Logger: logger}
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		next.ServeHTTP(w, r)
	})
}

func (m *Middlewares) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Logger.WithFields(map[string]interface{}{"method": r.Method, "path": r.URL.Path}).Infoln("Handled Request")
		next.ServeHTTP(w, r)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
