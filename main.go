/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"

	"github.com/berrybytes/zocli/cmd"
	"github.com/berrybytes/zocli/internal/config"
	userContext "github.com/berrybytes/zocli/pkg/utils/context"
	"github.com/berrybytes/zocli/pkg/utils/factory"
)

func main() {
	// will need this on near future to check for binary updates
	//info, _ := debug.ReadBuildInfo()
	//cobra.MousetrapHelpText = ""
	//fmt.Println(info)
	f := factory.New(context.Background(), config.New())
	userContext.Loader(f)
	err := cmd.Execute(f)
	if err != nil {
		f.Printer.Fatal(4, err)
	}
}
