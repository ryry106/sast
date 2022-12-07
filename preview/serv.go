package preview

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Serv struct {
	CsvPath string
}

func (s *Serv) Up() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ph := &prevHandler{templateHtmlPath: "assets/template.html"}
	rh := &resourceHandler{csvPath: s.CsvPath}

	// Routes
	e.GET("/preview", ph.handle)
	e.GET("/resource", rh.handle)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
