package main

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/labstack/echo"
)

func main() {
	color.White("Is this white?")
	color.Black("Is this black?")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
