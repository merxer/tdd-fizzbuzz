package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	fz "github.com/merxer/fizzbuzz/fizzbuzz"
)

type Result struct {
	Message string `json:"message"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/fizzbuzz/:number", func(c echo.Context) error {
		n := c.Param("number")
		num, err := strconv.Atoi(n)
		if err != nil {
			result := Result{
				Message: "Please input number",
			}
			return c.JSON(http.StatusBadRequest, result)
		}
		fizzbuzz := fz.FizzBuzz(num)
		result := Result{
			Message: fizzbuzz,
		}
		return c.JSON(http.StatusOK, result)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
