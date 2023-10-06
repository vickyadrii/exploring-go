package main

import "github.com/labstack/echo/v4"

type Users struct {
	Name string
	Email string
}

func main () {
	e := echo.New() // Create instance
	e.GET("/users", getUsersController)
	e.GET("/users/:id", getDetailUsersController)
	// e.Start(":5000")
	e.Logger.Fatal(e.Start(":8002"))

}

func getUsersController(c echo.Context) error {
	// Response Code
	// Response data
	var users = []Users{{"Alta", "altagmail.com"}}
	return c.JSON(200, users)
}

func getDetailUsersController(c echo.Context) error {
	var user = Users{"Alta", "altagmail.com"}
	return c.JSON(200, user)
}