package http

import (
	"errors"
	"net/http"

	"github.com/CallumKerrEdwards/neterrific"
	"github.com/gorilla/mux"
)

var (
	errRequiredID = errors.New("id is required")
)

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	fetched, err := h.Service.GetAllBooks(r.Context())
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusInternalServerError, err)
		return
	}

	neterrific.SendJSON(w, http.StatusOK, neterrific.Payload{
		"books": fetched,
	})
}

func (h *Handler) GetAllAudiobooks(w http.ResponseWriter, r *http.Request) {
	fetched, err := h.Service.GetAllAudiobooks(r.Context())
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusInternalServerError, err)
		return
	}

	neterrific.SendJSON(w, http.StatusOK, neterrific.Payload{
		"audiobooks": fetched,
	})
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		neterrific.SendHTTPJSONError(w, http.StatusBadRequest, errRequiredID)
		return
	}

	fetched, err := h.Service.GetBook(r.Context(), id)
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusNotFound, err)
		return
	}

	neterrific.SendJSON(w, http.StatusOK, fetched)
}
