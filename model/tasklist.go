package model

import (
	"sort"
	"time"
)

type TasksList struct {
	list []Tasks
}

type Tasks struct {
	name string
	list []*Task
}

type Task struct {
	sp       int
	createDt time.Time
	fixedDt  time.Time
}

func NewTasks(name string, list []*Task) *Tasks {
	return &Tasks{name: name, list: list}
}

func NewTask(sp int, createDt time.Time, fixedDt time.Time) *Task {
	return &Task{sp: sp, createDt: createDt, fixedDt: fixedDt}
}

func NewTasksList(list []Tasks) *TasksList {
	return &TasksList{list: list}
}

func (t *Tasks) Name() string {
	return t.name
}

func (tl *TasksList) Sort() *TasksList {
	var csvNameList []string
	for _, t := range tl.list {
		csvNameList = append(csvNameList, t.name)
	}
	sort.Strings(csvNameList)
	return NewTasksList(sortTasksList(tl.list, csvNameList))
}

func sortTasksList(tasksList []Tasks, csvList []string) []Tasks {
	var res []Tasks
	for _, csv := range csvList {
		for _, tasks := range tasksList {
			if csv == tasks.name {
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
	return tl.calculateSP(*NewSPDailyList(tl.name, start, end))
}

func (tl *Tasks) calculateSP(spdl SPDailyList) *SPDailyList {
	for _, task := range tl.list {
		for i, spd := range spdl.List {
			if isAddSP(spd.Dt, task) {
				spdl.List[i].SP += task.sp
			}
		}
	}
	return &spdl
}

func (tl *Tasks) mostEarlyDt(now time.Time) time.Time {
	var mostEarlyDt = now
	for _, task := range tl.list {
		if mostEarlyDt.After(task.createDt) {
			mostEarlyDt = task.createDt
		}
	}
	return mostEarlyDt
}

func isAddSP(addTargetDt time.Time, task *Task) bool {
	// ?????????????????????SP??????????????????????????????????????????
	if task.createDt.After(addTargetDt) {
		return false
	}
	// ?????????????????????SP???????????????????????????????????????????????????
	if !task.fixedDt.IsZero() && !task.fixedDt.After(addTargetDt) {
		return false
	}
	return true
}
