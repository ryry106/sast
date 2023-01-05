package usecase

import (
	"errors"
	"io"
	"sast/infra"
	"sast/presenter"
)

type Lint struct {
	csvDir string
}

func NewLint(csvDir string) *Lint {
	return &Lint{csvDir: csvDir}
}

func (l *Lint) Do(out io.Writer) {
	res, err := infra.ParseFromCSVDir(l.csvDir)
	if err != nil {
		panic(err)
	}

	for _, r := range res.List() {
		csv, es := r.Errors()
		var errorList []error
		for _, e := range es {
			errorList = append(errorList, errors.New(e.String()))
		}
		presenter.OutPut(csv, errorList, out)
	}
}
