package model

import (
	"reflect"
	"testing"
	"time"
)

var timeJst, _ = time.LoadLocation("Asia/Tokyo")

func TestToSPDaily(t *testing.T) {
	tl := Tasks{
		Name: "TestToSPDaily",
		List: []Task{
			{CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 1},
		},
	}
	now := time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst)
	expects := SPDailyList{
		Name: "TestToSPDaily",
		List: []SPDaily{
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
		tasklist Tasks
		now      time.Time
		expects  time.Time
	}{
		{
			name: "second is early date",
			tasklist: Tasks{
				Name: "TestMostEarlyDt",
				List: []Task{
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
			tasklist: Tasks{
				Name: "TestMostEarlyDt",
				List: []Task{},
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

func TestTasksFromCSV(t *testing.T) {
	timeJst, _ = time.LoadLocation("Asia/Tokyo")
	expects := &Tasks{
		Name: "tests/source_csv.csv",
		List: []Task{
			{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)},
			{SP: 2, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst)},
			{SP: 3, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 8, 0, 0, 0, 0, timeJst)},
		},
	}
	actual, _ := TasksFromCSV("tests/source_csv.csv")
	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects, actual)
	}
}

func TestTasksListFromCSVDir(t *testing.T) {
	expects := &TasksList{
		List: []Tasks{
			{
				Name: "tests/tgtdir/source1.csv",
				List: []Task{
					{SP: 3, CreateDt: time.Date(2022, 11, 3, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 13, 0, 0, 0, 0, timeJst)},
					{SP: 1, CreateDt: time.Date(2022, 11, 3, 0, 0, 0, 0, timeJst)},
					{SP: 2, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 11, 0, 0, 0, 0, timeJst)},
				},
			},
			{
				Name: "tests/tgtdir/source2.csv",
				List: []Task{
					{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)},
					{SP: 2, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst)},
					{SP: 3, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 8, 0, 0, 0, 0, timeJst)},
				},
			},
			{
				Name: "tests/tgtdir/tmp/source3.csv",
				List: []Task{
					{SP: 5, CreateDt: time.Date(2022, 12, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 12, 10, 0, 0, 0, 0, timeJst)},
					{SP: 1, CreateDt: time.Date(2022, 12, 1, 0, 0, 0, 0, timeJst)},
					{SP: 3, CreateDt: time.Date(2022, 12, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 12, 8, 0, 0, 0, 0, timeJst)},
				},
			},
		},
	}

	actual, _ := TasksListFromCSVDir("tests/tgtdir")
	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects, actual)
	}
}
