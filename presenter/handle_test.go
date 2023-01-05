package presenter

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestPreviewHandler(t *testing.T) {
	ph, err := newPrevHandler("plane", 8080)
	if err != nil {
		t.Error(err)
	}
	c := echo.New().NewContext(httptest.NewRequest("GET", "/stub", strings.NewReader("")), httptest.NewRecorder())
	err = ph.handle(c)
	if err != nil || http.StatusOK != c.Response().Status {
		t.Error(err, c.Response())
	}
}

func TestPreviewHandlerErrorNoTemplate(t *testing.T) {
	_, err := newPrevHandler("dummy", 8080)
	if err == nil {
		t.Error("not exists template. but create prevHandler.")
	}
}

func TestResourceHandler(t *testing.T) {
	ph, err := newResourceHandler("tests/sample", "2022-12-29")
	if err != nil {
		t.Error(err)
	}
	c := echo.New().NewContext(httptest.NewRequest("GET", "/stub", strings.NewReader("")), httptest.NewRecorder())
	err = ph.handle(c)
	if err != nil || http.StatusOK != c.Response().Status {
		t.Error(err, c.Response())
	}
}

func TestResourceHandlerErrorNoTemplate(t *testing.T) {
	_, err := newResourceHandler("dummypath", "2022-12-30")
	if err == nil {
		t.Error("not exists csv dir. but create resourceHandler.")
	}
}
