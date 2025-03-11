package organization

import (
	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/cmd/organization/advanced/cluster"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/spf13/cobra"
)

type Opts struct {
	F      *factory.Factory
	I      Interface
	Org    *api.OrganizationList
	Wide   bool
	Output string
	List   OrganizationList

	OrgID string
}

type OrganizationList struct {
	Organizations []models.Organization `json:"organizations" yaml:"organizations"`
}

func NewOrganizationCommand(f *factory.Factory) *cobra.Command {
	o := Opts{F: f}
	o.I = NewInterface(&o)
	o.Org = &api.OrganizationList{}
	org := &cobra.Command{
		Use:     "organization",
		Aliases: []string{"o", "or", "org", "orga", "organ", "organi", "organiz", "organiza", "organizat", "organizati", "organizatio"},
		Short:   "organization commands",
		Long:    grammar.OrganizationHelp,
		GroupID: "basic",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	org.AddGroup(&cobra.Group{
		ID:    "basic",
		Title: "Basic commands to maintain organization",
	})

	org.AddGroup(&cobra.Group{
		ID:    "context",
		Title: "Maintain contexts or defaults",
	})

	org.AddGroup(&cobra.Group{
		ID:    "admin",
		Title: "commands which admin can only use",
	})

	org.AddCommand(NewOrganizationGetCommand(&o))
	org.AddCommand(NewOrganizationUseDefaultCommand(&o))
	org.AddCommand(NewOrganizationDeleteCommand(&o))
	org.AddCommand(cluster.NewClusterCommand(o.F))
	return org
}
