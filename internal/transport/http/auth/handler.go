package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"

	"github.com/CallumKerrEdwards/library-api/pkg/log"
)

type Handler struct {
	Log log.Logger
}

type WelcomeResponse struct {
	User string `json:"user"`
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	jwtToken, err := getAuthHeader(r)
	if err != nil {
		http.Error(w, unauthorizedMessage, http.StatusUnauthorized)
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTSigningKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if err := json.NewEncoder(w).Encode(WelcomeResponse{User: claims.Username}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.WithError(err).Errorln("Cannot encode response")

		return
	}
}
