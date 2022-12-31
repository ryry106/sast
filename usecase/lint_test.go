package usecase

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestDo(t *testing.T) {
	b := new(bytes.Buffer)
	NewLint(filepath.Dir("tests/lint")).Do(b)

	actual := b.String()
	expects := `tests/lint/2.csv
    ,2022-12-05,2022-12-06 | strconv.Atoi: parsing "": invalid syntax
tests/lint/1.csv
    1,,2022-12-03 | parsing time "" as "2006-01-02": cannot parse "" as "2006"
`
	if expects != actual {
		t.Errorf("fail. expects %s, actual %s", expects, actual)
	}
}
