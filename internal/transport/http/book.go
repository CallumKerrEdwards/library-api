package http

import (
	"encoding/json"
	"net/http"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

type Response struct {
	Message string
}

type PostBookRequest struct {
	Title       string           `json:"title" validate:"required"`
	Authors     []books.Person   `json:"authors" validate:"required"`
	Description string           `json:"description"`
	ReleaseDate *time.Time       `json:"releaseDate"`
	Genres      []books.Genre    `json:"genres"`
	Series      books.Series     `json:"series"`
	Arefacts    []books.Artefact `json:"artefacts"`
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

	convertedBook := h.convertPostBookRequestToBook(&request)

	postedBook, err := h.Service.PostBook(r.Context(), &convertedBook)
	if err != nil {
		h.Log.WithError(err).Errorln("Could not post book")
		return
	}

	if err := json.NewEncoder(w).Encode(postedBook); err != nil {
		h.Log.WithError(err).Errorln("Error handling request", r)
	}
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	fetched, err := h.Service.GetAllBooks(r.Context())
	if err != nil {
		h.Log.WithError(err).Errorln("Cannot get all book")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(fetched); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.WithError(err).Errorln("Cannot encode response")

		return
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
		h.Log.WithError(err).Errorln("Cannot get all books with id", id)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(fetched); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.WithError(err).Errorln("Cannot encode response")

		return
	}
}

func (h *Handler) convertPostBookRequestToBook(r *PostBookRequest) books.Book {
	return books.NewBook(
		r.Title, r.Description, r.Authors, r.ReleaseDate, r.Genres, r.Series, r.Arefacts)
}
