package loadbalancer

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectLoadbalancerCreateCommand(opts *Opts) *cobra.Command {
	listPermProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "create",
		Aliases:               []string{"c", "cr"},
		Short:                 "create project scoped loadbalancer",
		Long:                  grammar.ProjectLoadbalancerCreateHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.I.CreateLoadbalancer,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
		},
	}
	listPermProj.Flags().StringVarP(&opts.ProjectID, "id", "i", "", "project id")
	listPermProj.Flags().StringVarP(&opts.Name, "name", "n", "", "loadbalancer name")
	listPermProj.Flags().StringVarP(&opts.Region, "region", "r", "", "region name")
	return listPermProj
}

func (o *Opts) CreateLoadbalancer(cmd *cobra.Command, _ []string) {
	o.F.IO.StartProgressIndicator()
	err := o.ValidateCreation()
	if err != nil {
		o.F.IO.StopProgressIndicator()
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Print("Missing required fields:")
		o.F.Printer.Fatal(1, err.Error())
	}
	o.ApplyLoadbalancer()
	o.F.IO.StopProgressIndicator()
	_ = cmd.Help()
}

func (o *Opts) ApplyLoadbalancer() {
	headers := o.F.GetAuth()
	var lb api.LoadBalancer
	lb.Name = o.Name
	pid, _ := strconv.Atoi(o.ProjectID)
	lb.ProjectID = uint64(pid)
	lb.Region = o.Region
	lbData, _ := json.Marshal(lb)
	reqConfig := defaults.CreateProjectLoadbalancer(o.F, map[string]interface{}{"headers": headers, "body": lbData})
	o.F.IO.StopProgressIndicator()
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Successfully created loadbalancer with name '" + lb.Name + "'\n")
	o.F.Printer.Exit(0)
}

func (l *Opts) ValidateCreation() error {
	if l.Name == "" {
		return errors.New("provide project name")
	}
	if l.ProjectID == "" {
		return errors.New("provide name")
	}
	if l.Region == "" {
		return errors.New("provide region")
	}
	return nil
}
