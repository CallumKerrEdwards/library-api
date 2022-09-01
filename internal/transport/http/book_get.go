package http

import (
	"encoding/json"
	"net/http"

	"github.com/CallumKerrEdwards/loggerrific"
	"github.com/gorilla/mux"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	fetched, err := h.Service.GetAllBooks(r.Context())
	if err != nil {
		h.Log.WithError(err).Errorln("Cannot get all books")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	respondWithBooksSlice("books", fetched, w, h.Log)
}

func (h *Handler) GetAllAudiobooks(w http.ResponseWriter, r *http.Request) {
	fetched, err := h.Service.GetAllAudiobooks(r.Context())
	if err != nil {
		h.Log.WithError(err).Errorln("Cannot get all audiobooks")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	respondWithBooksSlice("audiobooks", fetched, w, h.Log)
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
		h.Log.WithError(err).Errorln("Cannot get all book with id", id)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(w).Encode(fetched); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.WithError(err).Errorln("Cannot encode response")

		return
	}
}

func respondWithBooksSlice(key string, fetched []books.Book,
	w http.ResponseWriter, logger loggerrific.Logger) {
	booksResponse := make(map[string][]books.Book)

	booksResponse[key] = fetched

	if err := json.NewEncoder(w).Encode(booksResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.WithError(err).Errorln("Cannot encode response")

		return
	}
}
