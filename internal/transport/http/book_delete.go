package http

import (
	"net/http"

	"github.com/CallumKerrEdwards/neterrific"
	"github.com/gorilla/mux"
)

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookID := vars["id"]
	if bookID == "" {
		neterrific.SendHTTPJSONError(w, http.StatusUnprocessableEntity, errRequiredID)
		return
	}

	err := h.Service.DeleteBook(r.Context(), bookID)
	if err != nil {
		neterrific.SendHTTPJSONError(w, http.StatusInternalServerError, err)
		return
	}

	neterrific.SendJSON(w, http.StatusOK, neterrific.Payload{
		"message": "Successfully deleted book with ID " + bookID,
	})
}
