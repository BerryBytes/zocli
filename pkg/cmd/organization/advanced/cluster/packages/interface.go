package packages

import "github.com/spf13/cobra"

type Interface interface {
	StatusRunner(*cobra.Command, []string)
	InstallRunner(*cobra.Command, []string)
	UnInstallRunner(*cobra.Command, []string)
	StatusAll()
}

func NewInterface(o *Opts) Interface {
	return o
}
