package http

import (
	"net/http"

	"github.com/CallumKerrEdwards/neterrific"

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

	err := neterrific.ParseAndValidate(r, &request)
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusBadRequest, err)
		return
	}

	convertedBook := books.NewBook(
		request.Title, request.Description, request.Authors, request.ReleaseDate,
		request.Genres, request.Series, request.Audiobook)

	postedBook, err := h.Service.PostBook(r.Context(), &convertedBook)
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusInternalServerError, err)
		return
	}

	neterrific.SendJSON(w, http.StatusCreated, postedBook)
}
