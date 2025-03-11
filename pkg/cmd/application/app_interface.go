package application

import (
	"github.com/berrybytes/zocli/api"
	"github.com/spf13/cobra"
)

type Interface interface {
	GetRunner(*cobra.Command, []string)
	GetApps()
	GetSingleApplication(string)
	GetProjectDetailByName() *api.Project
	PrintApps()
	GetProjectDetail(id string) *api.Project
	RenameRunner(*cobra.Command, []string)
	RenameApplication()
	DeleteRunner(*cobra.Command, []string)
	DeleteApplication()
	AskDeleteConfirmation()
	GetDefaultApplicationRunner(*cobra.Command, []string)
	RemoveDefaultApplicationRunner(*cobra.Command, []string)
	SetDefaultRunner(*cobra.Command, []string)
}

func NewAppInterface(o *Opts) Interface {
	return o
}
