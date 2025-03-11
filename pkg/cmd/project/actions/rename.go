package actions

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectRenameCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	enableProject := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "rename",
		Aliases:               []string{"r", "rena", "renam", "renm", "re"},
		Short:                 "rename project",
		Long:                  grammar.ProjectRenameHelp,
		GroupID:               "basic",
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.rename,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID = args[0]
			}
			middleware.LoggedIn(opts.F)
			validation.CheckValidID(f, opts.ID)
		},
	}

	enableProject.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	enableProject.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	return enableProject
}

func (o *Opts) rename(_ *cobra.Command, _ []string) {
	o.askID()
	o.askName()
	if o.ID != "" && o.Name != "" {
		header := o.F.GetAuth()

		body := []byte(`{"name": "` + o.Name + `"}`)

		reqConfig := defaults.RenameProject(o.F, map[string]interface{}{"headers": header, "body": body})
		reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)

		res := reqConfig.Request()
		if res == nil {
			return
		}

		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully name changed to '"+o.Name+"'")
		o.F.Printer.Exit(0)
	}
}

func (o *Opts) askName() {
	if o.Name == "" {
		fmt.Print("New Project Name: ")
		text, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, "cannot get input")
			return
		}
		o.Name = strings.TrimSpace(strings.Trim(text, "\n"))
	}
}

func (o *Opts) askID() {
	if o.ID == "" {
		fmt.Print("Project ID: ")
		text, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, "cannot get input")
			return
		}
		o.ID = strings.TrimSpace(strings.Trim(text, "\n"))
	}
	validation.CheckValidID(o.F, o.ID)
}
