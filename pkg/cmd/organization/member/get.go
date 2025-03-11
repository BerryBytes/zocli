package member

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/formatter"
	manifestprocessor "github.com/berrybytes/zocli/pkg/utils/manifestProcessor"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/spf13/cobra"
)

func NewOrganizationMembersGetCommand(o *Opts) *cobra.Command {
	get := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g", "ge", "list", "lis", "retrieve"},
		Short:   "organizations members details",
		Long:    grammar.OrganizationMemberGetHelp,
		Example: grammar.OrganizationMemberGetExample,
		GroupID: "basic",
		Run:     o.I.GetRunner,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	get.Flags().BoolVarP(&o.Wide, "wide", "w", false, "more columns")
	get.Flags().StringVarP(&o.Output, "out", "o", "", "output in JSON/YAML")
	return get
}

// GetRunner
//
// responsible to fetch all the Organization Member List that the user
// has access to.
func (o *Opts) GetRunner(cmd *cobra.Command, _ []string) {
	if o.F.Config.ActiveContext.OrganizationID == 0 {
		fmt.Print("no need to check for members on default organization")
		return
	}
	o.I.GetOrganizationMembers()
	o.I.PrintOrganizationMembersTable()
	o.F.Printer.Exit(0)
}

func (o *Opts) GetOrganizationMembers() {
	headers := o.F.GetAuth()
	reqConfig := defaults.GetOrganization(o.F, map[string]interface{}{"headers": headers})
	res := reqConfig.Request()
	var org api.Organization
	err := utils.ConvertType(res.Data, &org)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal")
	}
	o.Org.Members = org.Members
}

func (o *Opts) GetSingleOrganizationMember(email string) {
	o.GetOrganizationMembers()
	for i, member := range o.Org.Members {
		if member.Owner.Email == email {

			o.Org.Members = append(o.Org.Members, o.Org.Members[i])
			fmt.Print("Entered the organization member email to proceed on : ", o.Org.Members[i])
		}
	}
}

func (o *Opts) PrintOrganizationMembersTable() {
	if o.Output != "" {
		switch strings.ToLower(o.Output) {
		case "json":
			o.makeManifest()
			if len(o.Org.Members) == 1 {
				formatter.PrintJson(o.F, o.Org.Members[0])
				return
			}
			formatter.PrintJson(o.F, o.List)
			return
		case "yaml":
			o.makeManifest()
			if len(o.Org.Members) == 1 {
				formatter.PrintYaml(o.F, o.Org.Members[0])
				return
			}
			formatter.PrintYaml(o.F, o.List)
			return
		default:
			o.F.Printer.Fatal(1, "no such output format available")
		}
	}

	tableprinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	defaultHeaders := []string{"name", "email", "role"}
	tableprinter.HeaderRow(color, defaultHeaders...)

	for _, member := range o.Org.Members {
		for _, header := range defaultHeaders {
			switch header {
			case "name":
				tableprinter.AddField(member.Owner.FirstName + " " + member.Owner.LastName)
			case "email":
				tableprinter.AddField(member.Owner.Email)
			case "role":
				tableprinter.AddField(models.GetRole(member.UserRole))
			}
		}
		tableprinter.EndRow()
	}
	_ = tableprinter.Print()
}

func (o *Opts) makeManifest() {
	manifest := manifestprocessor.New(o.F)
	for _, one := range o.Org.Members {
		manifestOrg := manifest.MakeManifest("member", one)
		org, _ := manifestOrg.(*models.OrganizationMember)
		o.List.Members = append(o.List.Members, *org)
	}

}
