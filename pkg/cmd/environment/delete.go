package environment

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewEnvironmentDeleteCommand(o *Opts) *cobra.Command {
	env := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d", "del", "de", "dele", "delete"},
		Short:   "delete environments",
		Long:    grammar.EnvironmentDeleteHelp,
		GroupID: "basic",
		Run:     o.I.DeleteRunner,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				o.EnvID, o.EnvName = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	env.Flags().StringVarP(&o.EnvID, "id", "i", "", "environment id")
	env.Flags().StringVarP(&o.EnvName, "name", "n", "", "environment name")

	return env
}

func (o *Opts) DeleteRunner(cmd *cobra.Command, _ []string) {
	if o.EnvID != "" && o.EnvName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.EnvID != "" {
		o.I.DeleteEnvByID()
		return
	}

	if o.EnvName != "" {
		o.F.Printer.Println("Please use id for now, as this is still under development")
		o.F.Printer.Exit(0)
	}
	_ = cmd.Help()
}

func (o *Opts) DeleteEnvByID() {
	o.I.AskDeleteConfirmation()
	headers := o.F.GetAuth()
	reqConf := defaults.EnvironmentDeleteByID(o.F, map[string]interface{}{"headers": headers})
	reqConf.URL = strings.ReplaceAll(reqConf.URL, "<:id>", o.EnvID)
	_ = reqConf.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print("Successfully deleted environment\n")
	o.F.Printer.Exit(0)
}

func (o *Opts) AskDeleteConfirmation() {
	proceed := false
	prompt := &survey.Confirm{
		Message: "Do you sure want to delete the environment ?",
		Default: false,
	}

	err := survey.AskOne(prompt, &proceed)
	if err != nil {
		o.F.Printer.Fatalf(5, "cannot proceed")
	}

	if !proceed {
		o.F.Printer.Exit(0)
	}
}
