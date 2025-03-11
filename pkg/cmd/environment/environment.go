package environment

import (
	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/spf13/cobra"
)

type Opts struct {
	F               *factory.Factory
	I               EnvironmentInterface
	ApplicationID   string
	ApplicationName string
	Output          string
	EnvID           string
	EnvName         string
	UpdatedName     string
	Wide            bool
	ShowPassword    bool
	Environment     []*api.SingleEnvironment
	Multi           *api.MultipleEnvironment
	Overview        *api.EnvironmentOverview
}

func NewEnvironmentCommand(f *factory.Factory) *cobra.Command {
	o := Opts{F: f}
	o.I = NewEnvironmentInterface(&o)
	env := &cobra.Command{
		Use:     "environment",
		Aliases: []string{"e", "en", "env", "envi", "envir", "environm", "environme", "envs"},
		Short:   "environment level commands",
		Long:    grammar.EnvironmentCommandHelp,
		GroupID: "basic",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	env.AddGroup(&cobra.Group{
		ID:    "basic",
		Title: "Basic commands to maintain environment",
	})

	env.AddGroup(&cobra.Group{
		ID:    "advanced",
		Title: "Advanced commands to maintain environment",
	})

	env.AddCommand(NewEnvironmentGetCommand(&o))
	env.AddCommand(NewEnvironmentStartCommand(&o))
	env.AddCommand(NewEnvironmentStopCommand(&o))
	env.AddCommand(NewEnvironmentDeleteCommand(&o))
	env.AddCommand(NewEnvironmentOverviewCommand(&o))
	env.AddCommand(NewEnvironmentRenameCommand(&o))
	return env
}
