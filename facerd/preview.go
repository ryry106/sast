package facerd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Run() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/preview", preview)
	e.GET("/resource", resource)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func preview(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func resource(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
