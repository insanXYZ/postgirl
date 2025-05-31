package log

import (
	"fmt"
	"time"
)

var log = []string{}

func AddLog(v string) {
	log = append(log, fmt.Sprintf("[%v] %v \n", time.Now().Format(time.TimeOnly), v))
}

func GetLogs() []string {
	return log
}
