package http

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/CallumKerrEdwards/loggerrific"
	"github.com/gorilla/mux"

	"github.com/CallumKerrEdwards/library-api/internal/transport/http/auth"
	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

var (
	address = "0.0.0.0:8080"
)

type BookService interface {
	PostBook(context.Context, books.Book) (books.Book, error)
	GetBook(ctx context.Context, ID string) (books.Book, error)
	GetAllBooks(ctx context.Context) ([]books.Book, error)
}

type Handler struct {
	Router      *mux.Router
	Service     BookService
	Server      *http.Server
	Log         loggerrific.Logger
	AuthHandler *auth.Handler
}

func NewHandler(service BookService, logger loggerrific.Logger) *Handler {
	h := &Handler{
		Service: service,
		Log:     logger,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	m := NewMiddlewares(logger)

	h.Router.Use(JSONMiddleware)
	h.Router.Use(m.LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)
	h.AuthHandler = &auth.Handler{Log: logger}

	h.Server = &http.Server{
		Addr:              address,
		Handler:           h.Router,
		ReadHeaderTimeout: 15 * time.Second,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	h.Router.HandleFunc("/api/v1/auth/login", h.AuthHandler.Login).Methods("GET")
	h.Router.HandleFunc("/api/v1/auth/refresh", h.AuthHandler.Refresh).Methods("GET")
	h.Router.HandleFunc("/api/v1/auth/welcome", h.AuthHandler.Welcome).Methods("GET")

	h.Router.HandleFunc("/api/v1/book", auth.JWTAuth(h.PostBook)).Methods("POST")
	h.Router.HandleFunc("/api/v1/book", h.GetAllBooks).Methods("GET")
	h.Router.HandleFunc("/api/v1/book/{id}", h.GetBook).Methods("GET")
}

func (h *Handler) Serve() error {
	h.Log.Infoln("Starting server at", address)

	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			h.Log.WithError(err).Errorln("Server Error")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := h.Server.Shutdown(ctx)
	if err != nil {
		h.Log.WithError(err).Errorln("Problem shutting down server")
		return err
	}

	h.Log.Infoln("Shut down server gracefully")

	return nil
}
