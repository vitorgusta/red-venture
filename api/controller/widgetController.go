package controller

import (
	"net/http"

	model "github.com/vitorgusta/red-venture/api/models"
	Repository "github.com/vitorgusta/red-venture/api/repository"

	"github.com/labstack/echo"
)

var (
	widget model.Widget
)

func CreateWidget(c echo.Context) error {
	w := new(model.Widget)

	if err := c.Bind(w); err != nil {
		return err
	}

	result := Repository.CreateWidget(*w)

	return c.JSON(http.StatusCreated, result)
}

func FindAllWidget(c echo.Context) error {
	result := Repository.FindAllWidget()

	return c.JSON(http.StatusOK, result)
}

func FindWidgetById(c echo.Context) error {
	id := c.Param("id")

	result := Repository.FindWidgetById(id)

	return c.JSON(http.StatusOK, result)
}

func UpdateWidget(c echo.Context) error {
	w := new(model.Widget)
	id := c.Param("id")
	if err := c.Bind(w); err != nil {
		return err
	}

	result := Repository.UpdateWidget(*w, id)
	return c.JSON(http.StatusCreated, result)
}
