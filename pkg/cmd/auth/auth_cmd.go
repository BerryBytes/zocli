package auth

import (
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/cmd/auth/login"
	"github.com/berrybytes/zocli/pkg/cmd/auth/logout"
	"github.com/berrybytes/zocli/pkg/cmd/auth/status"
	"github.com/berrybytes/zocli/pkg/utils/factory"

	"github.com/spf13/cobra"
)

// NewAuthCommand
//
// creates a new instance of the cobra.Command as the "auth" command.
// It takes a pointer to a factory.Factory object as a parameter.
// The "auth" command is used to authenticate a user.
// It has a short description and a longer description taken from the grammar.AuthHelp constant.
// It also has an example command taken from the grammar.AuthExamples constant.
// The "auth" command is part of the "basic" group.
// When the "auth" command is executed, it runs the help command of the cobra.Command object.
// If there is an error running the help command, it prints the error message to the console.
// It returns the created "auth" command.
func NewAuthCommand(f *factory.Factory) *cobra.Command {
	auth := &cobra.Command{
		Use:     "auth",
		Short:   "Authenticates a user",
		Long:    grammar.AuthHelp,
		Example: grammar.AuthExamples,
		GroupID: "basic",
		RunE:    authCmd,
	}

	// add subcommands
	auth.AddCommand(login.NewLoginCommand(f))
	auth.AddCommand(status.NewStatusCmd(f))
	auth.AddCommand(logout.NewLogOutCmd(f))
	return auth
}

func authCmd(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}
