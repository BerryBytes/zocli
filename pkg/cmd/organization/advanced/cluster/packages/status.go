package packages

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

// NewPackageStatusCommand
// @Description: return status command
// @param o
// @return packageInstallCmd
// Description: return status command
func NewPackageStatusCommand(o *Opts) *cobra.Command {
	packageInstallCmd := &cobra.Command{
		Use:                   "status",
		Aliases:               []string{"s", "st", "sta", "stat", "statu", "check"},
		Short:                 "get status for packages",
		Long:                  grammar.PackagesStatusHelp,
		Run:                   o.StatusRunner,
		DisableFlagsInUseLine: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			middleware.LoggedIn(o.F)
			if len(args) != 0 {
				id, _ := validation.GetIDOrName(args[0])
				o.ClusterID, _ = strconv.Atoi(id)
			}
		},
	}

	packageInstallCmd.Flags().IntVarP(&o.ClusterID, "id", "i", 0, "id of the cluster")
	return packageInstallCmd
}

// StatusRunner
// @Description: get status for packages
// @param cmd
// @param args
// Description: root runner for getting status for packages
func (o *Opts) StatusRunner(*cobra.Command, []string) {
	if o.ClusterID == 0 {
		o.F.Printer.Fatal(3, "cluster id is required")
		return
	}
	o.I.StatusAll()
}

// StatusAll
// Description: get status for all packages
// This function is used to get status for all packages
func (o *Opts) StatusAll() {
	headers := o.F.GetAuth()
	reqConfig := defaults.StatusForPackage(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(o.ClusterID))
	res := reqConfig.Request()
	var response map[string]interface{}
	byteData, _ := json.Marshal(res.Data)
	err := json.Unmarshal(byteData, &response)
	if err != nil {
		o.F.Printer.Fatal(3, "cannot unmarshal")
		return
	}
	o.PackagesStatus = make(map[string]bool)
	for key, value := range response {
		o.PackagesStatus[key] = value.(bool)
	}

	o.seperateTrueAndFalse()
	o.PrintAll()
}

// seperateTrueAndFalse
//
// Description: separate true and false packages
// This function is used to separate true and false packages
func (o *Opts) seperateTrueAndFalse() {
	truePackages := make(map[string]bool, 0)
	falsePackages := make(map[string]bool, 0)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for key, value := range o.PackagesStatus {
		wg.Add(1)
		go func(key string, value bool) {
			defer wg.Done()
			if value {
				mutex.Lock()
				truePackages[key] = value
				mutex.Unlock()
			} else {
				mutex.Lock()
				falsePackages[key] = value
				mutex.Unlock()
			}
		}(key, value)
	}
	wg.Wait()
	o.TrueStatus = truePackages
	o.FalsStatus = falsePackages
}

// PrintAll
// Description: print all packages status in tabular format
func (o *Opts) PrintAll() {
	tableprinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tableprinter.HeaderRow(color, "packages", "status")

	for key, value := range o.TrueStatus {
		tableprinter.AddField(key)
		tableprinter.AddField(fmt.Sprint(value))
		tableprinter.EndRow()
	}
	for key, value := range o.FalsStatus {
		tableprinter.AddField(key)
		tableprinter.AddField(fmt.Sprint(value))
		tableprinter.EndRow()
	}
	_ = tableprinter.Print()
}
