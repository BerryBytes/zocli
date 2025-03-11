package overview

import (
	"fmt"
	"strings"
	"sync"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

type Opts struct {
	F           *factory.Factory
	ID          string
	Name        string
	Wide        bool
	Output      string
	Project     *api.Project
	Resource    *api.Resource
	ProjectList *api.ProjectList
	ResourceMap map[int]api.Resource
	ProjectMap  map[int]api.Project
}

func NewProjectOverviewCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	overviewProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "overview",
		Aliases:               []string{"o", "over", "ove", "overv"},
		Short:                 "get projects memory, cores and other details you have access to",
		Long:                  grammar.ProjectOverviewHelp,
		GroupID:               "basic",
		Run:                   opts.checkArgs,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
			middleware.LoggedIn(opts.F)
		},
	}

	overviewProj.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	overviewProj.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	return overviewProj
}

func (o *Opts) checkArgs(_ *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.getProjectDetail(o.ID)
		if o.ProjectMap == nil {
			o.ProjectMap = make(map[int]api.Project)
		}
		o.ProjectMap[o.Project.ID] = *o.Project
		o.printTable()
		return
	}

	if o.Name != "" {
		o.getProjectDetailByName()
		if o.ProjectMap == nil {
			o.ProjectMap = make(map[int]api.Project)
		}
		o.ProjectMap[o.Project.ID] = *o.Project
		o.printTable()
		return
	}

	projects := o.GetAllProjects()
	var wg sync.WaitGroup
	for _, proj := range projects.Projects {
		wg.Add(1)
		go func(id int, proj api.Project) {
			defer wg.Done()
			o.fetchResource(id)
			if o.ProjectMap == nil {
				o.ProjectMap = make(map[int]api.Project)
			}
			o.ProjectMap[proj.ID] = proj
		}(proj.ID, proj)
	}
	wg.Wait()
	o.printTable()
}

func (o *Opts) GetAllProjects() *api.ProjectList {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectList(o.F, map[string]interface{}{"headers": header})
	res := reqConfig.Request()
	var projects api.ProjectList
	err := projects.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
		return nil
	}
	return &projects
}

func (o *Opts) getProjectDetail(id string) {
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

	o.Project = &project
	o.fetchResource(project.ID)
}

// getProjectDetailByName
//
// this function fetches the details related to project because of which the
// subscription details can be found, so while showing the overview details,
// we will need that.
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
	o.getProjectDetail(fmt.Sprint(project.ID))
}

func (o *Opts) fetchResource(id int) {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectResource(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(id))

	res := reqConfig.Request()
	if res == nil {
		return
	}

	var resource api.Resource
	err := resource.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err.Error())
	}
	if o.ResourceMap == nil {
		o.ResourceMap = make(map[int]api.Resource)
	}
	o.ResourceMap[id] = resource
	o.Resource = &resource
}

func (o *Opts) printTable() {
	datatransfer := fmt.Sprintf("%.2f", (o.Resource.DataTransfer.DataTransfer.Receive+o.Resource.DataTransfer.DataTransfer.Transmit)/1024)
	tablePrinter := table.New(o.F, 0)

	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tablePrinter.HeaderRow(color, "id", "projectname", "apps", "memory(gb)", "core(millicores)", "storage(gb)", "data transfer(gb)")

	for id, resource := range o.ResourceMap {
		tablePrinter.AddField(fmt.Sprint(o.ProjectMap[id].ID))
		tablePrinter.AddField(o.ProjectMap[id].Name)
		tablePrinter.AddField(fmt.Sprint(resource.Apps) + "/" + fmt.Sprint(o.ProjectMap[id].Subscription.Apps))
		tablePrinter.AddField(fmt.Sprint(resource.Memory/1024) + " / " + fmt.Sprint(o.ProjectMap[id].Subscription.Memory/1024))
		tablePrinter.AddField(fmt.Sprint(resource.Core) + " / " + fmt.Sprint(o.ProjectMap[id].Subscription.Cores))
		tablePrinter.AddField(fmt.Sprint(resource.Disk/1024) + " / " + fmt.Sprint(o.ProjectMap[id].Subscription.DiskSpace/1024))
		tablePrinter.AddField(datatransfer + " / " + fmt.Sprint(o.ProjectMap[id].Subscription.DataTransfer/1024))
		tablePrinter.EndRow()
	}
	_ = tablePrinter.Print()
}
