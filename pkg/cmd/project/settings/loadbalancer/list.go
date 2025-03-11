package loadbalancer

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/formatter"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectLoadbalancerListCommand(opts *Opts) *cobra.Command {
	listPermProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "list",
		Aliases:               []string{"l", "li", "lis", "get", "g", "ge"},
		Short:                 "list project scoped loadbalancers",
		Long:                  grammar.ProjectLoadbalancerListHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.I.ListLoadbalancer,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
		},
	}
	listPermProj.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	return listPermProj
}

func (o *Opts) ListLoadbalancer(cmd *cobra.Command, _ []string) {
	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.I.GetLoadbalancer()
		o.F.Printer.Exit(0)
		return
	}
	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.ID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
		o.I.GetLoadbalancer()
		o.F.Printer.Exit(0)
		return
	}

	_ = cmd.Help()
}

func (o *Opts) GetLoadbalancer() {
	headers := o.F.GetAuth()
	reqConfig := defaults.ListProjectLoadbalancer(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)
	res := reqConfig.Request()
	err := utils.ConvertType(res.Data, &o.Loadbalancers.LoadBalancers)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot marshal", err)
	}
	o.I.PrintLoadbalancer()
}

func (o *Opts) PrintLoadbalancer() {
	if o.Output == "json" {
		if len(o.Loadbalancers.LoadBalancers) == 1 {
			formatter.PrintJson(o.F, o.Loadbalancers.LoadBalancers[0])
			return
		}
		formatter.PrintJson(o.F, o.Loadbalancers.LoadBalancers)
		return
	} else if o.Output == "yaml" {
		if len(o.Loadbalancers.LoadBalancers) == 1 {
			formatter.PrintYaml(o.F, o.Loadbalancers.LoadBalancers[0])
			return
		}
		formatter.PrintYaml(o.F, o.Loadbalancers.LoadBalancers)
		return
	}

	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tablePrinter := table.New(o.F, 0)
	tablePrinter.HeaderRow(color, "id", "name", "project_id", "cluster_id", "custom_domain")

	for _, lb := range o.Loadbalancers.LoadBalancers {
		tablePrinter.AddField(fmt.Sprint(lb.Id))
		tablePrinter.AddField(lb.Name)
		tablePrinter.AddField(fmt.Sprint(lb.ProjectID))
		tablePrinter.AddField(fmt.Sprint(lb.ClusterID))
		tablePrinter.AddField(fmt.Sprint(lb.CustomDomain))
		tablePrinter.EndRow()
	}
	_ = tablePrinter.Print()
}

func (o *Opts) GetByIdLoadbalancer() {
	headers := o.F.GetAuth()
	reqConfig := defaults.GetByIdLoadbalancer(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)
	res := reqConfig.Request()
	err := utils.ConvertType(res.Data, &o.Loadbalancer)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot marshal", err)
	}
	o.I.PrintLoadbalancer()
}
