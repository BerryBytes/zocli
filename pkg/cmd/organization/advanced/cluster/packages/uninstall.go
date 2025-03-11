package packages

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

// NewPackageUnInstallCommand
//
// @Description: return uninstall command
// @param o
// @return packageUnInstallCmd
// Description: return uninstall command
func NewPackageUnInstallCommand(o *Opts) *cobra.Command {
	packageUnInstallCmd := &cobra.Command{
		Use:     "uninstall",
		Aliases: []string{"u", "ui", "un", "unin", "unins", "uninst", "uninsta", "uninstal", "uin", "uins", "uinst", "uinsta", "uinstal", "remove", "r"},
		Short:   "uninstall packages on the cluster",
		Long:    grammar.PackagesUnInstallHelp,
		Run:     o.UnInstallRunner,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				id, _ := validation.GetIDOrName(args[0])
				o.ClusterID, _ = strconv.Atoi(id)
			}
		},
		DisableFlagsInUseLine: true,
	}

	packageUnInstallCmd.Flags().StringVarP(&o.PackageName, "name", "n", "", "name of the package to uninstall")
	packageUnInstallCmd.Flags().IntVarP(&o.ClusterID, "clusterid", "i", 0, "id of the cluster to uninstall package from")
	// packageUnInstallCmd.Flags().StringVarP(&o.)
	return packageUnInstallCmd
}

// UnInstallRunner
// @Description: uninstall packages on cluster
// @param cmd
// @param args
// Description: root runner for uninstalling packages on cluster
func (o *Opts) UnInstallRunner(*cobra.Command, []string) {
	var configWaitGroup sync.WaitGroup
	configWaitGroup.Add(1)
	go func() {
		defer configWaitGroup.Done()
		o.FetchPackageConfigs()
	}()

	// check if the cluster id is provided
	if o.ClusterID == 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Fatal(3, " cluster id is required")
		return
	}

	// check if the package name is provided
	if o.PackageName == "" {
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Fatal(3, " package name is required")
		return
	}

	configWaitGroup.Wait()
	for _, packages := range o.PackageConfig.Packages {
		if strings.EqualFold(packages.Title, o.PackageName) {
			var body []map[string]interface{}
			body = append(body, map[string]interface{}{
				"name":         packages.Name,
				"chart":        packages.Chart,
				"required_dns": packages.RequiredDNS,
				"icon":         packages.Icon,
				"needs":        packages.Needs,
				"set":          []string{},
			})
			jsonData, err := json.Marshal(body)
			if err != nil {
				o.F.Printer.Fatalf(3, "Error marshaling JSON")
				return
			}
			o.genericUninstaller(jsonData)
			o.F.Printer.Printf(o.F.IO.ColorScheme().SuccessIcon() + " Sent command to uninstall package '" + o.PackageName + "'.")
			return
		}
	}
	o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
	o.F.Printer.Fatal(3, " package '"+o.PackageName+"' not found")
}

func (o *Opts) genericUninstaller(jsonData []byte) {
	reqConfig := defaults.UninstallPackage(o.F, map[string]interface{}{"headers": o.F.GetAuth(), "body": jsonData})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(o.ClusterID))
	_ = reqConfig.Request()
}
