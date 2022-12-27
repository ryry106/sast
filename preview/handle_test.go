package preview

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestPreviewHandler(t *testing.T) {
	ph := &prevHandler{templateHtmlPath: "assets/template.html"}
	c := echo.New().NewContext(httptest.NewRequest("GET", "/stub", strings.NewReader("")), httptest.NewRecorder())
	err := ph.handle(c)
	if err != nil || http.StatusOK != c.Response().Status {
		t.Error(err, c.Response())
	}
}

func TestPreviewHandlerErrorNoTemplate(t *testing.T) {
	ph := &prevHandler{templateHtmlPath: "dummy_template_path"}
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(httptest.NewRequest("GET", "/stub", strings.NewReader("")), rec)
	err := ph.handle(c)
	if err != nil || http.StatusInternalServerError != rec.Code {
		t.Error(rec.Code, rec.Body.String())
	}
}


func TestResourceHandler(t *testing.T) {
	ph := &resourceHandler{csvDir: "tests/sample"}
	c := echo.New().NewContext(httptest.NewRequest("GET", "/stub", strings.NewReader("")), httptest.NewRecorder())
	err := ph.handle(c)
	if err != nil || http.StatusOK != c.Response().Status {
		t.Error(err, c.Response())
	}
}

func TestResourceHandlerErrorNoTemplate(t *testing.T) {
	ph := &resourceHandler{csvDir: "dummypath"}
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(httptest.NewRequest("GET", "/stub", strings.NewReader("")), rec)
	err := ph.handle(c)
	if err != nil || http.StatusInternalServerError != rec.Code {
		t.Error(rec.Code, rec.Body.String())
	}
}
