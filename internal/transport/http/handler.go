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
	PostBook(context.Context, *books.Book) (books.Book, error)
	GetBook(ctx context.Context, id string) (books.Book, error)
	GetAllBooks(ctx context.Context) ([]books.Book, error)
	UpdateBook(ctx context.Context, id string, updatedBook *books.Book) (books.Book, error)
	DeleteBook(ctx context.Context, id string) error
	GetAllAudiobooks(ctx context.Context) ([]books.Book, error)
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

	h.Router.Use(JSONMiddleware)
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

	apiSubrouter := h.Router.PathPrefix("/api/v1").Subrouter()

	m := NewMiddlewares(h.Log)
	apiSubrouter.Use(m.LoggingMiddleware)

	apiSubrouter.HandleFunc("/auth/login", h.AuthHandler.Login).Methods(http.MethodGet)
	apiSubrouter.HandleFunc("/auth/refresh", h.AuthHandler.Refresh).Methods(http.MethodGet)
	apiSubrouter.HandleFunc("/auth/welcome", h.AuthHandler.Welcome).Methods(http.MethodGet)

	apiSubrouter.HandleFunc("/book", auth.JWTAuth(h.PostBook)).Methods(http.MethodPost)
	apiSubrouter.HandleFunc("/book", h.GetAllBooks).Methods(http.MethodGet)
	apiSubrouter.HandleFunc("/book/{id}", h.GetBook).Methods(http.MethodGet)
	apiSubrouter.HandleFunc("/book/{id}", auth.JWTAuth(h.UpdateBook)).Methods(http.MethodPut)
	apiSubrouter.HandleFunc("/book/{id}", auth.JWTAuth(h.DeleteBook)).Methods(http.MethodDelete)

	apiSubrouter.HandleFunc("/audiobooks", h.GetAllAudiobooks).Methods(http.MethodGet)
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
