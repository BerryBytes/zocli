package apply

import (
	"github.com/spf13/cobra"
)

type Interface interface {
	ManifestRunner(*cobra.Command, []string)
	MatchModel(interface{}, string)
}

func NewInterface(o *Opts) Interface {
	return o
}
