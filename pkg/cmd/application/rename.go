package application

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

func NewApplicationRenameCommand(o *Opts) *cobra.Command {
	app := &cobra.Command{
		Use:     "rename",
		Aliases: []string{"r", "re", "renam", "ren"},
		Short:   "rename apps on projects",
		Long:    grammar.ApplicationRenameHelp,
		GroupID: "basic",
		Run:     o.I.RenameRunner,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				o.ApplicationID, _ = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	app.Flags().StringVarP(&o.ApplicationID, "id", "i", "", "app id")
	app.Flags().StringVarP(&o.ApplicationName, "name", "n", "", "app name")

	app.Flags().StringVarP(&o.ProjectID, "pid", "p", "", "project id")
	app.Flags().StringVarP(&o.ProjectName, "pname", "N", "", "project name")

	return app
}

func (o *Opts) RenameRunner(cmd *cobra.Command, _ []string) {
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
			o.F.Printer.Fatal(1, o.F.IO.ColorScheme().FailureIcon()+" Error reading input")
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
	o.askChanges()
	o.I.RenameApplication()
}

func (o *Opts) askChanges() {
	input, _ := readline.New("New Application Name : ")
	defer func(input *readline.Instance) {
		_ = input.Close()
	}(input)

	datax := o.Change.Name
	data2 := []byte(datax)
	_, _ = input.WriteStdin(data2)

	updates, _ := input.Readline()

	r, _ := regexp.Compile("^[a-zA-Z0-9_ -]{3,30}$")
	pass := r.MatchString(updates)
	if !pass {
		o.F.Printer.Fatal(1, "Only 3 to 30 alphanumeric, underscore, space & hyphen characters allowed!")
		return
	}

	o.ApplicationName = updates
}

func (o *Opts) checkIfValidID() {
	valid := false
	for _, app := range o.List.Applications {
		if fmt.Sprint(app.Id) == o.ApplicationID {
			valid = true
			o.Change = app
		}
	}

	if !valid {
		o.F.Printer.Fatal(1, "no such application found.")
	}
}

func (o *Opts) RenameApplication() {
	headers := o.F.GetAuth()
	body := []byte(`{"name":"` + o.ApplicationName + `"}`)
	reqConfig := defaults.RenameApplication(o.F, map[string]interface{}{"headers": headers, "body": body})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(o.ApplicationID))
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Application renamed to '"+o.ApplicationName+"'\n")
}
