package main

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
