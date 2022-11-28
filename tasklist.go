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

func (tl *TaskList) SPDaily() *SPDailyList {
	var timeJst, _ = time.LoadLocation("Asia/Tokyo")
	return &SPDailyList{
		[]SPDaily{
			{
				Dt: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst),
				SP: 1,
			},
			{
				Dt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst),
				SP: 2,
			},
			{
				Dt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst),
				SP: 3,
			},
		},
	}

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
