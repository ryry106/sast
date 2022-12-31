package model

import (
	"sort"
	"time"
)

type TasksList struct {
	list []Tasks
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

func NewTasksList(list []Tasks) *TasksList {
	return &TasksList{list: list}
}

func (tl *TasksList) Sort() *TasksList {
	var csvNameList []string
	for _, t := range tl.list {
		csvNameList = append(csvNameList, t.Name)
	}
	sort.Strings(csvNameList)
	return NewTasksList(sortTasksList(tl.list, csvNameList))
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

func (tl *TasksList) ToSPDailyLists(start time.Time, end time.Time) *SPDailyLists {
	var sl []SPDailyList
	for _, t := range tl.list {
		sl = append(sl, *t.toSPDailyList(start, end))
	}
	return NewSPDailyLists(sl)
}

func (tl *TasksList) ToSPDailyListsEntirePeriod(now time.Time) *SPDailyLists {
	var sl []SPDailyList
	for _, t := range tl.list {
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
