package main

import (
	"reflect"
	"testing"
	"time"
)

var timeJst, _ = time.LoadLocation("Asia/Tokyo")

func TestSPDaily(t *testing.T) {
	tt := []struct {
		name     string
		tasklist TaskList
		now      time.Time
		expects  *SPDailyList
	}{
		{
			name: "20 - 28",
			tasklist: TaskList{
				[]Task{
					{Name: "first", CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
					{Name: "second", CreateDt: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst), SP: 1},
					{Name: "third", CreateDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
				},
			},
			now: time.Date(2022, 11, 28, 0, 0, 0, 0, timeJst),
			expects: &SPDailyList{
				[]SPDaily{
					{Dt: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 26, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 27, 0, 0, 0, 0, timeJst)},
					{Dt: time.Date(2022, 11, 28, 0, 0, 0, 0, timeJst)},
				},
			},
		},
	}

	for _, test := range tt {
		actual := test.tasklist.SPDaily()
		if !reflect.DeepEqual(test.expects, actual) {
			t.Errorf("%s is fail. expects: %v,actual: %v", test.name, test.expects, actual)
		}
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
