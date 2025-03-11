package loadbalancer

import (
	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/spf13/cobra"
)

type Opts struct {
	F             *factory.Factory
	ID            string
	Output        string
	Name          string
	Region        string
	ProjectID     string
	Loadbalancers api.LoadBalancers
	Loadbalancer  api.LoadBalancer
	I             Interface
}

func NewProjectLoadbalancerCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	lbProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "loadbalancer",
		Aliases:               []string{"l", "lb", "load", "lbc"},
		Short:                 "project scoped loadbalancer",
		Long:                  grammar.ProjectLoadbalancerHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
		},
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	lbProj.AddGroup(&cobra.Group{ID: "actions", Title: "Action commands"})
	o := &Opts{F: f}
	o.I = NewLoadbalancerInterface(o)
	lbProj.AddCommand(NewProjectLoadbalancerListCommand(o))
	lbProj.AddCommand(NewProjectLoadbalancerCreateCommand(o))
	lbProj.AddCommand(NewProjectLoadbalancerDeleteCommand(o))
	return lbProj
}
