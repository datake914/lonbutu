package controllers

import "github.com/labstack/echo"

// Controller interface.
type Controller interface {
	// GET
	Get(c echo.Context) error
	// POST
	Post(c echo.Context) error
	// DELETE
	Delete(c echo.Context) error
	// PUT
	Put(c echo.Context) error
	// HEAD
	Head(c echo.Context) error
	// PATCH
	Patch(c echo.Context) error
	// OPTIONS
	Options(c echo.Context) error
}
