package shared

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	NodeName string
}

var logger *Logger = nil

func GetLogger() *Logger {
	if logger == nil {
		logger = &Logger{NodeName: "NEW SERVER"}
	}
	return logger
}

func (l *Logger) SetNodeName(nodeName string) {
	l.NodeName = nodeName
}

func (l Logger) Log(message string) {
	fmt.Println(message)

	f, err := os.OpenFile("storage/logs", os.O_APPEND|os.O_WRONLY, 0644)

	defer f.Close()

	if err != nil {
		fmt.Println("failed to open log file")
	}

	t := time.Now()
	timestamp := t.Format(time.RFC3339)

	_, err = f.WriteString(fmt.Sprintf("[%v] [%v] %v\n", timestamp, l.NodeName, message))

	if err != nil {
		fmt.Println("failed append to log file")
	}
}
