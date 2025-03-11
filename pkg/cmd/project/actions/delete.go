package actions

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	deleteProject := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "delete",
		Aliases:               []string{"d", "del", "de", "dele"},
		Short:                 "delete project",
		GroupID:               "basic",
		Long:                  grammar.ProjectDeleteHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.deleteCheck,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(opts.F)
		},
	}

	deleteProject.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	deleteProject.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	return deleteProject
}

func (o *Opts) deleteCheck(cmd *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.fetchProjByID(o.ID)
		return
	}

	if o.Name != "" {
		o.fetchProjByName(o.Name)
	}

	_ = cmd.Help()
}

func (o *Opts) fetchProjByID(id string) {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectDetail(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(id))
	res := reqConfig.Request()
	if res == nil {
		return
	}
	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	o.project = &project
	o.printInfo()
	o.confirmAndDelete()
}

func (o *Opts) fetchProjByName(name string) {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectByName(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:name>", name)

	res := reqConfig.Request()
	if res == nil {
		return
	}

	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	o.project = &project
	o.printInfo()
	o.confirmAndDelete()
}

func (o *Opts) confirmAndDelete() {
	utils.ConfirmIfToProceed("Do you sure want to delete the project ?", o.F)
	fmt.Print("Enter the project name to proceed on : ")
	text, err := o.F.Term.ReadInput()
	if err != nil {
		o.F.Printer.Fatal(5, err)
	}
	projectName := strings.TrimSpace(strings.Trim(text, "\n"))
	if projectName != o.project.Name {
		o.F.Printer.Fatal(1, "invalid project name")
	}

	header := o.F.GetAuth()
	reqConfig := defaults.DeleteProject(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(o.project.ID))
	res := reqConfig.Request()
	if res == nil {
		return
	}

	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully deleted '"+o.project.Name+"'")
	o.F.Printer.Exit(0)
}

func (o *Opts) printInfo() {
	columns := []string{"id", "name", "email", "subscription", "active"}
	tablePrinter := table.New(o.F, 0)
	for _, column := range columns {
		switch column {
		case "id":
			tablePrinter.AddField("ID")
			tablePrinter.AddField(fmt.Sprintf("%d", o.project.ID))
			tablePrinter.EndRow()
		case "name":
			tablePrinter.AddField("NAME")
			tablePrinter.AddField(o.project.Name)
			tablePrinter.EndRow()
		case "description":
			tablePrinter.AddField("DESCRIPTION")
			tablePrinter.AddField(o.project.Description)
			tablePrinter.EndRow()
		case "email":
			tablePrinter.AddField("EMAIL")
			tablePrinter.AddField(o.project.User.Email)
			tablePrinter.EndRow()
		case "subscription":
			tablePrinter.AddField("SUBSCRIPTION")
			tablePrinter.AddField(o.project.Subscription.Name)
			tablePrinter.EndRow()
		case "active":
			tablePrinter.AddField("ACTIVE")
			tablePrinter.AddField(fmt.Sprintf("%v", o.project.Active))
			tablePrinter.EndRow()
		}
	}
	_ = tablePrinter.Print()
}
