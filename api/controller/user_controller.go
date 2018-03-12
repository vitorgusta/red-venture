package controller

import (
	"net/http"

	"github.com/vitorgusta/red-venture/api/repository"

	"github.com/vitorgusta/red-venture/api/models"

	"github.com/labstack/echo"
)

type (
	user models.User
)

func CreateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	result := repository.CreateUser(*u)
	return c.JSON(http.StatusCreated, result)
}

func FindAll(c echo.Context) error {
	result := repository.FindAllUsers()
	return c.JSON(http.StatusOK, result)
}

func FindById(c echo.Context) error {
	id := c.Param("id")
	result := repository.FindUserById(id)
	return c.JSON(http.StatusOK, result)
}
