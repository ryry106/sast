package preview

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Serv struct {
	CsvDir       string
	Port         int
	TemplateName string
}

func (s *Serv) Up() {
	ph, rh, err := s.createHandler()
	if err != nil {
		fmt.Printf("failed to start preview server. %v", err)
		return
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/preview", ph.handle)
	e.GET("/resource", rh.handle)

	s.dispPreviewPath()

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(s.Port)))
}

func (s *Serv) createHandler() (*prevHandler, *resourceHandler, error) {
	ph, err := newPrevHandler(s.TemplateName, s.Port)
	if err != nil {
		return nil, nil, err
	}

	rh, err := newResourceHandler(s.CsvDir)
	if err != nil {
		return nil, nil, err
	}
	return ph, rh, nil
}

func (s *Serv) dispPreviewPath() {
	fmt.Println("preview to http://localhost:" + strconv.Itoa(s.Port) + "/preview")
}
