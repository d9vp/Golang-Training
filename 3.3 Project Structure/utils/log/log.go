package log

import "fmt"

type Logger interface {
	Info(value ...interface{})
	Error(value ...interface{})
	Warning(value ...interface{})
}

type Log struct {
}

func GetLogger() Logger {
	return &Log{}
}

func (l *Log) Info(value ...interface{}) {
	fmt.Println("<<<<<INFO<<<<<")
	fmt.Println(value...)
	fmt.Println("<<<<<INFO<<<<<")

}

func (l *Log) Error(value ...interface{}) {
	fmt.Println("<<<<<ERROR<<<<<")
	fmt.Println(value...)
	fmt.Println("<<<<<ERROR<<<<<")
}

func (l *Log) Warning(value ...interface{}) {
	fmt.Println("<<<<<WARNING<<<<<")
	fmt.Println(value...)
	fmt.Println("<<<<<WARNING<<<<<")
}
