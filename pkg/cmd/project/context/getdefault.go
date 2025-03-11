package context

import (
	"fmt"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/cmd/project/get"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

// NewProjectGetDefaultCommand
//
// this command is responsible for printing the default project
// that is being used in this active orgranization
func NewProjectGetDefaultCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	getDefaultProject := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "getdefault",
		Aliases:               []string{"gedef", "getdef", "getde", "retrieve", "default", "def"},
		Short:                 "get the default project currently used",
		Long:                  grammar.ProjectGetDefaultHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.getDefaultProjectRunner,
		GroupID:               "context",
		PreRun: func(_ *cobra.Command, args []string) {
			if len(args) != 0 {
				opts.ID = args[0]
			}
			middleware.LoggedIn(opts.F)
			validation.CheckValidID(f, opts.ID)
		},
	}

	return getDefaultProject
}

func (o *Opts) getDefaultProjectRunner(_ *cobra.Command, _ []string) {
	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
		o.F.Printer.Printf(" The default project on active context is of ID %d", o.F.Config.ActiveContext.DefaultProject)
		getProjectOpts := get.Opts{F: o.F}
		getProjectOpts.GetProjectDetail(fmt.Sprint(o.F.Config.ActiveContext.DefaultProject))
		return
	}
	o.F.IO.ColorScheme().FailureIcon()
	o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
	o.F.Printer.Printf(" No any active project is found")
}
