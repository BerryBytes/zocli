package environment

import (
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewEnvironmentStopCommand(o *Opts) *cobra.Command {
	env := &cobra.Command{
		Use:     "stop",
		Aliases: []string{"s", "st", "sto", "halt", "hal"},
		Short:   "stop environments",
		Long:    grammar.EnvironmentStopHelp,
		GroupID: "basic",
		Run:     o.I.StopRunner,
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

func (o *Opts) StopRunner(cmd *cobra.Command, _ []string) {
	if o.EnvID != "" && o.EnvName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.EnvID != "" {
		o.I.StopEnvByID()
		return
	}

	if o.EnvName != "" {
		o.F.Printer.Println("Please use id for now, as this is still under development")
		o.F.Printer.Exit(0)
		return
	}
	_ = cmd.Help()
}

func (o *Opts) StopEnvByID() {
	headers := o.F.GetAuth()
	body := []byte(`{}`)
	reqConf := defaults.EnvironmentStopByID(o.F, map[string]interface{}{"headers": headers, "body": body})
	reqConf.URL = strings.ReplaceAll(reqConf.URL, "<:id>", o.EnvID)
	res := reqConf.Request()
	o.F.Printer.Println(res.Message)
	o.F.Printer.Exit(0)
}
