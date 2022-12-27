package preview

import (
	"embed"
	"net/http"
	"sast/model"
	"time"

	"github.com/labstack/echo"
)

//go:embed assets/*
var assets embed.FS

type prevHandler struct {
	templateHtmlPath string
}

func (ph *prevHandler) handle(c echo.Context) error {
	template, err := assets.ReadFile(ph.templateHtmlPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, "text/html; charset=UTF-8")

	return c.String(http.StatusOK, string(template))
}

type resourceHandler struct {
	csvDir string
}

func (r *resourceHandler) handle(c echo.Context) error {
	tl, err := model.NewTasksListFromCSVDir(r.csvDir)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	sls := *tl.ToSPDailyLists(time.Now().In(tz))
	return c.String(http.StatusOK, sls.ToJson())
}
