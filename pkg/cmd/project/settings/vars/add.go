package vars

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectVarsAddCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	addVar := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "add",
		Aliases:               []string{"a", "ad", "new"},
		Short:                 "add variable",
		GroupID:               "actions",
		Long:                  grammar.ProjectVarsUpdateHelp,
		Example:               grammar.ProjectVarsUpdateExample,
		Run:                   opts.AddVarsRunner,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
		},
	}

	addVar.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	addVar.Flags().StringVarP(&opts.VarName, "name", "n", "", "variable name")
	addVar.Flags().StringVarP(&opts.Value, "value", "p", "", "variable value")
	addVar.Flags().StringVarP(&opts.Type, "type", "t", "", "variable type")
	return addVar
}

func (o *Opts) AddVarsRunner(_ *cobra.Command, _ []string) {
	o.F.Printer.Printf("NOTE: WARNING THIS COMMAND REPLACES ALL THE VARIABLES ON THE PROJECT AND CANNOT BE UNDONE")
	utils.ConfirmIfToProceed("Do you sure want to continue?", o.F)
	o.checkProjectID()
	o.checkVarName()
	o.checkVarValue()
	o.checkVarType()
	newVar := api.Variable{
		Key:   o.VarName,
		Type:  o.Type,
		Value: o.Value,
	}
	o.AllVars = append(o.AllVars, newVar)
	o.postChanges(newVar)
}

// checkVarType
//
// check if variable type is provided
// if not, ask for variable type
// and check if variable type is valid i.e. of type secret or normal
func (o *Opts) checkVarType() {
	if o.Type == "" {
		o.F.Printer.Print("Variable type: ")
		val, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, err.Error())
			return
		}
		o.Type = strings.ReplaceAll(val, "\n", "")
	} else {
		o.F.Printer.Print("Variable type: " + o.Type + "\n")
	}
	if o.Type != "secret" && o.Type != "normal" {
		o.F.Printer.Fatal(1, "type must be either secret or normal")
		return
	}
}

// checkVarValue
//
// check if variable value is provided
// if not, ask for variable value
func (o *Opts) checkVarValue() {
	if o.Value == "" {
		o.F.Printer.Print("Variable value: ")
		val, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, err.Error())
			return
		}
		o.Value = strings.ReplaceAll(val, "\n", "")
	} else {
		o.F.Printer.Print("Variable value: " + o.Value + "\n")
	}
}

// checkVarName
//
// check if variable name is provided
// if not, ask for variable name
// if yes, check if variable name is valid
func (o *Opts) checkVarName() {
	if o.VarName == "" {
		o.F.Printer.Print("Variable name: ")
		val, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, err.Error())
			return
		}
		o.VarName = strings.ReplaceAll(val, "\n", "")
	} else {
		o.F.Printer.Print("Variable name: " + o.VarName + "\n")
	}
	r, _ := regexp.Compile("^[a-zA-Z0-9_]*$")
	pass := r.MatchString(o.VarName)
	if !pass {
		o.F.Printer.Fatal(1, "Err: key must be only alphanumeric")
		return
	}
}

// checkProjectID
//
// check if project id is provided
// if not, check if there is an active context
// if not, ask for project id
// if yes, check if project id is valid
// if yes, set project id
func (o *Opts) checkProjectID() {
	if o.ID == "" && o.F.Config.ActiveContext != nil {
		o.ID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
	}
	if o.ID == "" {
		o.F.Printer.Printf("Project ID: ")
		val, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, err.Error())
			return
		}
		val = strings.ReplaceAll(val, "\n", "")
		validation.CheckValidID(o.F, val)
		o.ID = val
	} else {
		o.F.Printer.Print("Project ID: " + o.ID + "\n")
	}
}
