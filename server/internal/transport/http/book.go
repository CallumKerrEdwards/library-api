package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CallumKerrEdwards/library/server/pkg/books"
	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string
}

type PostBookRequest struct {
	Title        string `json:"title" validate:"required"`
	Author       string `json:"author" validate:"required"`
	books.Series `json:"series,omitempty"`
}

func (h *Handler) PostBook(w http.ResponseWriter, r *http.Request) {
	var request PostBookRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "not a valid Book", http.StatusBadRequest)
		h.Log.WithError(err).Errorln("Cannot unmarshall book")
		return
	}
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		http.Error(w, "not a valid Book", http.StatusBadRequest)
		h.Log.WithError(err).Errorln("Book failed validation")
		return
	}

	convertedBook := h.convertPostBookRequestToBook(request)
	postedBook, err := h.Service.PostBook(r.Context(), convertedBook)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(postedBook); err != nil {
		h.Log.WithError(err).Errorln("Error handling request", r)
	}

}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fetched, err := h.Service.GetBook(r.Context(), id)
	if err != nil {
		h.Log.WithError(err).Errorln("Cannot get book with id", id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(fetched); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.WithError(err).Errorln("Cannot encode response")
		return
	}
}

func (h *Handler) convertPostBookRequestToBook(r PostBookRequest) books.Book {
	return books.NewBook(r.Title, r.Author, r.Series)
}
