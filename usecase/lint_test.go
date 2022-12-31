package usecase

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	b := new(bytes.Buffer)
	NewLint(filepath.Dir("tests/lint")).Do(b)

	actual := b.String()
	expects := []string{"tests/lint/2.csv", ",2022-12-05,2022-12-06 | strconv.Atoi: parsing \"\": invalid syntax", "tests/lint/1.csv", "1,,2022-12-03 | parsing time \"\" as \"2006-01-02\": cannot parse \"\" as \"2006\""}
	for _, e := range expects {
		if strings.Index(actual, e) == -1 {
			t.Errorf("fail. display %s", actual)
		}
	}
}
