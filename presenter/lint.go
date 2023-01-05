package presenter

import (
	"fmt"
	"io"
)

func OutPut(csvName string, errors []error, w io.Writer) {
	if len(errors) == 0 {
		return
	}
	fmt.Fprintln(w, csvName)
	for _, e := range errors {
		fmt.Fprintln(w, "    "+e.Error())
	}
}
