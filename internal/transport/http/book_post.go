package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
	"github.com/CallumKerrEdwards/library-api/pkg/books/genres"
)

type PostBookRequest struct {
	Title       string             `json:"title" validate:"required"`
	Authors     []books.Person     `json:"authors" validate:"required"`
	Description *books.Description `json:"description"`
	ReleaseDate *books.ReleaseDate `json:"releaseDate"`
	Genres      []genres.Genre     `json:"genres"`
	Series      books.Series       `json:"series"`
	Audiobook   *books.Audiobook   `json:"audiobook"`
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

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(postedBook); err != nil {
		h.Log.WithError(err).Errorln("Error handling request", r)
	}
}

func (h *Handler) convertPostBookRequestToBook(r *PostBookRequest) books.Book {
	return books.NewBook(
		r.Title, r.Description, r.Authors, r.ReleaseDate, r.Genres, r.Series, r.Audiobook)
}
