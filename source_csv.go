package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func ToTaskList(path string) (*TaskList, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	list := []Task{}
	s := bufio.NewScanner(fp)
	for s.Scan() {
		sp, cd, fd := lineParse(s.Text())
		list = append(list, Task{SP: sp, CreateDt: cd, FixedDt: fd})
	}

	return &TaskList{List: list}, nil
}

func lineParse(line string) (int, time.Time, time.Time) {
	la := strings.Split(line, ",")
	sp := convertSP(la[0])
	createDt := convertDt(la[1])
	fixedDt := convertDt(la[2])
	return sp, createDt, fixedDt
}

func convertDt(lineUnit string) time.Time {
	timeJst, _ := time.LoadLocation("Asia/Tokyo")
	dt, err := time.ParseInLocation("2006-01-02", lineUnit, timeJst)
	if err != nil {
		return time.Time{}
	}
	return dt
}

func convertSP(lineUnit string) int {
	sp, err := strconv.Atoi(lineUnit)
	if err != nil {
		return 0
	}
	return sp
}
