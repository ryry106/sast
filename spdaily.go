package main

import "time"

type SPDailyList struct {
	List []SPDaily
}
type SPDaily struct {
	Dt time.Time
	SP int
}

func NewSPDailyList(start time.Time, end time.Time) *SPDailyList {
	diffDays := diffDays(start, end)
	spdailyList := make([]SPDaily, diffDays, diffDays)
	for i := 0; i < diffDays; i++ {
		spdailyList[i].Dt = start.AddDate(0, 0, i)
	}
	return &SPDailyList{spdailyList}
}

func diffDays(start time.Time, end time.Time) int {
	return int(end.Sub(start).Hours()/24) + 1
}
