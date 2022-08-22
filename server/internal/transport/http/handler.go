package http

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/CallumKerrEdwards/library/server/pkg/books"
	"github.com/CallumKerrEdwards/library/server/pkg/log"
)

var (
	address = "0.0.0.0:8080"
)

type BookService interface {
	PostBook(context.Context, books.Book) (books.Book, error)
	GetBook(ctx context.Context, ID string) (books.Book, error)
}

type Handler struct {
	Router  *mux.Router
	Service BookService
	Server  *http.Server
	Log     log.Logger
}

func NewHandler(service BookService, logger log.Logger) *Handler {
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

	h.Server = &http.Server{
		Addr:    address,
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	h.Router.HandleFunc("/api/v1/book", h.PostBook).Methods("POST")
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
	h.Server.Shutdown(ctx)

	h.Log.Infoln("Shut down server gracefully")
	return nil
}
