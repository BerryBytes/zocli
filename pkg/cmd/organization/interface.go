package organization

import "github.com/spf13/cobra"

type Interface interface {
	GetRunner(*cobra.Command, []string)
	DeleteRunner(*cobra.Command, []string)
	UseDefaultRunner(*cobra.Command, []string)
	GetOrganizations()
	GetSingleOrganization()
	PrintOrganizationTable()
}

func NewInterface(o *Opts) Interface {
	return o
}
