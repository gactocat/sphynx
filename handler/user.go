package handler

import (
	"github.com/gactocat/snowshoe/config"
	"github.com/gactocat/snowshoe/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

func CreateUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func ReadUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	//return c.JSON(http.StatusOK, users[id])
	log.Println(id)
	u1 := models.User{}
	conn := config.GetContext().Db.Connection()
	err := conn.SelectOne(&u1, "select id, name from user where id = ?", id)
	if err != nil {
		log.Fatalln("Select failed", err)
		return err
	}
	return c.JSON(http.StatusOK, u1)
}

func UpdateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
