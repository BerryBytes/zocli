package loadbalancer

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectLoadbalancerDeleteCommand(opts *Opts) *cobra.Command {
	delPermProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "delete",
		Aliases:               []string{"d", "de", "del", "dele", "delete", "remove", "rem", "r"},
		Short:                 "delete project scoped permissions",
		Long:                  grammar.ProjectLoadbalancerDeleteHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.I.DeleteLoadbalancer,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)

			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])

			}
		},
	}

	delPermProj.Flags().StringVarP(&opts.ID, "id", "i", "", "loadbalancer id")
	delPermProj.Flags().StringVarP(&opts.Output, "out", "o", "", "output in json or yaml")
	return delPermProj
}

func (o *Opts) DeleteLoadbalancer(cmd *cobra.Command, _ []string) {
	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.I.GetByIdLoadbalancer()
	}
	if o.ID == "" {
		_ = cmd.Help()
		o.F.Printer.Exit(0)
		return
	}
	if o.Loadbalancer.Name == "" {
		o.F.Printer.Fatal(1, "no loadbalancer found")
		return
	}
	utils.ConfirmIfToProceed("Are you sure you want to delete the loadbalancer?", o.F)
	o.Delete()
}

func (o *Opts) Delete() {
	reqConfig := defaults.DeleteLoadbalancer(o.F, map[string]interface{}{"headers": o.F.GetAuth()})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully deleted loadbalancer with id '"+fmt.Sprint(o.ID)+"'")
}
