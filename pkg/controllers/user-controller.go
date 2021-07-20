package controllers

import (
	"fiber-jwt/pkg/cookies"
	"fiber-jwt/pkg/models"
	"html/template"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignInForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		formPath := path.Join("templates", "signin.html")
		templates, err := template.ParseFiles(formPath)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if err := templates.Execute(c.Response().Writer, nil); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

// Sign in function goes here ...
func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Load the test user
		user := models.LoadSampleUser()

		// !init-method, Instantiate user model
		u := new(models.User)

		// Get the form inputs, and parse to our user (myUser)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Compare the passwords
		if err := bcrypt.CompareHashAndPassword([]byte(models.LoadSampleUser().Password), []byte(u.Password)); err != nil {
			//  Return 401, since passwords didn't match
			// Don't specify that only password is incorrect
			return echo.NewHTTPError(http.StatusUnauthorized, "Username or Password is incorrect")
		}

		// If password is okay, create token and cookie
		err := cookies.GenerateTokenCookies(user, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
		}
		return c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}
