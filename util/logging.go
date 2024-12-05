package util

import "fmt"

var MyLogEnabled = false

func MyLog(format string, args ...interface{}) {
	if MyLogEnabled {
		fmt.Printf(format, args...)
	}
}
