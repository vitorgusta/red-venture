package controller

import (
	"fmt"
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

	result, err := repository.CreateUser(*u)
	fmt.Println("vitor", err)
	return c.JSON(http.StatusCreated, result)
}

func FindAll(c echo.Context) error {
	result, err := repository.FindAllUsers()
	fmt.Println("testeseila", err)

	return c.JSON(http.StatusOK, result)
}

func FindById(c echo.Context) error {
	id := c.Param("id")

	result, err := repository.FindUserById(id)
	fmt.Println("teste", err)
	return c.JSON(http.StatusOK, result)
}
