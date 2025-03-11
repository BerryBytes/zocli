package vars

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/formatter"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectVarsGetCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	getProjVars := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "get",
		Aliases:               []string{"ge", "g"},
		Short:                 "project scoped variables",
		GroupID:               "actions",
		Long:                  grammar.ProjectVarsGetHelp,
		Run:                   opts.getVars,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(opts.F)
		},
	}

	getProjVars.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	getProjVars.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	getProjVars.Flags().StringVarP(&opts.Output, "out", "o", "", "output in json or yaml")
	getProjVars.Flags().BoolVarP(&opts.ShowSecret, "show", "s", false, "shows plain text of secret")
	return getProjVars
}

func (o *Opts) getVars(cmd *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.printVars(o.getProjectDetail(o.ID).Variables)
		o.F.Printer.Exit(0)
		return
	}

	if o.Name != "" {
		o.printVars(o.getProjectDetailByName().Variables)
		o.F.Printer.Exit(0)
	}
	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.printVars(o.getProjectDetail(fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)).Variables)
		o.F.Printer.Exit(0)
		return
	}
	_ = cmd.Help()
}

func (o *Opts) printVars(vars []api.Variable) {
	if o.Output == "json" {
		if len(vars) == 1 {
			formatter.PrintJson(o.F, vars[0])
			return
		}
		formatter.PrintJson(o.F, vars)
		return
	} else if o.Output == "yaml" {
		if len(vars) == 1 {
			formatter.PrintYaml(o.F, vars[0])
			return
		}
		formatter.PrintYaml(o.F, vars)
		return
	}

	tablePrinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tablePrinter.HeaderRow(color, "id", "key", "value", "type")

	for _, each := range vars {
		tablePrinter.AddField(fmt.Sprint(each.Id))
		tablePrinter.AddField(each.Key)

		value := each.Value
		if each.Type == "secret" && !o.ShowSecret {
			value = strings.Repeat("*", len(value))
		}
		tablePrinter.AddField(value)
		tablePrinter.AddField(each.Type)
		tablePrinter.EndRow()
	}

	tablePrinter.Print()
}
