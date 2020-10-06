package main

import (
	"github.com/ericdube/echoexample/customMiddleware"
	"net/http"
	"sync/atomic"

	_ "github.com/ericdube/echoexample/docs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Simple Echo Test
// @version 1.0
// @description This is a simple test of the Echo framework.

// @host localhost:1234
// @BasePath /

// User struct
type User struct {
	Name string `json::"name"`
	UserID int64 `json:"userID"`
}

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
	logTest := customMiddleware.NewLogger()

	// Middleware
	e.Use(logTest.Process)
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.GET("/hello", hello)
	e.POST("/user", user)

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

// user
// @Summary Add a new User
// @Description Adds a new user to the list of Users
// @ID add-user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Returns new user"
// @Router /user [post]
func user(c echo.Context) error {
	// Create a new user
	u := new(User)
	// Bind the request body to the new user
	if err := c.Bind(u); err != nil {
		return err
	}

	u.UserID = incUserCount()

	//TODO: Do something with new User

	// Return new user
	return c.JSON(http.StatusOK, u)
}