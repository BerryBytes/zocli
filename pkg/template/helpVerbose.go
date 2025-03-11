package template

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func HelpVerbose(cmd *cobra.Command, _ []string) {
	fmt.Println(cmd.Long)
	fmt.Println("\nUsage:")
	fmt.Printf("\t%s [command] [flags]\n", cmd.CommandPath())
	fmt.Println("\nFlags:")
	err := printFlagSet(cmd.PersistentFlags())
	if err != nil {
		return
	}
	fmt.Println("\nNot recommended:")
	//err = printFlagSet(newToken)
	//if err != nil {
	//	return
	//}
	fmt.Println("\nGlobal Flags:")
	err = printFlagSet(cmd.InheritedFlags())
	if err != nil {
		return
	}
	fmt.Println("\nExample:")
	fmt.Println(cmd.Example)
}
func printFlagSet(flagSet *pflag.FlagSet) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			// Skip hidden flags (not recommended)
			return
		}
		_, err := fmt.Fprintf(w, "  -%s, --%s\t%s\t(default: %s)\n", flag.Shorthand, flag.Name, flag.Usage, flag.DefValue)
		if err != nil {
			return
		}
	})
	err := w.Flush()
	if err != nil {
		return err
	}
	return nil
}
