package application

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/spf13/cobra"
)

type Opts struct {
	F               *factory.Factory
	I               Interface
	ProjectID       string
	ProjectName     string
	Output          string
	ApplicationID   string
	ApplicationName string
	List            api.ApplicationList
	NewList         ApptList
	Change          api.Application
	ID              string
	Temporary       bool
}

type ApptList struct {
	Apps []models.App `json:"apps" yaml:"apps"`
}

func NewApplicationCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	opts.I = NewAppInterface(opts)
	app := &cobra.Command{
		Use:     "app",
		Aliases: []string{"a", "ap", "apps"},
		Short:   "app related commands",
		Long:    grammar.ApplicationHelp,
		GroupID: "basic",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	app.AddGroup(&cobra.Group{ID: "basic", Title: "basic commands"})
	app.AddGroup(&cobra.Group{ID: "context", Title: "Maintain context or defaults"})

	app.AddCommand(NewApplicationGetCommand(opts))
	app.AddCommand(NewApplicationRenameCommand(opts))
	app.AddCommand(NewApplicationDeleteCommand(opts))
	app.AddCommand(NewApplicationUseDefaultCommand(f))
	app.AddCommand(NewApplicationGetDefaultCommand(f))
	app.AddCommand(NewApplicationDeleteDefaultCommand(f))

	return app
}

func (o *Opts) GetProjectDetailByName() *api.Project {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectByName(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:name>", o.ProjectName)

	res := reqConfig.Request()
	if res == nil {
		return nil
	}

	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	return o.GetProjectDetail(fmt.Sprint(project.ID))
}

func (o *Opts) GetProjectDetail(id string) *api.Project {
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
