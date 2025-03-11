package printer

import (
	"fmt"
	"os"
)

func (p *printerStruct) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}

func (p *printerStruct) Print(v ...interface{}) {
	fmt.Print(v...)
}

func (p *printerStruct) Println(v ...interface{}) {
	fmt.Printf("\n%v\n", v...)
}

func (p *printerStruct) Exit(exitCode int) {
	os.Exit(exitCode)
}
