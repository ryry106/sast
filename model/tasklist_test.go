package model

import (
	"reflect"
	"testing"
	"time"
)

var timeJst, _ = time.LoadLocation("Asia/Tokyo")

func TestCalculateSP(t *testing.T) {
	tl := TaskList{
		[]Task{
			{CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 1},
		},
	}
	now := time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst)
	expects := SPDailyList{
		[]SPDaily{
			{Dt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 0},
			{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 2},
		},
	}
	actual := tl.ToSPDaily(now)
	if !reflect.DeepEqual(expects, *actual) {
		t.Errorf("fail. expects: %v,actual: %v", expects, actual)
	}
}

func TestMostEarlyDt(t *testing.T) {
	tt := []struct {
		name     string
		tasklist TaskList
		now      time.Time
		expects  time.Time
	}{
		{
			name: "second is early date",
			tasklist: TaskList{
				[]Task{
					{CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
					{CreateDt: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst), SP: 1},
					{CreateDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
				},
			},
			now:     time.Date(2022, 11, 28, 0, 0, 0, 0, timeJst),
			expects: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst),
		},
		{
			name: "return now",
			tasklist: TaskList{
				[]Task{},
			},
			now:     time.Date(2022, 11, 28, 0, 0, 0, 0, timeJst),
			expects: time.Date(2022, 11, 28, 0, 0, 0, 0, timeJst),
		},
	}

	for _, test := range tt {
		actual := test.tasklist.mostEarlyDt(test.now)
		if !actual.Equal(test.expects) {
			t.Errorf("%s is fail. expects: %v,actual: %v", test.name, test.expects, actual)
		}
	}
}
