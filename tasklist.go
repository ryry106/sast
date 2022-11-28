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

type SPDailyList struct {
	List []SPDaily
}
type SPDaily struct {
	Dt time.Time
	SP int
}

// todo nowを引数
func (tl *TaskList) SPDaily() *SPDailyList {
	var timeJst, _ = time.LoadLocation("Asia/Tokyo")
	now := time.Date(2022, 11, 28, 0, 0, 0, 0, timeJst)
	mostEarlyDt := tl.mostEarlyDt(now)
	diffDays := int(now.Sub(mostEarlyDt).Hours() / 24) + 1

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
