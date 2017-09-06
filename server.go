package main

import (
	"github.com/datake914/lonbutu/controllers"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routing
	e.GET("/words", controllers.NewWordController().Get)

	e.Start(":8000")
}
