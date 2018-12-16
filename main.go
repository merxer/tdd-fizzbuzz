package main


import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	fz "github.com/merxer/fizzbuzz/fizzbuzz"
)

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
		num, _  := strconv.Atoi(n)
	        fizzbuzz := fz.FizzBuzz(num)
	        return c.String(http.StatusOK, "" + fizzbuzz )
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
