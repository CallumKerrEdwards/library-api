package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookID := vars["id"]
	if bookID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteBook(r.Context(), bookID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted book with ID " + bookID}); err != nil {
		h.Log.WithError(err).Errorln("Cannot encode response")
	}
}
