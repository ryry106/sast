package model

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type TaskList struct {
	Name string
	List []Task
}
type Task struct {
	SP       int
	CreateDt time.Time
	FixedDt  time.Time
}

func (tl *TaskList) ToSPDaily(now time.Time) *SPDailyList {
	mostEarlyDt := tl.mostEarlyDt(now)
	spdailyList := NewSPDailyList(tl.Name, mostEarlyDt, now)
	return tl.calculateSP(*spdailyList)
}

func (tl *TaskList) calculateSP(spdl SPDailyList) *SPDailyList {
	for _, task := range tl.List {
		for i, spd := range spdl.List {
			if isAddSP(spd.Dt, task) {
				spdl.List[i].SP += task.SP
			}
		}
	}
	return &spdl
}

func (tl *TaskList) mostEarlyDt(now time.Time) time.Time {
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

func TaskListFromCSV(path string) (*TaskList, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	list := []Task{}
	s := bufio.NewScanner(fp)
	for s.Scan() {
		sp, cd, fd := lineParse(s.Text())
		if cd.IsZero() {
			continue
		}
		list = append(list, Task{SP: sp, CreateDt: cd, FixedDt: fd})
	}

	return &TaskList{Name: path, List: list}, nil
}

func lineParse(line string) (int, time.Time, time.Time) {
	la := strings.Split(line, ",")
	sp := convertSP(la[0])
	createDt := convertDt(la[1])
	fixedDt := convertDt(la[2])
	return sp, createDt, fixedDt
}

func convertDt(lineUnit string) time.Time {
	timeJst, _ := time.LoadLocation("Asia/Tokyo")
	dt, err := time.ParseInLocation("2006-01-02", lineUnit, timeJst)
	if err != nil {
		return time.Time{}
	}
	return dt
}

func convertSP(lineUnit string) int {
	sp, err := strconv.Atoi(lineUnit)
	if err != nil {
		return 0
	}
	return sp
}
