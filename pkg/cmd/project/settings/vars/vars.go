package vars

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/spf13/cobra"
)

type Opts struct {
	F          *factory.Factory
	ID         string
	Name       string
	Output     string
	ShowSecret bool
	VarID      int64
	VarName    string
	AllVars    []api.Variable
	Value      string
	Type       string
}

func NewProjectVarsCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	varProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "vars",
		Aliases:               []string{"v", "va", "var", "variable", "variables", "vari"},
		Short:                 "project scoped variables",
		Long:                  grammar.ProjectVarsHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
		},
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	varProj.AddGroup(&cobra.Group{ID: "actions", Title: "Action commands"})
	varProj.AddCommand(NewProjectVarsGetCommand(f))
	varProj.AddCommand(NewProjectVarsDeleteCommand(f))
	varProj.AddCommand(NewProjectVarsUpdateCommand(f))
	varProj.AddCommand(NewProjectVarsAddCommand(f))
	return varProj
}

func (o *Opts) getProjectDetailByName() *api.Project {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectByName(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:name>", o.Name)

	res := reqConfig.Request()
	if res == nil {
		return nil
	}

	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	return o.getProjectDetail(fmt.Sprint(project.ID))
}

func (o *Opts) getProjectDetail(id string) *api.Project {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectDetail(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(id))
	res := reqConfig.Request()
	if res == nil {
		return nil
	}
	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}
	return &project
}
