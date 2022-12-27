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
	actual := tl.toSPDailyList(now)
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

func TestParseCSV(t *testing.T) {
	timeJst, _ = time.LoadLocation("Asia/Tokyo")
	expects := resultParsedCSV{
		t: &Tasks{
			Name: "tests/source_csv.csv",
			List: []Task{
				{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)},
				{SP: 2, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst)},
				{SP: 3, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 8, 0, 0, 0, 0, timeJst)},
			},
		},
	}
	actual := parseCSV("tests/source_csv.csv")
	if !reflect.DeepEqual(actual.t, expects.t) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects.t, actual.t)
	}
}

func TestNewTasksListFromCSVDir(t *testing.T) {
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

	actual, _ := NewTasksListFromCSVDir("tests/tgtdir")
	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects, actual)
	}
}

func TestParseCSVRow(t *testing.T) {
	tests := []struct {
		name            string
		row             string
		expectsSP       int
		expectsCraeteDt time.Time
		expectsFixedDt  time.Time
	}{
		{
			name:            "test",
			row:             "1,2022-11-01,2022-11-10",
			expectsSP:       1,
			expectsCraeteDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst),
			expectsFixedDt:  time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst),
		},
		{
			name:            "no fixedDt",
			row:             "1,2022-11-01,",
			expectsSP:       1,
			expectsCraeteDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst),
			expectsFixedDt:  time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		actualSP, actualCreateDt, actualFixedDt, err := parseCSVRow(tt.row)
		if actualSP != tt.expectsSP || !actualCreateDt.Equal(tt.expectsCraeteDt) || !actualFixedDt.Equal(tt.expectsFixedDt) || err != nil {
			t.Errorf("%s is fail.", tt.name)
		}
	}

}

func TestParseCSVRowError(t *testing.T) {
	tests := []struct {
		name string
		row  string
	}{
		{name: "empty", row: ""},
		{name: "SP is not exists", row: ",2022-11-01,2022-11-10"},
		{name: "SP is not int", row: "a,2022-11-01,2022-11-10"},
		{name: "createDt is not exists", row: "1,,2022-11-10"},
		{name: "createDt is invalid format", row: "1,2022/11/01,2022-11-10"},
		{name: "fixedDt is invalid format", row: "1,2022-11-01,2022/11/10"},
	}

	for _, tt := range tests {
		_, _, _, err := parseCSVRow(tt.row)
		if err == nil {
			t.Errorf("%s is fail.", tt.name)
		}
	}

}
