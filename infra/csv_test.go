package infra

import (
	"reflect"
	"sast/model"
	"testing"
	"time"
)

var timeJst, _ = time.LoadLocation("Asia/Tokyo")

func TestToTasksList(t *testing.T) {
	rp := resultParsedCSVDir{
		list: []resultParsedCSV{
			{
				t: model.NewTasks("tasks1", []*model.Task{{SP: 3, CreateDt: time.Date(2022, 11, 3, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 13, 0, 0, 0, 0, timeJst)}}),
			},
			{
				t: model.NewTasks("tasks2", []*model.Task{{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)}}),
			},
			{
				t: model.NewTasks("tasks3", []*model.Task{{SP: 5, CreateDt: time.Date(2022, 12, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 12, 10, 0, 0, 0, 0, timeJst)}}),
			},
		},
	}
	expects := model.NewTasksList([]model.Tasks{
		*model.NewTasks("tasks1", []*model.Task{{SP: 3, CreateDt: time.Date(2022, 11, 3, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 13, 0, 0, 0, 0, timeJst)}}),
		*model.NewTasks("tasks2", []*model.Task{{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)}}),
		*model.NewTasks("tasks3", []*model.Task{{SP: 5, CreateDt: time.Date(2022, 12, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 12, 10, 0, 0, 0, 0, timeJst)}}),
	})

	actual := rp.ToTasksList()
	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects, actual)
	}
}

func TestParseCSV(t *testing.T) {
	timeJst, _ = time.LoadLocation("Asia/Tokyo")
	expects := resultParsedCSV{
		t: model.NewTasks(
			"tests/source_csv.csv",
			[]*model.Task{
				{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)},
				{SP: 2, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst)},
				{SP: 3, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 8, 0, 0, 0, 0, timeJst)},
			},
		),
	}
	actual := parseCSV("tests/source_csv.csv")
	if !reflect.DeepEqual(actual.t, expects.t) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects.t, actual.t)
	}
}

func TestParseFromCSVDir(t *testing.T) {
	actual, _ := ParseFromCSVDir("tests/tgtdir")
	expects := resultParsedCSVDir{
		list: []resultParsedCSV{
			{
				t: model.NewTasks(
					"tests/tgtdir/source1.csv",
					[]*model.Task{
						{SP: 3, CreateDt: time.Date(2022, 11, 3, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 13, 0, 0, 0, 0, timeJst)},
						{SP: 1, CreateDt: time.Date(2022, 11, 3, 0, 0, 0, 0, timeJst)},
						{SP: 2, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 11, 0, 0, 0, 0, timeJst)},
					},
				),
			},
			{
				t: model.NewTasks(
					"tests/tgtdir/source2.csv",
					[]*model.Task{
						{SP: 1, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 10, 0, 0, 0, 0, timeJst)},
						{SP: 2, CreateDt: time.Date(2022, 11, 1, 0, 0, 0, 0, timeJst)},
						{SP: 3, CreateDt: time.Date(2022, 11, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 8, 0, 0, 0, 0, timeJst)},
					},
				),
			},
			{
				t: model.NewTasks(
					"tests/tgtdir/tmp/source3.csv",
					[]*model.Task{
						{SP: 5, CreateDt: time.Date(2022, 12, 1, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 12, 10, 0, 0, 0, 0, timeJst)},
						{SP: 1, CreateDt: time.Date(2022, 12, 1, 0, 0, 0, 0, timeJst)},
						{SP: 3, CreateDt: time.Date(2022, 12, 4, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 12, 8, 0, 0, 0, 0, timeJst)},
					},
				),
			},
		},
	}

	if len(expects.list) != len(actual.list) {
		t.Errorf("fail. \nexpects: %v\nactual:  %v", expects, actual)
	}

	for _, e := range expects.list {
		ok := false
		for _, a := range actual.list {
			if reflect.DeepEqual(e.t, a.t) {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("fail. \nexpects: %v\nactual:  %v", expects, actual)
			break
		}
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
