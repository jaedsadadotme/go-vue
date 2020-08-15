package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Static("views"))
	e.File("*", "views/index.html")
	e.Logger.Fatal(e.Start(":9000"))
}
