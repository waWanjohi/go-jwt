package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminController() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi, You have access!")
	}
}
