package packages

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

// NewPackageInstallCommand
//
// @param o
// @return packageInstallCmd
// Description: initializes a new command for installing packages on cluster and returns it
func NewPackageInstallCommand(o *Opts) *cobra.Command {
	packageInstallCmd := &cobra.Command{
		Use:                   "install",
		Aliases:               []string{"i", "in", "ins", "inst", "insta", "install", "add", "a"},
		Short:                 "install packages on cluster",
		Long:                  grammar.PackagesInstallHelp,
		Run:                   o.InstallRunner,
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
	packageInstallCmd.Flags().StringVarP(&o.PackageName, "name", "n", "", "name of the package")
	packageInstallCmd.Flags().StringVarP(&o.PackageGroupToInstall, "template", "t", "", "standard template to install")
	packageInstallCmd.Flags().BoolVarP(&o.ShowTemplates, "show-templates", "s", false, "show the standard templates and exit")
	return packageInstallCmd
}

// InstallRunner
//
// @param cmd
// @param args
// @return err
//
// Description: root runner for installing packages on cluster
func (o *Opts) InstallRunner(*cobra.Command, []string) {
	// check if the user wants to see the templates
	// if yes then call the function to fetch the configs and print the templates and exit
	if o.ShowTemplates {
		o.FetchPackageConfigs()
		o.PrintPackageTemplates()
		return
	}
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

	// check if the package name or group is provided
	if o.PackageName == "" && o.PackageGroupToInstall == "" {
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Fatal(3, " package name or group is required")
		return
	}

	if o.PackageName != "" {
		// get the status for all the packages
		o.StatusAll()
		configWaitGroup.Wait()
		o.InstallPackage()
	} else {
		configWaitGroup.Wait()
		o.checkPackageGroup()
		o.InstallPackageGroup()
	}
	// now as the package is installed, we need to check if the package is already installed
	configWaitGroup.Wait()
}

// InstallPackageGroup
//
// @Description: Install the package by template on the cluster
func (o *Opts) InstallPackageGroup() {
	packageInfoMap := make(map[string]api.Packages)
	for _, packageInfo := range o.PackageConfig.Packages {
		packageInfoMap[packageInfo.Name] = packageInfo
	}

	var packageNames []string
	for _, packages := range o.PackageConfig.Templates {
		if packages.Name == o.PackageGroupToInstall {
			packageNames = packages.Packages
			break
		}
	}

	var body []map[string]interface{}
	for _, packageName := range packageNames {
		if packageInfo, ok := packageInfoMap[packageName]; ok {
			body = append(body, map[string]interface{}{
				"name":         packageInfo.Name,
				"chart":        packageInfo.Chart,
				"required_dns": packageInfo.RequiredDNS,
				"icon":         packageInfo.Icon,
				"needs":        packageInfo.Needs,
				"set":          []string{},
			})
		}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		o.F.Printer.Fatalf(3, "Error marshaling JSON")
		return
	}
	o.genericCaller(jsonData)
	o.F.Printer.Printf(o.F.IO.ColorScheme().SuccessIcon() + " Sent command to install package group '" + o.PackageGroupToInstall + "'.")
}

// genericCaller
//
// @param jsonData
// @Description: generic caller for install package and install package group
func (o *Opts) genericCaller(jsonData []byte) {
	reqConfig := defaults.InstallPackage(o.F, map[string]interface{}{"headers": o.F.GetAuth(), "body": jsonData})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", strconv.Itoa(o.ClusterID))
	_ = reqConfig.Request()
}

// InstallPackage
//
// @Description: Install the package on the cluster
func (o *Opts) InstallPackage() {
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
			o.genericCaller(jsonData)
			o.F.Printer.Printf(o.F.IO.ColorScheme().SuccessIcon() + " Sent command to install package '" + o.PackageName + "'.")
			return
		}
	}
	o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
	o.F.Printer.Fatal(3, " package '"+o.PackageName+"' not found")
}

// checkPackageGroup
//
// Description: check if the package group is valid
func (o *Opts) checkPackageGroup() {
	for _, v := range o.PackageConfig.Templates {
		if v.Name == o.PackageGroupToInstall {
			o.PackageGroupToInstall = v.Name
			return
		}
	}
	o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
	o.F.Printer.Fatal(3, " package group '"+o.PackageGroupToInstall+"' not found")
}

// FetchPackageConfigs
//
// @Description: fetch the package configs
// Description: fetch the package configs
func (o *Opts) FetchPackageConfigs() {
	headers := o.F.GetAuth()
	reqConfig := defaults.PackageConfig(o.F, map[string]interface{}{"headers": headers})
	res := reqConfig.Request()
	err := utils.ConvertType(res.Data, &o.PackageConfig)
	if err != nil {
		o.F.Printer.Fatalf(9, "cannot unmarchal data")
		return
	}
}

// PrintPackageTemplates
//
// @Description: print the package templates
func (o *Opts) PrintPackageTemplates() {
	tableprinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tableprinter.HeaderRow(color, "name", "packages")
	for _, v := range o.PackageConfig.Templates {
		tableprinter.AddField(v.Name)
		all := ""
		for i, packageInfo := range v.Packages {
			if i == len(v.Packages)-1 {
				all += packageInfo
			} else {
				all += packageInfo + ", "
			}
		}
		tableprinter.AddField(all)
		tableprinter.EndRow()
	}
	tableprinter.Print()
}
