package environment

import (
	"github.com/spf13/cobra"
)

type EnvironmentInterface interface {
	GetRunner(*cobra.Command, []string)
	StopRunner(*cobra.Command, []string)
	StartRunner(*cobra.Command, []string)
	DeleteRunner(*cobra.Command, []string)
	RenameRunner(*cobra.Command, []string)
	OverviewRunner(*cobra.Command, []string)

	FetchEnvironment(id string)
	PrintSingleEnvTable()
	PrintMultiEnvs()
	GetEnvironmentOverview(string)
	PrintOverviewTable()
	StopEnvByID()
	StartEnvByID()
	DeleteEnvByID()
	AskDeleteConfirmation()
	RenameEnv()
	AskChanges()
}

func NewEnvironmentInterface(o *Opts) EnvironmentInterface {
	return o
}
