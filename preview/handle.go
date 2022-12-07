package preview

import (
	"io"
	"net/http"
	"os"
	"burndownchart/model"
	"time"

	"github.com/labstack/echo"
)

type prevHandler struct {
	templateHtmlPath string
}

func (ph *prevHandler) handle(c echo.Context) error {
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

type resourceHandler struct {
	csvPath string
}

func (r *resourceHandler) handle(c echo.Context) error {
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
