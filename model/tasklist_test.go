package model

import (
	"reflect"
	"testing"
	"time"
)

var timeJst, _ = time.LoadLocation("Asia/Tokyo")

func TestSort(t *testing.T) {
	tl := &TasksList{
		list: []Tasks{
			{Name: "task5"},
			{Name: "task0"},
			{Name: "task3"},
			{Name: "task1"},
		},
	}
	expects := &TasksList{
		list: []Tasks{
			{Name: "task0"},
			{Name: "task1"},
			{Name: "task3"},
			{Name: "task5"},
		},
	}
	actual := tl.Sort()
	if !reflect.DeepEqual(expects, actual) {
		t.Errorf("fail. expects: %v,actual: %v", expects, actual)
	}

}

func TestToSPDailyList(t *testing.T) {
	tl := Tasks{
		Name: "ToSPDailyList",
		List: []Task{
			{CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 1},
		},
	}
	tests := []struct {
		name    string
		start   time.Time
		end     time.Time
		expects SPDailyList
	}{
		{
			name:  "タスクリスト外の開始日と内の終了日",
			start: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst),
			end:   time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst),
			expects: SPDailyList{
				Name: "ToSPDailyList",
				List: []SPDaily{
					{Dt: time.Date(2022, 11, 20, 0, 0, 0, 0, timeJst), SP: 0},
					{Dt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst), SP: 2},
				},
			},
		},
		{
			name:  "タスクリストと開始日、終了日が同じ",
			start: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst),
			end:   time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst),
			expects: SPDailyList{
				Name: "ToSPDailyList",
				List: []SPDaily{
					{Dt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst), SP: 2},
					{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 3},
				},
			},
		},
		{
			name:  "タスクリスト内の開始日と外の終了日",
			start: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst),
			end:   time.Date(2022, 11, 26, 0, 0, 0, 0, timeJst),
			expects: SPDailyList{
				Name: "ToSPDailyList",
				List: []SPDaily{
					{Dt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
					{Dt: time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst), SP: 2},
					{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 3},
					{Dt: time.Date(2022, 11, 26, 0, 0, 0, 0, timeJst), SP: 3},
				},
			},
		},
		{
			name:  "開始日と終了日が同じ",
			start: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst),
			end:   time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst),
			expects: SPDailyList{
				Name: "ToSPDailyList",
				List: []SPDaily{
					{Dt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		actual := tl.toSPDailyList(tt.start, tt.end)
		if !reflect.DeepEqual(tt.expects, *actual) {
			t.Errorf("%s fail. expects: %v,actual: %v", tt.name, tt.expects, actual)
		}
	}
}

func TestToSPDailyListEntirePeriod(t *testing.T) {
	tl := Tasks{
		Name: "ToSPDailyListEntirePeriod",
		List: []Task{
			{CreateDt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), FixedDt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{CreateDt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 1},
		},
	}
	now := time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst)
	expects := SPDailyList{
		Name: "ToSPDailyListEntirePeriod",
		List: []SPDaily{
			{Dt: time.Date(2022, 11, 21, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 22, 0, 0, 0, 0, timeJst), SP: 0},
			{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 24, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 2},
		},
	}
	actual := tl.toSPDailyListEntirePeriod(now)
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
