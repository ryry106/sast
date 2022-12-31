package usecase

import (
	"fmt"
	"io"
	"sast/infra"
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
		csv, errors := r.Errors()
		if len(errors) == 0 {
			continue
		}
		fmt.Fprintln(out, csv)
		for _, e := range errors {
			fmt.Fprintln(out, "    "+e.String())
		}
	}
}
