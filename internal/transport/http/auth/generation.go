package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	JWT string `json:"jwt"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: "anon",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSigningKey())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(LoginResponse{JWT: tokenString}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.WithError(err).Errorln("Cannot encode response")
		return
	}
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	jwtToken, err := getAuthHeader(r)
	if err != nil {
		http.Error(w, unauthorizedMessage, http.StatusUnauthorized)
		return
	}

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
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
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = &jwt.NumericDate{Time: expirationTime}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshedTokenString, err := token.SignedString(getJWTSigningKey())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(LoginResponse{JWT: refreshedTokenString}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Log.WithError(err).Errorln("Cannot encode response")
		return
	}
}
