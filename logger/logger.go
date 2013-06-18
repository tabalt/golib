package logger

import (
	"fmt"
	"github.com/tabalt/golib/array"
	"github.com/tabalt/golib/file"
	"os"
)

var logTypeList = []string{"Notice", "Fatal", "Error", "Default"}

func saveLog(logType string, logContent string) {
	if !array.InArray(logType, logTypeList) {
		logType = "Default"
	}
	content := fmt.Sprintf("%s : %s\n", logType, logContent)
	file.WriteFile("log.txt", content)
}

func Notice(logContent string) {
	saveLog("Notice", logContent)
}

func Fatal(logContent string) {
	saveLog("Fatal", logContent)
	os.Exit(1)
}

func Error(logContent string) {
	saveLog("Error", logContent)
}

func Default(logContent string) {
	saveLog("Default", logContent)
}
