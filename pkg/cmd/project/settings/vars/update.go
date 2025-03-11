package vars

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

func NewProjectVarsUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	delProjVar := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "update",
		Aliases:               []string{"u", "up", "upda", "upd"},
		Short:                 "update variables",
		GroupID:               "actions",
		Long:                  grammar.ProjectVarsUpdateHelp,
		Example:               grammar.ProjectVarsUpdateExample,
		Run:                   opts.updatevar,
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
	delProjVar.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	delProjVar.Flags().BoolVarP(&opts.ShowSecret, "show", "s", false, "shows plain text of secret")
	return delProjVar
}

func (o *Opts) updatevar(cmd *cobra.Command, _ []string) {
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
		o.ID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
	}

	if len(o.AllVars) == 0 {
		o.F.Printer.Fatal(1, "no variables")
		return
	}

	if o.ID == "" {
		_ = cmd.Help()
		return
	}

	o.printVars(o.AllVars)

	if o.VarID == -1 {
		fmt.Printf("ID to update : ")
		var id string
		_, _ = fmt.Scanln(&id)
		id, _ = validation.GetIDOrName(id)
		if id == "" {
			o.F.Printer.Fatal(1, "invalid id")
		}
		o.VarID, _ = strconv.ParseInt(id, 10, 64)
	}
	for _, each := range o.AllVars {
		if each.Id == int(o.VarID) {
			o.askUpdates(each)
			o.F.Printer.Exit(0)
		}
	}
	o.F.Printer.Fatal(1, "no such id found for variable")
}

func (o *Opts) askUpdates(vars api.Variable) {
	original := vars
	var updates string
	input, _ := readline.New("Key : ")
	defer input.Close()

	datax := vars.Key
	data2 := []byte(datax)
	_, _ = input.WriteStdin(data2)

	updates, _ = input.Readline()
	r, _ := regexp.Compile("^[a-zA-Z0-9_]*$")
	pass := r.MatchString(updates)
	if !pass {
		o.F.Printer.Fatal(1, "Err: key must be only alphanumeric")
	}

	vars.Key = updates

	input, _ = readline.New("Value : ")
	defer input.Close()

	data2 = []byte(vars.Value)
	_, _ = input.WriteStdin(data2)

	updates, _ = input.Readline()
	vars.Value = updates

	input, _ = readline.New("Type : ")
	defer input.Close()

	data2 = []byte(vars.Type)
	_, _ = input.WriteStdin(data2)

	updates, _ = input.Readline()
	if updates != "secret" && updates != "normal" {
		o.F.Printer.Fatal(1, "type must be either secret or normal")
	}
	vars.Type = updates

	if original == vars {
		o.F.Printer.Print("no changes detected")
		o.F.Printer.Exit(0)
	}
	o.postChanges(vars)
}

func (o *Opts) postChanges(vars api.Variable) {
	var postVariable []api.Variable
	for _, each := range o.AllVars {
		appendThis := each
		if each.Id == vars.Id {
			appendThis = vars
		}
		postVariable = append(postVariable, appendThis)
	}
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
	fmt.Println(o.F.IO.ColorScheme().SuccessIcon(), " Successfully Updated variable")
}
