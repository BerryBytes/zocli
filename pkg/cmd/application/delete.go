package application

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewApplicationDeleteCommand(o *Opts) *cobra.Command {
	app := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d", "del", "de", "dele", "delete"},
		Short:   "delete apps on projects",
		Long:    grammar.ApplicationDeleteHelp,
		GroupID: "basic",
		Run:     o.I.DeleteRunner,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				o.ApplicationID, _ = validation.GetIDOrName(args[0])
				validation.CheckValidID(o.F, o.ApplicationID)
			}
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}
	app.Flags().StringVarP(&o.ProjectID, "pid", "p", "", "project id")
	app.Flags().StringVarP(&o.ProjectName, "pname", "N", "", "project name")
	app.Flags().StringVarP(&o.ApplicationID, "id", "i", "", "application id")

	return app
}

func (o *Opts) DeleteRunner(cmd *cobra.Command, _ []string) {
	if o.ProjectID != "" && o.ProjectName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ProjectID != "" {
		validation.CheckValidID(o.F, o.ProjectID)
	}

	if o.ProjectName != "" {
		proj := o.I.GetProjectDetailByName()
		o.ProjectID = fmt.Sprint(proj.ID)
	}

	if o.ProjectID == "" && o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.ProjectID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
	}

	if o.ProjectID == "" && o.ProjectName == "" {
		_ = cmd.Help()
		return
	}

	o.I.GetApps()
	o.I.PrintApps()

	if o.ApplicationID == "" {
		o.F.Printer.Print("Application ID : ")
		var id string
		_, err := fmt.Scanln(&id)
		if err != nil {
			o.F.Printer.Fatal(1, o.F.IO.ColorScheme().FailureIcon(), " invalid id")
			return
		}
		id, _ = validation.GetIDOrName(id)
		if id == "" {
			o.F.Printer.Fatal(1, "invalid id")
			return
		}
		o.ApplicationID = id
	}

	o.checkIfValidID()
	o.I.AskDeleteConfirmation()
	o.I.DeleteApplication()
}

func (o *Opts) AskDeleteConfirmation() {
	proceed := false
	prompt := &survey.Confirm{
		Message: "Do you sure want to delete the application ?",
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

func (o *Opts) DeleteApplication() {
	headers := o.F.GetAuth()
	reqConfig := defaults.DeleteApplication(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(o.ApplicationID))
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Application deleted of id '"+o.ApplicationID+"'\n")
}
