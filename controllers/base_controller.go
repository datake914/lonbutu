package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type baseController struct {
}

func (bc *baseController) Get(c echo.Context) error {
	return echo.NewHTTPError(http.StatusMethodNotAllowed)
}

func (bc *baseController) Post(c echo.Context) error {
	return echo.NewHTTPError(http.StatusMethodNotAllowed)
}

func (bc *baseController) Delete(c echo.Context) error {
	return echo.NewHTTPError(http.StatusMethodNotAllowed)
}

func (bc *baseController) Put(c echo.Context) error {
	return echo.NewHTTPError(http.StatusMethodNotAllowed)
}

func (bc *baseController) Head(c echo.Context) error {
	return echo.NewHTTPError(http.StatusMethodNotAllowed)
}

func (bc *baseController) Patch(c echo.Context) error {
	return echo.NewHTTPError(http.StatusMethodNotAllowed)
}

func (bc *baseController) Options(c echo.Context) error {
	return echo.NewHTTPError(http.StatusMethodNotAllowed)
}
