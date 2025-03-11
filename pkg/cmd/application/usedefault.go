package application

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/context"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewApplicationUseDefaultCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	defaultApplication := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "use",
		Aliases:               []string{"u", "default", "def", "us", "defa"},
		Short:                 "set this as a default application",
		Long:                  grammar.ApplicationUseDefaultHelp,
		Example:               grammar.ApplicationUseDefaultExample,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.SetDefaultRunner,
		GroupID:               "context",
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID = args[0]
			}
			middleware.LoggedIn(opts.F)
			validation.CheckValidID(f, opts.ID)
		},
	}

	defaultApplication.Flags().StringVarP(&opts.ID, "id", "i", "", "application id")
	defaultApplication.Flags().StringVarP(&opts.ProjectID, "pid", "p", "", "project id")
	defaultApplication.Flags().StringVarP(&opts.ProjectName, "pname", "n", "", "project name")
	defaultApplication.Flags().BoolVarP(&opts.Temporary, "temp", "t", false, "use for temporary bash session")
	return defaultApplication
}

func (o *Opts) SetDefaultRunner(cmd *cobra.Command, _ []string) {
	if o.ProjectID != "" && o.ProjectName != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ProjectID != "" {
		validation.CheckValidID(o.F, o.ProjectID)
	}

	if o.ProjectName != "" {
		proj := o.I.GetProjectDetailByName()
		o.ProjectID = fmt.Sprint(proj.ID)
	}

	if o.ProjectID == "" && o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.ProjectID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
	}

	if o.ProjectID == "" && o.ProjectName == "" {
		_ = cmd.Help()
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		o.GetApps()
	}()

	var id int
	if o.ID == "" {
		o.F.Printer.Print("Enter ID : ")
		var input string
		_, _ = fmt.Scanln(&input)

		validation.CheckValidID(o.F, input)
		o.ID = input
	}
	wg.Wait()
	id, _ = strconv.Atoi(o.ID)

	found := false
	for _, app := range o.List.Applications {
		if id == app.Id {
			found = true
		}
	}

	if !found {
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon(), " no such Application Found.\n")
		return
	}

	if o.Temporary {
		// export variable in bash
		o.F.Printer.Println("Feat: Might be coming on near future.")
		return
	}

	// save the default as persistent value
	context.SaveContextChanges(o.F, id, "application", "")
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully changed the default Application to "+o.ID+"\n")
}
