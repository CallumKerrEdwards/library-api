package http

import (
	"net/http"

	"github.com/CallumKerrEdwards/neterrific"
	"github.com/gorilla/mux"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		neterrific.SendHTTPJSONError(w, http.StatusBadRequest, errRequiredID)
		return
	}

	var newBook books.Book

	err := neterrific.ParseAndValidate(r, &newBook)
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusBadRequest, err)
		return
	}

	newlyUpdatedBook, err := h.Service.UpdateBook(r.Context(), id, &newBook)
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusInternalServerError, err)
		return
	}

	neterrific.SendJSON(w, http.StatusOK, newlyUpdatedBook)
}
