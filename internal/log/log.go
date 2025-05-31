package log

import (
	"fmt"
	"time"
)

var log = []string{}

func AddLog(v string) {
	log = append(log, fmt.Sprintf("[%v] %v", time.Now().Format(time.TimeOnly), v))
}

func GetLogs() []string {
	return log
}

func GetStringLogs() string {
	var logs string

	for _, v := range log {
		logs += v + "\n"
	}

	return logs
}
