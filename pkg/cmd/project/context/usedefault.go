package context

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/cmd/project/get"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/context"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

type Opts struct {
	F         *factory.Factory
	ID        string
	Temporary bool
}

func NewProjectUseDefaultCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	defaultProject := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "use",
		Aliases:               []string{"u", "default", "def", "us", "defa"},
		Short:                 "set this as a default project",
		Long:                  grammar.ProjectUseDefaultHelp,
		Example:               grammar.ProjectUseDefaultExample,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.setDefaultRunner,
		GroupID:               "context",
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID = args[0]
			}
			middleware.LoggedIn(opts.F)
			validation.CheckValidID(f, opts.ID)
		},
	}

	defaultProject.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	defaultProject.Flags().BoolVarP(&opts.Temporary, "temp", "t", false, "use for temporary bash session")
	return defaultProject
}

func (o *Opts) setDefaultRunner(cmd *cobra.Command, _ []string) {
	var wg sync.WaitGroup
	var projects api.ProjectList
	wg.Add(1)
	go func() {
		defer wg.Done()

		projectOpts := get.Opts{F: o.F}
		projectOpts.GetAllProjects()
		projects = projectOpts.List
	}()

	var id int
	if o.ID == "" {
		o.F.Printer.Print("Enter ID : ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon(), " Invalid input.\n")
			return
		}
		validation.CheckValidID(o.F, input)
		o.ID = input
	}
	wg.Wait()
	id, _ = strconv.Atoi(o.ID)

	found := false
	for _, project := range projects.Projects {
		if id == project.ID {
			found = true
		}
	}

	if !found {
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon(), " no such Project Found.\n")
		return
	}

	if o.Temporary {
		// export variable in bash
		o.F.Printer.Println("Feat: Might be coming on near future.")
		return
	}

	// save the default as persistent value
	context.SaveContextChanges(o.F, id, "project", "")
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully changed the default project to "+o.ID+"\n")
}
