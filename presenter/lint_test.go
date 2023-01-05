package presenter

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestOutPut(t *testing.T) {
	b := new(bytes.Buffer)
	OutPut("testcsv", []error{errors.New("err1"), errors.New("err2")}, b)

	actual := b.String()
	expects := []string{"testcsv", "err1", "err2"}
	for _, e := range expects {
		if strings.Index(actual, e) == -1 {
			t.Errorf("fail. display %s", actual)
		}
	}
}

func TestOutPutEmpty(t *testing.T) {
	b := new(bytes.Buffer)
	OutPut("testcsv", []error{}, b)

	actual := b.String()
	if actual != "" {
		t.Errorf("fail. display %s", actual)
	}
}
