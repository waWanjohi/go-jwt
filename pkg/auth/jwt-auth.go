package auth

import (
	"fiber-jwt/env"

	"github.com/dgrijalva/jwt-go"
)

// Get a secret key and token name
const (
	keyString             = "Awesome"
	accessTokenCookieName = "access-token"
)

// TODO: Add environment variable later
var jwtSecretKey = env.SetKey()

func GetAccessToken() string {
	return accessTokenCookieName
}

func GetJwtSecretKey() string {
	return jwtSecretKey
}

// Add required JWT claims
// In our case, we only need a username
type Claims struct {
	Name string `json: "name"` // Remember json is case-sensitive
	jwt.StandardClaims
}
