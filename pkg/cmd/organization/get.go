package organization

import (
	"fmt"
	"strings"
	"time"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/formatter"
	manifestprocessor "github.com/berrybytes/zocli/pkg/utils/manifestProcessor"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/spf13/cobra"
)

func NewOrganizationGetCommand(o *Opts) *cobra.Command {
	get := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g", "ge", "list", "lis", "retrieve"},
		Short:   "organizations details",
		Long:    grammar.OrganizationGetHelp,
		Example: grammar.OrganizationGetExample,
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
// responsiuble to fetch all the Organization List that the user
// has access to.
func (o *Opts) GetRunner(cmd *cobra.Command, _ []string) {
	o.I.GetOrganizations()
	o.I.PrintOrganizationTable()
	o.F.Printer.Exit(0)
}

func (o *Opts) PrintOrganizationTable() {
	if o.Output != "" {
		switch strings.ToLower(o.Output) {
		case "json":
			o.makeManifest()
			if len(o.Org.Organizations) == 1 {
				formatter.PrintJson(o.F, o.List.Organizations[0])
				return
			}
			formatter.PrintJson(o.F, o.List)
			return
		case "yaml":
			o.makeManifest()
			if len(o.Org.Organizations) == 1 {
				formatter.PrintYaml(o.F, o.List.Organizations[0])
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
	defaultHeaders := []string{"id", "name", "domain", "owner", "createdat"}
	if o.Wide {
		defaultHeaders = append(defaultHeaders, []string{"plan", "image", "total members"}...)
	}
	tableprinter.HeaderRow(color, defaultHeaders...)

	for _, org := range o.Org.Organizations {
		for _, header := range defaultHeaders {
			switch header {
			case "id":
				tableprinter.AddField(fmt.Sprint(org.ID))
			case "name":
				tableprinter.AddField(org.Name)
			case "domain":
				tableprinter.AddField(org.Domain)
			case "owner":
				tableprinter.AddField(org.User.Email)
			case "createdat":
				tableprinter.AddField(table.RelativeTimeAgo(time.Now(), org.Createdat))
			case "plan":
				if org.OrganizationPlan != nil {
					tableprinter.AddField(org.OrganizationPlan.Name)
				} else {
					tableprinter.AddField("")
				}
			case "image":
				tableprinter.AddField(org.Image)
			case "total members":
				tableprinter.AddField(fmt.Sprint(len(org.Members)))
			}
		}
		tableprinter.EndRow()
	}
	_ = tableprinter.Print()
}

func (o *Opts) GetOrganizations() {
	headers := o.F.GetAuth()
	reqConfig := defaults.GetOrganizations(o.F, map[string]interface{}{"headers": headers})
	res := reqConfig.Request()
	var org api.OrganizationList
	err := org.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal")
	}
	o.Org = &org
}

// NOTE:
// below portion code is yet to be tested, as till now there is no any organization on the context
func (o *Opts) GetSingleOrganization() {
	headers := o.F.GetAuth()
	reqConfig := defaults.GetOrganization(o.F, map[string]interface{}{"headers": headers})
	res := reqConfig.Request()
	var org api.Organization
	err := org.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal")
	}
	o.Org.Organizations = append(o.Org.Organizations, org)
}

func (o *Opts) makeManifest() {
	manifest := manifestprocessor.New(o.F)
	for _, one := range o.Org.Organizations {
		manifestOrg := manifest.MakeManifest("organization", one)
		org, _ := manifestOrg.(*models.Organization)
		o.List.Organizations = append(o.List.Organizations, *org)
	}

}
