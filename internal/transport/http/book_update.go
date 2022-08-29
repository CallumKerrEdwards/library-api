package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newBook books.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		h.Log.WithError(err).Errorln("Error updating book with ID", id)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	newlyUpdatedBook, err := h.Service.UpdateBook(r.Context(), id, &newBook)

	h.Log.WithField("updated book", true).Infoln(newlyUpdatedBook)

	if err != nil {
		h.Log.WithError(err).Errorln("Error updating book with ID", id)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(newlyUpdatedBook); err != nil {
		h.Log.WithError(err).Errorln("Cannot encode response")
	}
}
