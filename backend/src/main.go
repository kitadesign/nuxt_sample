package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	e.Debug = true
	// e.Use(middleware.Recover())

	e.POST("/", TopHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		e.Logger.Info("Defaulting to port %s", port)
	}
	_ = os.Setenv("LISTEN_PORT", port)

	e.Logger.Info("Listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

func TopHandler(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"code":    http.StatusOK,
			"message": http.StatusText(http.StatusOK),
		},
	)
}
