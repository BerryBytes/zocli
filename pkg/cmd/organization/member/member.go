package member

import (
	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/spf13/cobra"
)

type Opts struct {
	F       *factory.Factory
	I       Interface
	Org     *api.OrganizationMembersList
	Wide    bool
	Output  string
	List    MemberList
	NewList MemberList
}

type MemberList struct {
	Members []models.OrganizationMember `json:"members" yaml:"members"`
}

func NewOrganizationMembersCommand(f *factory.Factory) *cobra.Command {
	o := Opts{F: f}
	o.I = NewInterface(&o)
	o.Org = &api.OrganizationMembersList{}
	org := &cobra.Command{
		Use:     "members",
		Aliases: []string{"m", "me", "mem", "memb", "membe"},
		Short:   "organization members commands",
		Long:    grammar.OrganizationMembersHelp,
		GroupID: "basic",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	org.AddGroup(&cobra.Group{
		ID:    "basic",
		Title: "Basic commands to maintain organization member",
	})

	org.AddGroup(&cobra.Group{
		ID:    "context",
		Title: "Maintain contexts or defaults",
	})

	org.AddCommand(NewOrganizationMembersGetCommand(&o))
	org.AddCommand(NewOrganizationMemberDeleteCommand(&o))
	return org
}
