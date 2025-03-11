package get

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/formatter"
	manifestprocessor "github.com/berrybytes/zocli/pkg/utils/manifestProcessor"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

type Opts struct {
	F       *factory.Factory
	ID      string
	Name    string
	List    api.ProjectList
	NewList ProjectList
	Wide    bool
	Output  string
}

type ProjectList struct {
	Projects []models.Project `json:"projects" yaml:"projects"`
}

func NewProjectGetCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	getProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "get",
		Aliases:               []string{"g", "ge"},
		Short:                 "get projects and project details you have access to",
		GroupID:               "basic",
		Long:                  grammar.ProjectGetHelp,
		Run:                   opts.checkArgs,
		Example:               grammar.ProjectGetExample,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(opts.F)
		},
	}

	getProj.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	getProj.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	getProj.Flags().BoolVarP(&opts.Wide, "wide", "w", false, "more columns")
	getProj.Flags().StringVarP(&opts.Output, "out", "o", "", "output in json or yaml")
	return getProj
}

func (o *Opts) checkArgs(_ *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.GetProjectDetail(o.ID)
		return
	}

	if o.Name != "" {
		o.getProjectDetailByName()
	}

	o.GetAllProjects()
	o.printProjects()
	o.F.Printer.Exit(0)
}

func (o *Opts) getProjectDetailByName() {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectByName(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:name>", o.Name)

	res := reqConfig.Request()
	if res == nil {
		return
	}

	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	o.GetProjectDetail(fmt.Sprint(project.ID))
}

func (o *Opts) GetProjectDetail(id string) {
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
	o.List.Projects = append(o.List.Projects, project)
	o.printProjects()
	o.F.Printer.Exit(0)
}

func (o *Opts) GetAllProjects() {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectList(o.F, map[string]interface{}{"headers": header})
	res := reqConfig.Request()
	err := o.List.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}
}

func (o *Opts) printProjects() {
	if o.Output != "" {
		switch strings.ToLower(o.Output) {
		case "json":
			if len(o.List.Projects) == 1 {
				manifest := manifestprocessor.New(o.F)
				manifestProj := manifest.MakeManifest("project", o.List.Projects[0])
				formatter.PrintJson(o.F, manifestProj)
				return
			}
			var manifestYamls []models.Project
			manifest := manifestprocessor.New(o.F)
			for _, one := range o.List.Projects {
				manifestProj := manifest.MakeManifest("project", one)
				proj, _ := manifestProj.(*models.Project)
				manifestYamls = append(manifestYamls, *proj)
			}
			o.NewList = ProjectList{Projects: manifestYamls}
			formatter.PrintJson(o.F, o.NewList)
			return
		case "yaml":
			if len(o.List.Projects) == 1 {
				manifest := manifestprocessor.New(o.F)
				manifestProj := manifest.MakeManifest("project", o.List.Projects[0])
				formatter.PrintYaml(o.F, manifestProj)
				return
			}
			var manifestYamls []models.Project
			manifest := manifestprocessor.New(o.F)
			for _, one := range o.List.Projects {
				manifestProj := manifest.MakeManifest("project", one)
				proj, _ := manifestProj.(*models.Project)
				manifestYamls = append(manifestYamls, *proj)
			}
			o.NewList = ProjectList{Projects: manifestYamls}
			formatter.PrintYaml(o.F, o.NewList)
			return
		default:
			o.F.Printer.Fatal(1, "no such output format available")
		}
	}

	defaultHeaders := []string{"id", "name", "email", "subscription", "cores", "memory(gb)", "active"}
	printHeaders := defaultHeaders

	var wideHeaders []string
	wideHeaders = append(wideHeaders, defaultHeaders...)
	wideHeaders = append(wideHeaders, "description", "region", "logging", "diskspace(gb)", "lb", "monitoring")
	if o.Wide {
		printHeaders = wideHeaders
	}

	tablePrinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tablePrinter.HeaderRow(color, printHeaders...)

	for _, project := range o.List.Projects {
		for _, header := range printHeaders {
			switch header {
			case "id":
				tablePrinter.AddField(fmt.Sprintf("%d", project.ID))
			case "name":
				tablePrinter.AddField(project.Name)
			case "description":
				tablePrinter.AddField(project.Description)
			case "email":
				tablePrinter.AddField(project.User.Email)
			case "subscription":
				tablePrinter.AddField(project.Subscription.Name)
			case "cores":
				tablePrinter.AddField(fmt.Sprintf("%d", (project.Subscription.Cores / 1024)))
			case "memory(gb)":
				tablePrinter.AddField(fmt.Sprintf("%d", project.Subscription.Memory/1024))
			case "active":
				tablePrinter.AddField(fmt.Sprintf("%v", project.Active))
			case "region":
				tablePrinter.AddField(project.Region)
			case "logging":
				tablePrinter.AddField(project.Logging)
			case "diskspace(gb)":
				tablePrinter.AddField(strconv.Itoa(project.Subscription.DiskSpace / 1024))
			case "lb":
				tablePrinter.AddField(fmt.Sprintf("%v", project.DedicatedLb))
			case "monitoring":
				tablePrinter.AddField(fmt.Sprintf("%v", project.Monitoring))
			}
		}
		tablePrinter.EndRow()
	}
	_ = tablePrinter.Print()
}
