package main

import (
	"fmt"
	"time"
)

type TaskList struct {
	List []Task
}
type Task struct {
	Name     string
	SP       int
	CreateDt time.Time
	FixedDt  time.Time
}

func (tl *TaskList) SPDaily() {


  fmt.Println("test")

}
