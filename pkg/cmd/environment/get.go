package environment

import (
	"fmt"
	"strings"
	"time"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/formatter"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewEnvironmentGetCommand(o *Opts) *cobra.Command {
	env := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g", "ge", "list", "lis", "retrieve"},
		Short:   "get environments",
		Long:    grammar.EnvironmentGetHelp,
		GroupID: "basic",
		Run:     o.I.GetRunner,
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

	env.Flags().StringVarP(&o.ApplicationID, "aid", "a", "", "application id")
	env.Flags().StringVarP(&o.ApplicationName, "aname", "N", "", "application name")

	env.Flags().StringVarP(&o.Output, "out", "o", "", "output in json or yaml")
	env.Flags().BoolVarP(&o.ShowPassword, "password", "p", false, "show password for user [NOTE: CAN ONLY BE USED IN CASE OF ONE ENVIRONMENT]")
	return env
}

func (o *Opts) GetRunner(cmd *cobra.Command, _ []string) {
	if o.EnvID != "" && o.EnvName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}
	if o.EnvID != "" {
		o.I.FetchEnvironment(o.EnvID)
		o.I.PrintSingleEnvTable()
		return
	}
	if o.EnvName != "" {
		o.F.Printer.Println("Please use id for now, as this is still under development")
		o.F.Printer.Exit(0)
		return
	}

	if o.ApplicationID != "" {
		o.FetchEnvironments(o.ApplicationID)
		return
	}

	if o.ApplicationName != "" {
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Printf(" Please use id for now, as this is still under development")
		o.F.Printer.Exit(0)
		return
	}

	if o.F.Config.ActiveContext.DefaultApplication != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Application Value.\n")
		o.FetchEnvironments(fmt.Sprint(o.F.Config.ActiveContext.DefaultApplication))
		return
	}
	_ = cmd.Help()
}

func (o *Opts) FetchEnvironments(id string) {
	reqConfig := defaults.EnvironmentGetList(o.F, map[string]interface{}{"headers": o.F.GetAuth()})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(id))
	res := reqConfig.Request()
	var env api.MultipleEnvironment
	err := env.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal data")
		return
	}
	o.Multi = &env
	o.I.PrintMultiEnvs()
	o.F.Printer.Exit(0)
}

func (o *Opts) FetchEnvironment(id string) {
	headers := o.F.GetAuth()
	reqConfig := defaults.EnvironmentGet(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", id)
	res := reqConfig.Request()
	var env api.SingleEnvironment
	err := env.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal data")
		return
	}
	o.Environment = append(o.Environment, &env)
}

func (o *Opts) PrintSingleEnvTable() {
	if o.Output != "" {
		var toPrint api.EnvironmentPresenter
		err := toPrint.FromJSON(o.Environment[0])
		if err != nil {
			o.F.Printer.Fatal(9, "cannot unmarshal data")
			return
		}
		switch strings.ToLower(o.Output) {
		case "json":
			formatter.PrintJson(o.F, toPrint)
			o.F.Printer.Exit(0)
			return
		case "yaml":
			formatter.PrintYaml(o.F, toPrint)
			o.F.Printer.Exit(0)
			return
		}
	}
	headers := []string{"id", "name", "isrunning", "url", "username", "deployed", "created"}
	if o.ShowPassword {
		headers = append(headers, "password")
	}
	tableprinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tableprinter.HeaderRow(color, headers...)
	for _, env := range o.Environment {
		for _, header := range headers {
			switch header {
			case "id":
				tableprinter.AddField(fmt.Sprint(env.Environment.ID))
			case "name":
				tableprinter.AddField(env.Environment.Name)
			case "isrunning":
				isRunning := false
				for _, o := range env.Overview {
					if o.IsRunning {
						isRunning = true
					}
				}
				tableprinter.AddField(fmt.Sprint(isRunning))
			case "url":
				for _, o := range env.Overview {
					if strings.Contains(o.Name, "Url") {
						tableprinter.AddField(o.Value)
					}
				}
			case "username":
				val := ""
				for _, o := range env.Overview {
					if o.Name == "User Name" {
						val = o.Value
					}
				}
				tableprinter.AddField(val)
			case "deployed":
				added := false
				for _, o := range env.Overview {
					if strings.Contains(o.Name, "Deployed") {
						added = true
						check, err := time.Parse("2006-01-02T15:04:05", strings.Split(o.Value, ".")[0])
						if err != nil {
							tableprinter.AddField(o.Value)
						} else {
							tableprinter.AddField(table.RelativeTimeAgo(time.Now(), check))
						}
					}
				}
				if !added {
					tableprinter.AddField("")
				}
			case "created":
				tableprinter.AddField(table.RelativeTimeAgo(time.Now(), env.Environment.Createdat))
			case "password":
				for _, o := range env.Overview {
					if o.Name == "Password" {
						tableprinter.AddField(o.Value)
					}
				}
			}
		}
		tableprinter.EndRow()
	}
	tableprinter.Print()
	o.F.Printer.Exit(0)
}

func (o *Opts) PrintMultiEnvs() {
	if o.Output != "" {
		var toPrint api.EnvironmentsPresenter
		err := toPrint.FromJSON(o.Multi.Environments)
		if err != nil {
			o.F.Printer.Fatal(9, "cannot unmarshal data")
		}
		switch strings.ToLower(o.Output) {
		case "json":
			formatter.PrintJson(o.F, toPrint)
			o.F.Printer.Exit(0)
		case "yaml":
			formatter.PrintYaml(o.F, toPrint)
			o.F.Printer.Exit(0)
		}
	}
	headers := []string{"id", "name", "status", "active", "branch", "created"}
	tableprinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tableprinter.HeaderRow(color, headers...)
	for _, env := range o.Multi.Environments {
		tableprinter.AddField(fmt.Sprint(env.ID))
		tableprinter.AddField(env.Name)
		tableprinter.AddField(env.Status)
		tableprinter.AddField(fmt.Sprint(env.Active))
		tableprinter.AddField(env.GitBranch)
		tableprinter.AddField(table.RelativeTimeAgo(time.Now(), o.Multi.Environments[0].Createdat))
		tableprinter.EndRow()
	}

	tableprinter.Print()
}
