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

func TestSPDaylyListToJson(t *testing.T) {
	sl := &SPDailyList{
		Name: "TestToJson",
		List: []SPDaily{
			{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1},
			{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 2},
			{Dt: time.Date(2022, 11, 26, 0, 0, 0, 0, timeJst), SP: 3},
		},
	}
	expects := "[{\"name\":\"TestToJson\",\"list\":[{\"dt\":\"2022-11-23T00:00:00+09:00\",\"sp\":1},{\"dt\":\"2022-11-25T00:00:00+09:00\",\"sp\":2},{\"dt\":\"2022-11-26T00:00:00+09:00\",\"sp\":3}]}]"
	actual := sl.ToJson()
	if expects != actual {
		t.Errorf("fail. %s, %s", expects, actual)
	}
}

func TestToJson(t *testing.T) {
	tests := []struct {
		name    string
		sls     *SPDailyLists
		expects string
	}{
		{
			name: "データあり",
			sls: &SPDailyLists{
				Lists: []SPDailyList{
					{Name: "TJ1", List: []SPDaily{{Dt: time.Date(2022, 11, 23, 0, 0, 0, 0, timeJst), SP: 1}, {Dt: time.Date(2022, 11, 26, 0, 0, 0, 0, timeJst), SP: 3}}},
					{Name: "TJ2", List: []SPDaily{{Dt: time.Date(2022, 11, 25, 0, 0, 0, 0, timeJst), SP: 2}}},
				},
			},
			expects: "[{\"name\":\"TJ1\",\"list\":[{\"dt\":\"2022-11-23T00:00:00+09:00\",\"sp\":1},{\"dt\":\"2022-11-26T00:00:00+09:00\",\"sp\":3}]},{\"name\":\"TJ2\",\"list\":[{\"dt\":\"2022-11-25T00:00:00+09:00\",\"sp\":2}]}]",
		},
		{
			name:    "空配列",
			sls:     &SPDailyLists{Lists: []SPDailyList{}},
			expects: "[]",
		},
		{
			name:    "初期値未設定",
			sls:     &SPDailyLists{},
			expects: "[]",
		},
	}

	for _, tt := range tests {
		actual := tt.sls.ToJson()
		if tt.expects != actual {
			t.Errorf("%s is fail. %s, %s", tt.name, tt.expects, actual)
		}
	}

}
