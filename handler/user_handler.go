package handler

import (
	"github.com/gactocat/snowshoe/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userRepo models.UserRepo
}

func NewUserHandler() UserHandler {
	return UserHandler{
		userRepo: models.NewUserRepo(),
	}
}

func (handler UserHandler) CreateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	err := handler.userRepo.Insert(u)
	if err != nil {
		log.Fatalln("Insert failed", err)
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func (handler UserHandler) ReadUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.NewUserRepo().Find(id)
	if err != nil {
		log.Fatalln("Select failed", err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (handler UserHandler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Id = id

	err := handler.userRepo.Update(u)
	if err != nil {
		log.Fatalln("Update failed", err)
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := handler.userRepo.Delete(id)
	if err != nil {
		log.Fatalln("Delete failed", err)
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
