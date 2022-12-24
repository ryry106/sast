package preview

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Serv struct {
	CsvDir string
}

func (s *Serv) Up() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ph := &prevHandler{templateHtmlPath: "assets/template.html"}
	rh := &resourceHandler{csvDir: s.CsvDir}

	// Routes
	e.GET("/preview", ph.handle)
	e.GET("/resource", rh.handle)

	dispPreviewPath()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func dispPreviewPath() {
	fmt.Println("preview to http://localhost:8080/preview")
}
