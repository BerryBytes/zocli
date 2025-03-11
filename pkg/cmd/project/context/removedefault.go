package context

import (
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/context"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectDeleteDefaultCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	getDefaultProject := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "removedefault",
		Aliases:               []string{"remove", "rem", "rede", "remde", "redefault", "remdef"},
		Short:                 "remove the default project from the configuration",
		Long:                  grammar.ProjectGetDefaultHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.removeDefaultProjectRunner,
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

func (o *Opts) removeDefaultProjectRunner(_ *cobra.Command, _ []string) {
	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		context.SaveContextChanges(o.F, 0, "project", "")
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
		o.F.Printer.Printf(" Removed the default project information from the configuration")
		return
	}
	o.F.IO.ColorScheme().FailureIcon()
	o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
	o.F.Printer.Printf(" No any active project is found")
}
