package main

import (
	"fiber-jwt/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	// Define route group
	adminGroup := app.Group("/admin")
	userGroup := app.Group("/user")
	// Get admin controller
	adminGroup.GET("", controllers.AdminController())
	// Signin user
	userGroup.GET("/signin", controllers.SignInForm()).Name = "userSignInForm"
	userGroup.POST("/signin", controllers.SignIn())
	// Start the server
	app.Logger.Fatal(app.Start(":3000"))
}
