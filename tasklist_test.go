package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSPDaily(t *testing.T) {
	var timeJst, _ = time.LoadLocation("Asia/Tokyo")
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
				SP: 2,
			},
		},
	}

	actual := tasklist.SPDaily()
	if !reflect.DeepEqual(expects, actual) {
    t.Errorf("expects: %v,actual: %v", expects, actual)
	}
}
