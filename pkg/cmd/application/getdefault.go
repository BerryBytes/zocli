package application

import (
	"fmt"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewApplicationGetDefaultCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	getDefaultApplication := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "getdefault",
		Aliases:               []string{"gedef", "getdef", "getde", "retrieve", "default", "def"},
		Short:                 "get the default application currently used",
		Long:                  grammar.ApplicationGetDefaultHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.GetDefaultApplicationRunner,
		GroupID:               "context",
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
			validation.CheckValidID(f, opts.ID)
		},
	}

	return getDefaultApplication
}

func (o *Opts) GetDefaultApplicationRunner(_ *cobra.Command, _ []string) {
	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultApplication != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
		o.F.Printer.Printf(" The default application on active context is of ID %d", o.F.Config.ActiveContext.DefaultApplication)
		o.ProjectID = fmt.Sprint(o.F.Config.ActiveContext.DefaultApplication)
		o.GetSingleApplication(fmt.Sprint(o.F.Config.ActiveContext.DefaultApplication))
		return
	}
	o.F.IO.ColorScheme().FailureIcon()
	o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
	o.F.Printer.Printf(" No any active application is found")
}
