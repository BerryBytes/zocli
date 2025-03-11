package environment

import (
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewEnvironmentStartCommand(o *Opts) *cobra.Command {
	env := &cobra.Command{
		Use:     "start",
		Aliases: []string{"sta", "b", "boot"},
		Short:   "start environments",
		Long:    grammar.EnvironmentStartHelp,
		GroupID: "basic",
		Run:     o.I.StartRunner,
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

func (o *Opts) StartRunner(cmd *cobra.Command, _ []string) {
	if o.EnvID != "" && o.EnvName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.EnvID != "" {
		o.I.StartEnvByID()
		return
	}

	if o.EnvName != "" {
		o.F.Printer.Println("Please use id for now, as this is still under development")
		o.F.Printer.Exit(0)
	}
	_ = cmd.Help()
}

func (o *Opts) StartEnvByID() {
	headers := o.F.GetAuth()
	body := []byte(`{}`)
	reqConf := defaults.EnvironmentStartByID(o.F, map[string]interface{}{"headers": headers, "body": body})
	reqConf.URL = strings.ReplaceAll(reqConf.URL, "<:id>", o.EnvID)
	res := reqConf.Request()
	o.F.Printer.Println(res.Message)
	o.F.Printer.Exit(0)
}
