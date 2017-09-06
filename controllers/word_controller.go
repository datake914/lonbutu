package controllers

import (
	"github.com/datake914/lonbutu/services"
	"github.com/labstack/echo"
)

type wordController struct {
	baseController
}

// NewWordController creates instance an instane of wordController
func NewWordController() Controller {
	return &wordController{}
}

func (wc *wordController) Get(c echo.Context) error {
	services.NewServiceExecuter(&services.WordGetService{}).Execute()
	return nil
}

func (wc *wordController) Post(c echo.Context) error {
	services.NewServiceExecuter(&services.WordGetService{}).Execute()
	return nil
}
