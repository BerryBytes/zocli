package organization

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/context"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/spf13/cobra"
)

func NewOrganizationDeleteCommand(o *Opts) *cobra.Command {
	del := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d", "del", "de", "dele"},
		Short:   "organizations details",
		Long:    grammar.OrganizationDeleteHelp,
		GroupID: "basic",
		Run:     o.I.DeleteRunner,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	return del
}

func (o *Opts) DeleteRunner(cmd *cobra.Command, _ []string) {
	o.GetSingleOrganization()
	o.PrintOrganizationTable()
	org := o.Org.Organizations[0]
	o.confirmAndDelete(&org)
	_ = cmd.Help()
}

func (o *Opts) confirmAndDelete(org *api.Organization) {
	utils.ConfirmIfToProceed("Do you sure want to delete the organization ?", o.F)
	fmt.Print("Enter the organization name to proceed on : ")
	text, err := o.F.Term.ReadInput()
	if err != nil {
		o.F.Printer.Fatal(5, err)
	}
	orgName := strings.TrimSpace(strings.Trim(text, "\n"))
	if orgName != org.Name {
		o.F.Printer.Fatal(1, "invalid org name")
	}

	header := o.F.GetAuth()
	reqConfig := defaults.DeleteOrganization(o.F, map[string]interface{}{"headers": header})
	res := reqConfig.Request()
	if res == nil {
		return
	}

	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully deleted '"+org.Name+"'")

	o.OrgID = "0"
	orgSwitch := o.switchOrg()
	o.SaveTokenChanges(&orgSwitch)
	o.F = context.SetActiveFalse(o.F)
	o.SaveContextChanges(0, &orgSwitch)
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Printf(" Successfully set organization with id '" + o.OrgID + "' as default")
	o.F.Printer.Exit(0)
}
