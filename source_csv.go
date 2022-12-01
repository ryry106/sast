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
		line := strings.Split(s.Text(), ",")
		sp, _ := strconv.Atoi(line[0])
		createDt := convertDt(line[1])
		fixedDt := convertDt(line[2])
		list = append(list, Task{SP: sp, CreateDt: createDt, FixedDt: fixedDt})
	}

	return &TaskList{List: list}, nil
}

func convertDt(lineUnit string) time.Time {
	timeJst, _ := time.LoadLocation("Asia/Tokyo")
	dt, err := time.ParseInLocation("2006-01-02", lineUnit, timeJst)
	if err != nil {
		return time.Time{}
	}
	return dt
}
