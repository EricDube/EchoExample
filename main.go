package main

import (
	"net/http"

	"github.com/ericdube/benefitfinder/customMiddleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	//Declare new custom logger
	logTest := customMiddleware.NewLogger()
	
	// Middleware
	e.Use(logTest.Process)
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}