package environment

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewEnvironmentRenameCommand(o *Opts) *cobra.Command {
	env := &cobra.Command{
		Use:     "rename",
		Aliases: []string{"re", "r", "ren", "renam", "rena"},
		Short:   "rename environments",
		Long:    grammar.EnvironmentRenameHelp,
		GroupID: "basic",
		Run:     o.I.RenameRunner,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				o.EnvID, _ = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	env.Flags().StringVarP(&o.EnvID, "id", "i", "", "environment id")
	env.Flags().StringVarP(&o.UpdatedName, "name", "n", "", "environment updated name")

	return env
}

func (o *Opts) RenameRunner(cmd *cobra.Command, _ []string) {
	if o.EnvID != "" && o.UpdatedName != "" {
		utils.ConfirmIfToProceed("Do you sure want to rename ?", o.F)
		o.I.RenameEnv()
		return
	}

	if o.EnvID != "" && o.UpdatedName == "" {
		o.I.AskChanges()
		utils.ConfirmIfToProceed("Do you sure want to rename ?", o.F)
		o.I.RenameEnv()
		return
	}

	_ = cmd.Help()
}

func (o *Opts) AskChanges() {
	fmt.Print("Environment New Name: ")
	text, err := o.F.Term.ReadInput()
	if err != nil {
		o.F.Printer.Fatal(1, err)
	}
	text = strings.ReplaceAll(text, "\n", "")
	o.UpdatedName = text
}

func (o *Opts) RenameEnv() {
	headers := o.F.GetAuth()
	body := []byte(`{"name":"` + o.UpdatedName + `"}`)
	reqConf := defaults.EnvironmentRenameByID(o.F, map[string]interface{}{"headers": headers, "body": body})
	reqConf.URL = strings.ReplaceAll(reqConf.URL, "<:id>", o.EnvID)
	_ = reqConf.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Successfully renamed environment\n")
	o.F.Printer.Exit(0)
}
