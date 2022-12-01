package main

import (
	"reflect"
	"testing"
	"time"
)

func TestToTaskList(t *testing.T) {
	timeJst, _ = time.LoadLocation("Asia/Tokyo")
	expects := &TaskList{
		List: []Task{
			{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)},
			{SP: 2, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst)},
			{SP: 3, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 8, 0, 0, 0, 0, timeJst)},
			{},
		},
	}
	actual, _ := ToTaskList("tests/source_csv.csv")
	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects, actual)
	}
}
