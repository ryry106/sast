package presenter

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type serv struct {
	csvDir       string
	port         int
	templateName string
	startDate    string
}

func NewServ(csvDir string, port int, templateName string, startDate string) *serv {
	return &serv{csvDir: csvDir, port: port, templateName: templateName, startDate: startDate}
}

func (s *serv) Up() {
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
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(s.port)))
}

func (s *serv) createHandler() (*prevHandler, *resourceHandler, error) {
	ph, err := newPrevHandler(s.templateName, s.port)
	if err != nil {
		return nil, nil, err
	}

	rh, err := newResourceHandler(s.csvDir, s.startDate)
	if err != nil {
		return nil, nil, err
	}
	return ph, rh, nil
}

func (s *serv) dispPreviewPath() {
	fmt.Println("preview to http://localhost:" + strconv.Itoa(s.port) + "/preview")
}
