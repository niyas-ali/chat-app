package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/get-message", func(c echo.Context) error {
		return c.String(http.StatusOK, "Getting message")
	})
	e.POST("/send-message", func(c echo.Context) error {
		param := c.FormValue("message")
		return c.String(http.StatusOK, "Sended message : "+param)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
