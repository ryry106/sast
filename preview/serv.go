package preview

import (
	"io"
	"net/http"
	"os"
	"ryry/prev-bdc/model"
	"time"

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

	ph := &PrevHandler{templateHtmlPath: "assets/template.html"}
	rh := &ResourceHandler{csvPath: s.CsvPath}

	// Routes
	e.GET("/preview", ph.handle)
	e.GET("/resource", rh.handle)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}

type PrevHandler struct {
	templateHtmlPath string
}

func (ph *PrevHandler) handle(c echo.Context) error {
	f, err := os.Open(ph.templateHtmlPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	template, err := io.ReadAll(f)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set(echo.HeaderContentType, "text/html; charset=UTF-8")
	return c.String(http.StatusOK, string(template))
}

type ResourceHandler struct {
	csvPath string
}

func (r *ResourceHandler) handle(c echo.Context) error {
	tl, err := model.ToTaskList(r.csvPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, tl.ToSPDaily(time.Now().In(tz)).ToJson())
}
