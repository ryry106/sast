package model

import (
	"encoding/json"
	"time"
)

type SPDailyLists struct {
	Lists []SPDailyList
}
type SPDailyList struct {
	Name string    `json:"name"`
	List []SPDaily `json:"list"`
}
type SPDaily struct {
	Dt time.Time `json:"dt"`
	SP int       `json:"sp"`
}

func NewSPDailyList(name string, start time.Time, end time.Time) *SPDailyList {
	diffDays := diffDays(start, end)
	spdailyList := make([]SPDaily, diffDays, diffDays)
	for i := 0; i < diffDays; i++ {
		spdailyList[i].Dt = start.AddDate(0, 0, i)
	}
	return &SPDailyList{Name: name, List: spdailyList}
}

func (sl *SPDailyList) ToJson() string {
	j, _ := json.Marshal(sl)
	return "[" + string(j) + "]"
}

func (sls *SPDailyLists) ToJson() string {
	j, _ := json.Marshal(sls.Lists)
	return string(j)
}

func diffDays(start time.Time, end time.Time) int {
	return int(end.Sub(start).Hours()/24) + 1
}
