package actions

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectEnableCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	enableProject := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "enable",
		Aliases:               []string{"e", "ena", "able", "enab"},
		Short:                 "enable project",
		GroupID:               "basic",
		Long:                  grammar.ProjectEnableHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.enable,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(opts.F)
		},
	}

	enableProject.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	enableProject.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")

	return enableProject
}

func (o *Opts) enable(cmd *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.enableByID(o.ID)
		return
	}

	if o.Name != "" {
		o.enableByName(o.Name)
	}

	_ = cmd.Help()
}

func (o *Opts) enableByID(id string) {
	utils.ConfirmIfToProceed("Do you sure want to enable the project ?", o.F)
	header := o.F.GetAuth()
	body := []byte(`{"is_operational": true}`)
	reqConfig := defaults.GetProjectEnableByID(o.F, map[string]interface{}{"headers": header, "body": body})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", id)

	res := reqConfig.Request()
	if res == nil {
		return
	}
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully activated project.")
	o.F.Printer.Exit(0)
}

func (o *Opts) enableByName(name string) {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectByName(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:name>", name)

	res := reqConfig.Request()
	if res == nil {
		return
	}

	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	o.enableByID(fmt.Sprint(project.ID))
}
