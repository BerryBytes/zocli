package cmd

import (
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/cmd/application"
	"github.com/berrybytes/zocli/pkg/cmd/apply"
	"github.com/berrybytes/zocli/pkg/cmd/auth"
	"github.com/berrybytes/zocli/pkg/cmd/environment"
	"github.com/berrybytes/zocli/pkg/cmd/organization"
	"github.com/berrybytes/zocli/pkg/cmd/organization/member"
	"github.com/berrybytes/zocli/pkg/cmd/project"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/fs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:           "zocli",
	Short:         "zocli is a CLI tool for 01cloud.com",
	Long:          grammar.RootHelp,
	RunE:          RootCmdRunE,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func RootCmdRunE(cmd *cobra.Command, _ []string) error {
	version, _ := cmd.Flags().GetBool("version")

	if version {
		logrus.Print(grammar.Version)
		return nil
	}

	return cmd.Help()
}

// Execute runs the main function of the program.
func Execute(f *factory.Factory) error {
	// Check if the configuration folder has been created.
	if !f.ConfigCreated {
		// Run the CheckConfigDir function in a separate goroutine.
		fs.CheckConfigDir(f)
		f.ConfigCreated = true
	}

	// Check if the user is logged in.
	fs.CheckIsLoggedIn(f)

	// Create a couple of global flags for the CLI.
	Cmd.PersistentFlags().BoolVarP(&f.Verbose, "verbose", "v", false, "Show verbose output for commands")
	Cmd.PersistentFlags().BoolVarP(&f.Quiet, "quiet", "q", false, "Quiet mode for commands")
	Cmd.PersistentFlags().BoolVarP(&f.NoInteractive, "no-interactive", "", false, "Disable interactive mode for commands")
	// Add a new command group to the CLI.
	Cmd.AddGroup(&cobra.Group{
		ID:    "basic",
		Title: "Basic Commands",
	})

	// Add the authentication command to the CLI.
	Cmd.AddCommand(auth.NewAuthCommand(f))
	Cmd.AddCommand(project.NewProjectCommand(f))
	Cmd.AddCommand(application.NewApplicationCommand(f))
	Cmd.AddCommand(environment.NewEnvironmentCommand(f))
	Cmd.AddCommand(apply.NewApplyManifestCommand(f))
	Cmd.AddCommand(organization.NewOrganizationCommand(f))
	Cmd.AddCommand(member.NewOrganizationMembersCommand(f))

	// Add global flags that will dynamically change the CLI's behavior.
	Cmd.Flags().Bool("help", false, "Show help for commands")
	Cmd.Flags().Bool("version", false, "Show version for commands")
	Cmd.Flags().Bool("update", false, "Show version for commands")
	err := Cmd.Execute()
	if err != nil {
		f.Printer.Errorf("Error executing command: %s", err.Error())
		f.Printer.Print(grammar.CliStatusCodeInfo)
		return err
	}
	return nil
}
