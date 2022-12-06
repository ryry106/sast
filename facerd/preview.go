package facerd

import (
	"fmt"
	"net/http"
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
	return c.String(http.StatusOK, "Hello, World!")
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
