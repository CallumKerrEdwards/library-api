package auth

import (
	"encoding/json"
	"net/http"

	"github.com/CallumKerrEdwards/library/server/pkg/log"
	"github.com/golang-jwt/jwt/v4"
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
		if err == jwt.ErrSignatureInvalid {
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
