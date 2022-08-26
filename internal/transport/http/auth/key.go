package auth

import "os"

var (
	jwtSigningKey []byte
)

func getJWTSigningKey() []byte {
	if len(jwtSigningKey) < 1 {
		jwtSigningKey = []byte(os.Getenv("JWT_SIGNING_KEY"))
	}
	return jwtSigningKey
}
