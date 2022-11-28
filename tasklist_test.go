package main

import (
	"reflect"
	"testing"
	"time"
)

var timeJst, _ = time.LoadLocation("Asia/Tokyo")

func TestSPDaily(t *testing.T) {
	tasklist := &TaskList{
		[]Task{
			{
				Name:     "first",
				CreateDt: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst),
				SP:       1,
			},
			{
				Name:     "second",
				CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst),
				SP:       1,
			},
			{
				Name:     "third",
				CreateDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst),
				SP:       1,
			},
		},
	}

	expects := &SPDailyList{
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

	actual := tasklist.SPDaily()
	if !reflect.DeepEqual(expects, actual) {
		t.Errorf("expects: %v,actual: %v", expects, actual)
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
					{Name: "first", CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
					{Name: "second", CreateDt: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst), SP: 1},
					{Name: "third", CreateDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
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
