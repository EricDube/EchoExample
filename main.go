package main

import (
	"github.com/ericdube/echoexample/logger"
	"net/http"
	"sync/atomic"

	_ "github.com/ericdube/echoexample/docs"
	"github.com/ericdube/echoexample/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Simple Echo Test
// @version 1.0
// @description This is a simple test of the Echo framework.

// @host localhost:1234
// @BasePath /

// Declare a local userCount
var userCount int64

// increments the userCount using the atomic library so it can only be done one at a time
func incUserCount() int64 {
	return atomic.AddInt64(&userCount, 1)
}

// returns the current userCount
func getUserCount() int64 {
	return atomic.LoadInt64(&userCount)
}

func main() {

	// Echo instance
	e := echo.New()

	//Declare new custom logger
	if err := logger.NewLogger(); err != nil {
		return
	}

	// Middleware
	e.Use(logger.Logger.Process)
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.GET("/hello", hello)
	e.POST("/user", addUser)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// Hello
// @Summary Returns Hello, World!
// @Description Returns Hello, World!
// @ID hello
// @Success 200 {string} string	"ok"
// @Router /hello [get]
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// addUser
// @Summary Add a new User
// @Description Adds a new user to the list of Users
// @ID add-user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Returns new user"
// @Router /user [post]
func addUser(c echo.Context) error {
	// Create a new user
	u := new(models.User)
	// Bind the request body to the new user
	if err := c.Bind(u); err != nil {
		return err
	}

	u.UserID = incUserCount()

	//Save user to persistance store
	models.SaveUser(u)

	// Return new user
	return c.JSON(http.StatusOK, u)
}