package settings

import (
	"github.com/berrybytes/zocli/pkg/cmd/project/settings/loadbalancer"
	"github.com/berrybytes/zocli/pkg/cmd/project/settings/permissions"
	"github.com/berrybytes/zocli/pkg/cmd/project/settings/vars"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/spf13/cobra"
)

func NewSettingsCommand(f *factory.Factory) *cobra.Command {
	settings := &cobra.Command{
		Use:     "settings",
		Short:   "settings related to project",
		Aliases: []string{"set", "s", "setting", "sett"},
		GroupID: "advanced",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	settings.AddCommand(vars.NewProjectVarsCommand(f))
	settings.AddCommand(permissions.NewProjectPermissionsCommand(f))
	settings.AddCommand(loadbalancer.NewProjectLoadbalancerCommand(f))
	return settings
}
