package infra

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"sast/model"
	"strconv"
	"strings"
	"time"
)

type resultParsedCSVDir struct {
	list []resultParsedCSV
}

type resultParsedCSV struct {
	t *model.Tasks
	e []parseError
}

type parseError struct {
	csv string
	row string
	err error
}

func (rp *resultParsedCSVDir) List() []resultParsedCSV {
	return rp.list
}

func (pe *parseError) String() string {
	return pe.row + " | " + pe.err.Error()
}

func (rp *resultParsedCSVDir) ToTasksList() *model.TasksList {
	var list []model.Tasks
	for _, r := range rp.list {
		list = append(list, *r.t)

	}
	return model.NewTasksList(list)
}

func (rp *resultParsedCSV) Errors() (string, []parseError) {
	return rp.t.Name, rp.e
}

func ParseFromCSVDir(dirPath string) (*resultParsedCSVDir, error) {
	csvList, err := dirwalk(dirPath)
	if err != nil {
		return nil, err
	}

	fileNum := len(csvList)
	rcvCh := make(chan resultParsedCSV, fileNum)
	defer close(rcvCh)
	for _, csv := range csvList {
		go func(rcvCh chan<- resultParsedCSV, csv string) {
			rcvCh <- parseCSV(csv)
		}(rcvCh, csv)
	}

	var results []resultParsedCSV
	for i := 0; i < fileNum; i++ {
		cpr := <-rcvCh
		results = append(results, cpr)
	}

	return &resultParsedCSVDir{results}, nil

}

func parseCSV(path string) resultParsedCSV {
	fp, err := os.Open(path)
	if err != nil {
		return resultParsedCSV{e: []parseError{{csv: path, err: err}}}
	}
	defer fp.Close()

	list := []model.Task{}
	errors := []parseError{}
	s := bufio.NewScanner(fp)
	for s.Scan() {
		row := s.Text()
		sp, cd, fd, err := parseCSVRow(row)
		if err != nil {
			errors = append(errors, parseError{csv: path, row: row, err: err})
			continue
		}
		list = append(list, model.Task{SP: sp, CreateDt: cd, FixedDt: fd})
	}

	return resultParsedCSV{t: &model.Tasks{Name: path, List: list}, e: errors}
}

func parseCSVRow(line string) (int, time.Time, time.Time, error) {
	la := strings.Split(line, ",")
	sp, err := convertSP(la[0])
	if err != nil {
		return 0, time.Time{}, time.Time{}, err
	}
	createDt, err := convertCreateDt(la[1])
	if err != nil {
		return 0, time.Time{}, time.Time{}, err
	}
	fixedDt, err := convertFixedDt(la[2])
	if err != nil {
		return 0, time.Time{}, time.Time{}, err
	}
	return sp, createDt, fixedDt, nil
}

func convertSP(lineUnit string) (int, error) {
	sp, err := strconv.Atoi(lineUnit)
	if err != nil {
		return 0, err
	}
	return sp, nil
}

func convertCreateDt(str string) (time.Time, error) {
	return convertDt(str)
}

func convertFixedDt(str string) (time.Time, error) {
	if str == "" {
		return time.Time{}, nil
	}
	return convertDt(str)
}

func convertDt(lineUnit string) (time.Time, error) {
	timeJst, _ := time.LoadLocation("Asia/Tokyo")
	dt, err := time.ParseInLocation("2006-01-02", lineUnit, timeJst)
	if err != nil {
		return time.Time{}, err
	}
	return dt, nil
}

func dirwalk(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			res, _ := dirwalk(filepath.Join(dirPath, file.Name()))

			paths = append(paths, res...)
			continue
		}
		paths = append(paths, filepath.Join(dirPath, file.Name()))
	}

	return paths, nil
}
