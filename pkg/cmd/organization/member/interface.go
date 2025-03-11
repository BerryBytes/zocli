package member

import "github.com/spf13/cobra"

type Interface interface {
	GetRunner(*cobra.Command, []string)
	DeleteRunner(*cobra.Command, []string)
	GetOrganizationMembers()
	GetSingleOrganizationMember(email string)
	PrintOrganizationMembersTable()
}

func NewInterface(o *Opts) Interface {
	return o
}
