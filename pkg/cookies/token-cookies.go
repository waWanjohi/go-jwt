package cookies

import (
	"fiber-jwt/pkg/auth"
	"fiber-jwt/pkg/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Generates a token
func generateToken(u *models.User, expiryTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &auth.Claims{
		Name: u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(), // set to expire and count in milliseconds
		},
	}

	// Create a token using the HS256 Algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create a jwt tokenstring
	tokenString, err := token.SignedString(auth.GetJwtSecretKey())
	if err != nil {
		return "", time.Now(), err
	}

	// Show the token and time
	return tokenString, expiryTime, nil
}

// Generates an access token that expires every ten minutes
func generateAccessToken(user *models.User) (string, time.Time, error) {
	expiryTime := time.Now().Add(10 * time.Minute)

	return generateToken(user, expiryTime, []byte(auth.GetJwtSecretKey()))
}

// set a token cookie
func setTokenCookie(name, token string, expiry time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiry
	cookie.Path = "/"

	// To prevent client-side script from accessing the server, I'll use http only
	// Ideally the client will be using https
	cookie.HttpOnly = true

	// Now add the cookie to echo context
	c.SetCookie(cookie)

}

// Same as the token cookie, only that this cookie will store the username
// And won't use http only
func setUserCookie(user *models.User, expiry time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = expiry
	cookie.Path = "/"

	c.SetCookie(cookie)
}

func GenerateTokenCookies(user *models.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(auth.GetAccessToken(), accessToken, exp, c)
	setUserCookie(user, exp, c)

	return nil
}
