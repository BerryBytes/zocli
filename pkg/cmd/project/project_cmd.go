package project

import (
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/cmd/project/actions"
	"github.com/berrybytes/zocli/pkg/cmd/project/context"
	"github.com/berrybytes/zocli/pkg/cmd/project/get"
	"github.com/berrybytes/zocli/pkg/cmd/project/overview"
	"github.com/berrybytes/zocli/pkg/cmd/project/settings"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/spf13/cobra"
)

func NewProjectCommand(f *factory.Factory) *cobra.Command {
	project := &cobra.Command{
		Use:     "project",
		Aliases: []string{"proj", "pro", "projects", "p"},
		Short:   "project related commands",
		Long:    grammar.ProjectHelp,
		GroupID: "basic",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	project.AddGroup(&cobra.Group{
		ID:    "basic",
		Title: "Basic commands for interacting with project",
	})

	project.AddGroup(&cobra.Group{
		ID:    "context",
		Title: "Maintain context or defaults",
	})

	project.AddGroup(&cobra.Group{
		ID:    "advanced",
		Title: "Advanced commands to maintain project",
	})

	project.AddCommand(get.NewProjectGetCommand(f))
	project.AddCommand(actions.NewProjectEnableCommand(f))
	project.AddCommand(actions.NewProjectDisableCommand(f))
	project.AddCommand(actions.NewProjectRenameCommand(f))
	project.AddCommand(actions.NewProjectDeleteCommand(f))
	project.AddCommand(overview.NewProjectOverviewCommand(f))
	project.AddCommand(settings.NewSettingsCommand(f))
	project.AddCommand(context.NewProjectUseDefaultCommand(f))
	project.AddCommand(context.NewProjectGetDefaultCommand(f))
	project.AddCommand(context.NewProjectDeleteDefaultCommand(f))
	return project
}
