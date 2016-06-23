package api

import (
	"github.com/gactocat/snowshoe/config"
	"github.com/gactocat/snowshoe/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Start() {
	config.Init()

	echo := echo.New()

	// Middleware
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	// Handler
	userHandler := handler.NewUserHandler()

	// Routers
	echo.POST("users", userHandler.CreateUser)
	echo.GET("/users/:id", userHandler.ReadUser)
	echo.PUT("/users/:id", userHandler.UpdateUser)
	echo.DELETE("/users/:id", userHandler.DeleteUser)

	// Start server
	echo.Run(standard.New(":8080"))
}
