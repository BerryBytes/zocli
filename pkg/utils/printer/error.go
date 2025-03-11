package printer

import (
	"fmt"
	"os"
)

func (p *printerStruct) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func (p *printerStruct) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (p *printerStruct) Fatal(exitCode int, args ...interface{}) {
	fmt.Println(args...)
	os.Exit(exitCode)
}
func (p *printerStruct) Fatalf(exitCode int, format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
	os.Exit(exitCode)
}
