package main

import (
	"github.com/gactocat/snowshoe/Godeps/_workspace/src/github.com/labstack/echo/engine/standard"
	"github.com/gactocat/snowshoe/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Handler
	userHandler := handler.NewUserHandler()

	// Routers
	e.POST("users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.ReadUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	// Start server
	e.Run(standard.New(":1323"))
}
