package printer

import (
	"fmt"
)

type DebugInterface interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type PrinterInterface interface {
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fatal(exitCode int, args ...interface{})
	Fatalf(exitCode int, format string, args ...interface{})
	Exit(exitCode int)
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

type printerStruct struct{}

type debugStruct struct{}

var debug bool

func New() PrinterInterface {
	return &printerStruct{}
}
func NewDebug(isDebugOn bool) DebugInterface {
	debug = isDebugOn
	return &debugStruct{}
}

func (ds *debugStruct) Debugf(format string, args ...interface{}) {
	if !debug {
		return
	}
	fmt.Printf(format+"\n", args...)
}

func (ds *debugStruct) Debug(args ...interface{}) {
	if !debug {
		return
	}
	fmt.Println(args...)
}
