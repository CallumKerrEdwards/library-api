package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var (
	unauthorizedMessage = "not authorized"
	errUnauthorized     = errors.New(unauthorizedMessage)
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := getAuthHeader(r)
		if err != nil {
			http.Error(w, unauthorizedMessage, http.StatusUnauthorized)
			return
		}

		if validateToken(jwtToken) {
			original(w, r)
		} else {
			http.Error(w, unauthorizedMessage, http.StatusUnauthorized)
			return
		}
	}
}

func getAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header["Authorization"]
	if authHeader == nil {
		return "", errUnauthorized
	}

	authHeaderParts := strings.Split(authHeader[0], " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errUnauthorized
	}
	return authHeaderParts[1], nil
}

func validateToken(accessToken string) bool {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return getJWTSigningKey(), nil
	})

	if err != nil {
		return false
	}
	return token.Valid
}
