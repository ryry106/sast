package model

import (
	"testing"
	"time"
)

func TestDiffDays(t *testing.T) {
	expects := 5
	actual := diffDays(time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), time.Date(2022, 11, 27, 0, 0, 0, 0, timeJst))
	if expects != actual {
		t.Errorf("fail. %d, %d", expects, actual)
	}
}

func TestToJson(t *testing.T) {
	sl := &SPDailyList{
		List: []SPDaily{
			{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 2},
			{Dt: time.Date(2022, 11, 26, 0, 0, 0, 0, timeJst), SP: 3},
		},
	}
	expects := "{\"List\":[{\"Dt\":\"2022-11-23T00:00:00+09:00\",\"SP\":1},{\"Dt\":\"2022-11-25T00:00:00+09:00\",\"SP\":2},{\"Dt\":\"2022-11-26T00:00:00+09:00\",\"SP\":3}]}"
	actual := sl.ToJson()
	if expects != actual {
		t.Errorf("fail. %s, %s", expects, actual)
	}
}
