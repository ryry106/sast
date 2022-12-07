package model

import (
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
