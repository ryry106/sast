package main

import (
	"time"
)

type TaskList struct {
	List []Task
}
type Task struct {
	Name     string
	SP       int
	CreateDt time.Time
	FixedDt  time.Time
}

// todo 日付のリストを元にcreateDtを見てSPを加算するメソッドを切り出し
// todo 日付のリストを元にFixedDtを見てSPを減算するメソッドを切り出す
func (tl *TaskList) SPDaily(now time.Time) *SPDailyList {
	spdailyList := tl.toSPDailyListOnlyDt(now)

	return spdailyList
}

func (tl *TaskList) toSPDailyListOnlyDt(now time.Time) *SPDailyList {
	mostEarlyDt := tl.mostEarlyDt(now)
	diffDays := diffDays(mostEarlyDt, now)

	spdailyList := make([]SPDaily, diffDays, diffDays)
	for i := 0; i < diffDays; i++ {
		spdailyList[i].Dt = mostEarlyDt.AddDate(0, 0, i)
	}

	return &SPDailyList{spdailyList}
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

func diffDays(start time.Time, end time.Time) int {
	return int(end.Sub(start).Hours()/24) + 1
}
