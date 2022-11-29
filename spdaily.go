package main

import "time"

type SPDailyList struct {
	List []SPDaily
}
type SPDaily struct {
	Dt time.Time
	SP int
}
