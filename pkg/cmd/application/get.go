package application

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/formatter"
	manifestprocessor "github.com/berrybytes/zocli/pkg/utils/manifestProcessor"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewApplicationGetCommand(o *Opts) *cobra.Command {
	app := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g", "ge", "list", "lis", "retrieve"},
		Short:   "get apps on a project",
		Long:    grammar.ApplicationGetHelp,
		GroupID: "basic",
		Example: grammar.ApplicationGetExample,
		Run:     o.I.GetRunner,
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
	app.Flags().StringVarP(&o.ProjectName, "pname", "n", "", "project name")
	app.Flags().StringVarP(&o.ApplicationID, "id", "i", "", "application id")
	app.Flags().StringVarP(&o.Output, "out", "o", "", "output in json or yaml")
	return app
}

func (o *Opts) GetRunner(cmd *cobra.Command, _ []string) {
	if o.ApplicationID != "" {
		o.GetSingleApplication(o.ApplicationID)
		return
	}

	if o.ProjectID != "" && o.ProjectName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ProjectID != "" {
		validation.CheckValidID(o.F, o.ProjectID)
		o.I.GetApps()
		o.I.PrintApps()
		o.F.Printer.Exit(0)
		return
	}

	if o.ProjectName != "" {
		proj := o.I.GetProjectDetailByName()
		o.ProjectID = fmt.Sprint(proj.ID)
		o.I.GetApps()
		o.I.PrintApps()
		o.F.Printer.Exit(0)
		return
	}

	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.ProjectID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
		o.I.GetApps()
		o.I.PrintApps()
		o.F.Printer.Exit(0)
		return
	}

	_ = cmd.Help()
}

func (o *Opts) GetApps() {
	headers := o.F.GetAuth()
	reqConfig := defaults.GetAppsList(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ProjectID)

	res := reqConfig.Request()
	var apps api.ApplicationList
	err := apps.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	o.List = apps
}

func (o *Opts) GetSingleApplication(appID string) {
	headers := o.F.GetAuth()
	reqConfig := defaults.GetSingleApplication(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", appID)
	res := reqConfig.Request()
	var apps api.Application
	err := apps.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "unable to Marshal")
	}
	o.List.Applications = append(o.List.Applications, apps)
	o.PrintApps()
}

func (o *Opts) PrintApps() {
	if o.Output != "" {
		// trim the format as there are many unnecessary data on the struct
		final := api.ApplicationPresenterList{}
		for _, app := range o.List.Applications {
			var presentedApp api.ApplicationPresenter
			err := presentedApp.FromJSON(app)
			if err != nil {
				o.F.Printer.Fatal(9, "cannot Unmarshal")
				return
			}

			final.Applications = append(final.Applications, presentedApp)
		}
		switch strings.ToLower(o.Output) {
		case "json":
			if len(o.List.Applications) == 1 {
				manifest := manifestprocessor.New(o.F)
				manifestApp := manifest.MakeManifest("application", o.List.Applications[0])
				formatter.PrintJson(o.F, manifestApp)
				return
			}
			var manifestYamls []models.App
			manifest := manifestprocessor.New(o.F)
			for _, one := range o.List.Applications {
				manifestApp := manifest.MakeManifest("application", one)
				app, _ := manifestApp.(*models.App)
				manifestYamls = append(manifestYamls, *app)
			}
			o.NewList = ApptList{Apps: manifestYamls}
			formatter.PrintJson(o.F, o.NewList)
			return
		case "yaml":
			if len(o.List.Applications) == 1 {
				manifest := manifestprocessor.New(o.F)
				manifestApp := manifest.MakeManifest("application", o.List.Applications[0])
				formatter.PrintYaml(o.F, manifestApp)
				return
			}
			var manifestYamls []models.App
			manifest := manifestprocessor.New(o.F)
			for _, one := range o.List.Applications {
				manifestApp := manifest.MakeManifest("application", one)
				app, _ := manifestApp.(*models.App)
				manifestYamls = append(manifestYamls, *app)
			}
			o.NewList = ApptList{Apps: manifestYamls}
			formatter.PrintYaml(o.F, o.NewList)
			return
		default:
			o.F.Printer.Fatal(1, "no such output format available")
		}
	}

	tablePrinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")

	tableHeaders := []string{"id", "name", "servicetype", "owner"}
	dockerHeadersAdded := false
	gitHeadersAdded := false
	for _, app := range o.List.Applications {
		if app.ServiceType == api.GIT && !gitHeadersAdded {
			tableHeaders = append(tableHeaders, "giturl", "repository")
			gitHeadersAdded = true
		}
		if app.ServiceType == api.DOCKER && !dockerHeadersAdded {
			tableHeaders = append(tableHeaders, "imagerepo", "imageurl")
			dockerHeadersAdded = true
		}
	}

	tableHeaders = append(tableHeaders, "active")
	// determine the headers
	tablePrinter.HeaderRow(color, tableHeaders...)

	for _, app := range o.List.Applications {
		for _, header := range tableHeaders {
			switch header {
			case "id":
				tablePrinter.AddField(fmt.Sprint(app.Id))
			case "name":
				tablePrinter.AddField(app.Name)
			case "servicetype":
				tablePrinter.AddField(api.GetServiceType(app.ServiceType))
			case "active":
				tablePrinter.AddField(fmt.Sprint(app.Active))
			case "owner":
				tablePrinter.AddField(app.Owner.FirstName + " " + app.Owner.LastName)
			case "giturl":
				tablePrinter.AddField(app.GitRepoUrl)
			case "repository":
				tablePrinter.AddField(app.GitRepository)
			case "imagerepo":
				tablePrinter.AddField(app.ImageRepo)
			case "imageurl":
				tablePrinter.AddField(app.ImageUrl)
			}
		}
		tablePrinter.EndRow()
	}
	tablePrinter.Print()
}
