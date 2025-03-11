package globalparser

import (
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/spf13/cobra"
)

type GlobalFlags struct {
	NoInteractive bool
	Verbose       bool
	Quiet         bool
}

func ParseGlobal(cmd *cobra.Command, _ []string, f *factory.Factory) *GlobalFlags {
	globalFlags := GlobalFlags{}
	interactive, err := cmd.Flags().GetBool("no-interactive")
	if err != nil {
		f.Printer.Fatal(6, err)
	}
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		f.Printer.Fatal(6, err)
	}
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		f.Printer.Fatal(6, err)
	}

	globalFlags.NoInteractive = interactive
	globalFlags.Verbose = verbose
	globalFlags.Quiet = quiet

	return &globalFlags
}
