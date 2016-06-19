package main

import (
	"database/sql"

	"github.com/gactocat/snowshoe/Godeps/_workspace/src/github.com/labstack/echo/engine/standard"
	"github.com/gactocat/snowshoe/handler"
	"github.com/gactocat/snowshoe/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
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

	sqlDb, _ := sql.Open("mysql", "sarasota-usr:sarasota-pwd@tcp(dev-dotmoney-dbm01.amb-stg-295.a4c.jp:3306)/sarasota")

	database := models.NewDatabase(sqlDb, models.Config{})

	u1 := models.User{Name: "John"}
	err := database.Connection().Insert(&u1)
	if err != nil {
		log.Fatalln(err)
	}

	// Start server
	e.Run(standard.New(":1323"))
}
