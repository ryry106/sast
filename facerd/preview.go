package facerd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"ryry/prev-bdc/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(csvPath string) {
	// todo file存在チェック

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := &Resource{csvPath: csvPath}

	// Routes
	e.GET("/preview", preview)
	e.GET("/resource", r.resource)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func preview(c echo.Context) error {
	f, err := os.Open("facerd/template.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	template, err := io.ReadAll(f)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	c.Response().Header().Set(echo.HeaderContentType, "text/html; charset=UTF-8")
	return c.String(http.StatusOK, string(template))
}

type Resource struct {
	csvPath string
}

// Handler
func (r *Resource) resource(c echo.Context) error {
	tl, err := model.ToTaskList(r.csvPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	fmt.Println(tl)
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	return c.String(http.StatusOK, tl.ToSPDaily(time.Now().In(tz)).ToJson())
}
