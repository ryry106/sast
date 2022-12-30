package model

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type TasksList struct {
	List []Tasks
}

type Tasks struct {
	Name string
	List []Task
}

type Task struct {
	SP       int
	CreateDt time.Time
	FixedDt  time.Time
}

func NewTasksListFromCSVDir(dirPath string) (*TasksList, error) {
	csvList, err := dirwalk(dirPath)
	if err != nil {
		return &TasksList{}, err
	}

	fileNum := len(csvList)
	rcvCh := make(chan resultParsedCSV, fileNum)
	defer close(rcvCh)
	for _, csv := range csvList {
		go func(rcvCh chan<- resultParsedCSV, csv string) {
			rcvCh <- parseCSV(csv)
		}(rcvCh, csv)
	}

	var tasksList []Tasks
	for i := 0; i < fileNum; i++ {
		cpr := <-rcvCh
		tasksList = append(tasksList, *cpr.t)
	}

	return &TasksList{sortTasksList(tasksList, csvList)}, nil
}

func (tl *TasksList) ToSPDailyLists(start time.Time, end time.Time) *SPDailyLists {
	var sl []SPDailyList
	for _, t := range tl.List {
		sl = append(sl, *t.toSPDailyList(start, end))
	}
	return NewSPDailyLists(sl)
}

func (tl *TasksList) ToSPDailyListsEntirePeriod(now time.Time) *SPDailyLists {
	var sl []SPDailyList
	for _, t := range tl.List {
		sl = append(sl, *t.toSPDailyListEntirePeriod(now))
	}
	return NewSPDailyLists(sl)
}

func (tl *Tasks) toSPDailyListEntirePeriod(end time.Time) *SPDailyList {
	return tl.toSPDailyList(tl.mostEarlyDt(end), end)
}

func (tl *Tasks) toSPDailyList(start time.Time, end time.Time) *SPDailyList {
	return tl.calculateSP(*NewSPDailyList(tl.Name, start, end))
}

func (tl *Tasks) calculateSP(spdl SPDailyList) *SPDailyList {
	for _, task := range tl.List {
		for i, spd := range spdl.List {
			if isAddSP(spd.Dt, task) {
				spdl.List[i].SP += task.SP
			}
		}
	}
	return &spdl
}

func (tl *Tasks) mostEarlyDt(now time.Time) time.Time {
	var mostEarlyDt = now
	for _, task := range tl.List {
		if mostEarlyDt.After(task.CreateDt) {
			mostEarlyDt = task.CreateDt
		}
	}
	return mostEarlyDt
}

func isAddSP(addTargetDt time.Time, task Task) bool {
	// タスク生成日がSP加算日付よりも後に生成された
	if task.CreateDt.After(addTargetDt) {
		return false
	}
	// タスク完了日がSP加算対象日よりも前に設定されている
	if !task.FixedDt.IsZero() && !task.FixedDt.After(addTargetDt) {
		return false
	}
	return true
}

type resultParsedCSV struct {
	t *Tasks
	e []errorParsedCSV
}

type errorParsedCSV struct {
	csv string
	row string
	err error
}

func parseCSV(path string) resultParsedCSV {
	fp, err := os.Open(path)
	if err != nil {
		return resultParsedCSV{e: []errorParsedCSV{{csv: path, err: err}}}
	}
	defer fp.Close()

	list := []Task{}
	errors := []errorParsedCSV{}
	s := bufio.NewScanner(fp)
	for s.Scan() {
		row := s.Text()
		sp, cd, fd, err := parseCSVRow(row)
		if err != nil {
			errors = append(errors, errorParsedCSV{csv: path, row: row, err: err})
			continue
		}
		list = append(list, Task{SP: sp, CreateDt: cd, FixedDt: fd})
	}

	return resultParsedCSV{t: &Tasks{Name: path, List: list}, e: errors}
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

func sortTasksList(tasksList []Tasks, csvList []string) []Tasks {
	var res []Tasks
	for _, csv := range csvList {
		for _, tasks := range tasksList {
			if csv == tasks.Name {
				res = append(res, tasks)
				break
			}
		}
	}
	return res
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
