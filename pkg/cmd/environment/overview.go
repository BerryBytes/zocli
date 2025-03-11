package environment

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/formatter"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewEnvironmentOverviewCommand(o *Opts) *cobra.Command {
	env := &cobra.Command{
		Use:     "overview",
		Aliases: []string{"o", "over", "ove", "overv"},
		Short:   "overview of environment",
		Long:    grammar.EnvironmentOverviewHelp,
		GroupID: "basic",
		Run:     o.I.OverviewRunner,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				o.EnvID, o.EnvName = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	env.Flags().StringVarP(&o.EnvID, "id", "i", "", "environment id")
	env.Flags().StringVarP(&o.EnvName, "name", "n", "", "environment name")

	env.Flags().StringVarP(&o.Output, "out", "o", "", "output in json or yaml")
	return env
}

func (o *Opts) OverviewRunner(cmd *cobra.Command, _ []string) {
	if o.EnvID != "" && o.EnvName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}
	if o.EnvID != "" {
		o.I.GetEnvironmentOverview(o.EnvID)
		o.F.Printer.Exit(0)
		return
	}
	if o.EnvName != "" {
		o.F.Printer.Println("Please use id for now, as this is still under development")
		o.F.Printer.Exit(0)
		return
	}
	_ = cmd.Help()
}

func (o *Opts) GetEnvironmentOverview(id string) {
	headers := o.F.GetAuth()
	lastminute := time.Now().Add(time.Minute * -1).Unix()
	now := time.Now().Unix()

	body := []byte(`{"start_time": ` + fmt.Sprint(lastminute) + `,"end_time": ` + fmt.Sprint(now) + `}`)
	reqConfig := defaults.EnvironmentGetOverview(o.F, map[string]interface{}{"headers": headers, "body": body})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", id)
	res := reqConfig.Request()
	var overview api.EnvironmentOverview
	err := overview.FromJSON(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal data")
	}
	o.Overview = &overview

	o.I.PrintOverviewTable()
}

// PrintOverviewTable
//
// prints the tabular view for the overview of the environment.
// NOTE: this only prints: id, cpu, ram and disk usage
func (o *Opts) PrintOverviewTable() {
	if o.Output != "" {
		switch strings.ToLower(o.Output) {
		case "json":
			formatter.PrintJson(o.F, o.Overview)
			o.F.Printer.Exit(0)
			return
		case "yaml":
			formatter.PrintYaml(o.F, o.Overview)
			o.F.Printer.Exit(0)
			return
		}
	}
	tableprinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tableprinter.HeaderRow(color, "envid", "name", "cpu", "ram", "disk")

	printCPUPercentage := "NaN"
	printMemoryUsage := "NaN"
	printDiskUsages := "NaN"

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if len(o.Overview.CPUUsages) != 0 {
			cpu := o.Overview.CPUUsages[len(o.Overview.CPUUsages)-1]
			percentage, err := strconv.ParseFloat(cpu.Values[0][1].(string), 64)
			if err == nil {
				printCPUPercentage = fmt.Sprintf("%.2f", percentage*1000) + " %"
			}
		}
	}()

	go func() {
		defer wg.Done()
		if len(o.Overview.MemoryUsages) != 0 {
			mem := o.Overview.MemoryUsages[len(o.Overview.MemoryUsages)-1]
			percentage, err := strconv.Atoi(mem.Values[0][1].(string))
			if err == nil {
				floatValue := float64(percentage) / 1000000.0
				printMemoryUsage = fmt.Sprintf("%.2f", floatValue) + " MB"
			}
		}
	}()

	if len(o.Overview.DiskUsages) != 0 {
		disk := o.Overview.DiskUsages[len(o.Overview.DiskUsages)-1]
		percentage, err := strconv.Atoi(disk.Values[0][1].(string))
		if err == nil {
			floatValue := float64(percentage) / 1000000.0
			printDiskUsages = fmt.Sprintf("%.2f", floatValue) + " MB"
		}
	}
	wg.Wait()

	tableprinter.AddField(o.EnvID)
	tableprinter.AddField(o.EnvName)
	tableprinter.AddField(printCPUPercentage)
	tableprinter.AddField(printMemoryUsage)
	tableprinter.AddField(printDiskUsages)
	tableprinter.Print()
}
