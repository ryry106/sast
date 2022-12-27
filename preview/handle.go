package preview

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
	"sast/model"
	"time"

	"github.com/labstack/echo"
)

//go:embed assets/*
var assets embed.FS

type prevHandler struct {
	templateHtmlPath string
	port             string
}

func (ph *prevHandler) handle(c echo.Context) error {
	blob, err := ph.concretePreview()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.HTMLBlob(http.StatusOK, blob)
}

func (ph *prevHandler) concretePreview() ([]byte, error) {
	tb, err := assets.ReadFile(ph.templateHtmlPath)
	if err != nil {
		return nil, err
	}
	tpl, err := template.New("preview").Parse(string(tb))
	if err != nil {
		return nil, err
	}

	v := struct {
		Port string
	}{
		Port: ph.port,
	}

	var res bytes.Buffer
	if err := tpl.Execute(&res, v); err != nil {
		return nil, err
	}
	return res.Bytes(), nil
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
