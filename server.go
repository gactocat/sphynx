package main

import (

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/gactocat/snowshoe/handler"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routers
	e.POST("users", handler.CreateUser)
	e.GET("/users/:id", handler.ReadUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	// Start server
	e.Run(standard.New(":1323"))
}
