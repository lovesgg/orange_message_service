package eventdispatcher

import (
	"fmt"
	"log"
)

var Logger *log.Logger

func Logf(s string, args ...interface{}) {
	if Logger == nil {
		return
	}
	_ = Logger.Output(1, fmt.Sprintf(s, args...))
}

func Warnf(s string, args ...interface{}) {
	if Logger == nil {
		return
	}
	_ = Logger.Output(2, fmt.Sprintf(s, args...))
}
