package vars

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectVarsDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	delProjVar := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "delete",
		Aliases:               []string{"del", "d", "de", "dele", "delete", "remove", "r"},
		Short:                 "project scoped variables",
		GroupID:               "actions",
		Long:                  grammar.ProjectVarsDeleteHelp,
		Run:                   opts.delVar,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(opts.F)
		},
	}

	delProjVar.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	delProjVar.Flags().Int64VarP(&opts.VarID, "vid", "I", -1, "variable id")
	delProjVar.Flags().StringVarP(&opts.VarName, "vname", "k", "", "variable key")
	delProjVar.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	delProjVar.Flags().BoolVarP(&opts.ShowSecret, "show", "s", false, "shows plain text of secret")
	return delProjVar
}

func (o *Opts) delVar(cmd *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.AllVars = o.getProjectDetail(o.ID).Variables
	}

	if o.Name != "" {
		project := o.getProjectDetailByName()
		o.ID = fmt.Sprint(project.ID)
		o.AllVars = project.Variables
	}
	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.AllVars = o.getProjectDetail(fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)).Variables
	}
	o.printVars(o.AllVars)

	if len(o.AllVars) == 0 {
		o.F.Printer.Fatal(1, "no any variables")
		return
	}

	if o.VarID != -1 {
		for _, each := range o.AllVars {
			if each.Id == int(o.VarID) {
				o.confirm(each)
				o.F.Printer.Exit(0)
			}
		}
		o.F.Printer.Fatal(1, "no such id found for variable")
	}

	if o.VarName != "" {
		for _, each := range o.AllVars {
			if each.Key == o.VarName {
				o.confirm(each)
				o.F.Printer.Exit(0)
			}
		}
		o.F.Printer.Fatal(1, "no such key found for variable")
	}

	_ = cmd.Help()
}

func (o *Opts) confirm(vars api.Variable) {
	var proceed bool
	prompt := &survey.Confirm{
		Message: "Do you sure want to delete the Variable with ID " + fmt.Sprint(vars.Id) + " ?",
		Default: false,
	}

	err := survey.AskOne(prompt, &proceed)
	if err != nil {
		o.F.Printer.Fatalf(5, "cannot proceed.\nErr: %v\n", err)
	}

	if !proceed {
		o.F.Printer.Exit(0)
	}
	o.deleteVar(vars)
}

func (o *Opts) deleteVar(vars api.Variable) {
	var postVariable []api.Variable
	for _, each := range o.AllVars {
		if each != vars {
			postVariable = append(postVariable, each)
		}
	}
	o.AllVars = postVariable

	a := struct {
		Variables []api.Variable `json:"variables"`
	}{
		Variables: postVariable,
	}

	body, err := json.Marshal(a)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot Marshal")
	}

	header := o.F.GetAuth()
	reqConfig := defaults.UpdateVarProject(o.F, map[string]interface{}{"body": body, "headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully Deleted variable")
}
