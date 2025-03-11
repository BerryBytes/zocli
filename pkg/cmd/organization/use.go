package organization

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/cmd/auth/login"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/context"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func NewOrganizationUseDefaultCommand(o *Opts) *cobra.Command {
	useDef := &cobra.Command{
		Use:                   "use",
		Aliases:               []string{"u", "default", "def", "us", "defa"},
		Short:                 "set default organization",
		Long:                  grammar.OrganizationUseDefaultHelp,
		GroupID:               "context",
		Run:                   o.I.UseDefaultRunner,
		DisableFlagsInUseLine: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				o.OrgID, _ = validation.GetIDOrName(args[0])
				validation.CheckValidID(o.F, o.OrgID)
			}
			middleware.LoggedIn(o.F)
		},
	}

	useDef.Flags().StringVarP(&o.OrgID, "id", "i", "", "organization id")
	return useDef
}

func (o *Opts) UseDefaultRunner(cmd *cobra.Command, _ []string) {
	id, _ := strconv.Atoi(o.OrgID)
	if o.OrgID != "0" {
		o.GetOrganizations()
		found := false
		for _, org := range o.Org.Organizations {
			if org.ID == id {
				found = true
				break
			}
		}
		if !found {
			o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
			o.F.Printer.Fatal(1, " no such organization found")
		}
	}
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Printf(" Valid organization found")

	orgSwitch := o.switchOrg()
	o.SaveTokenChanges(&orgSwitch)
	o.F = context.SetActiveFalse(o.F)
	o.SaveContextChanges(id, &orgSwitch)
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Printf(" Successfully set organization with id '" + o.OrgID + "' as default")
	o.F.Printer.Exit(0)
}

// saveTokenChanges
//
// this function opens the token.yaml file from the config folder and then saves the new token to the file
func (o *Opts) SaveTokenChanges(orgSwitch *api.OrganizationSwitch) {
	var saveConfig login.SaveConfig
	file, err := os.ReadFile(o.F.Config.ConfigFolder + o.F.Config.AuthFile)
	if err != nil {
		o.F.Printer.Fatal(10, "cannot access config folder")
	}
	err = yaml.Unmarshal(file, &saveConfig)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal")
	}
	saveConfig.AuthToken = orgSwitch.Token
	save, err := yaml.Marshal(saveConfig)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot marshal")
	}
	changes, err := os.Create(o.F.Config.ConfigFolder + o.F.Config.AuthFile)
	if err != nil {
		o.F.Printer.Fatal(10, "cannot create the config file")
		return
	}
	if _, err := changes.Write(save); err != nil {
		o.F.Printer.Fatal(10, "cannot save changes")
	}

}

func (o *Opts) SaveContextChanges(id int, orgSwitch *api.OrganizationSwitch) {
	exists := context.CheckIfOrgInContext(id, o.F)
	if exists {
		var err error
		err, o.F = context.SetActive(id, o.F)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	} else {
		active := config.Context{OrganizationID: id, Active: true, OrganizationName: orgSwitch.Organization.Name}
		if id == 0 {
			active.OrganizationName = "Default"
		}
		o.F.Config.Contexts = append(o.F.Config.Contexts, &active)
		o.F.Config.ActiveContext = &active
	}

	contexts, err := yaml.Marshal(o.F.Config.Contexts)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot marshal")
	}
	contextsChanges, err := os.Create(o.F.Config.ConfigFolder + o.F.Config.ContextFile)
	if err != nil {
		o.F.Printer.Fatal(10, "cannot create the config file")
		return
	}
	if _, err := contextsChanges.Write(contexts); err != nil {
		o.F.Printer.Fatal(10, "cannot save changes")
	}
}

func (o *Opts) switchOrg() api.OrganizationSwitch {
	headers := o.F.GetAuth()
	reqConfig := defaults.SwitchOrganization(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.OrgID)
	res := reqConfig.Request()
	var orgSwitch api.OrganizationSwitch
	err := orgSwitch.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal")
		return api.OrganizationSwitch{}
	}
	return orgSwitch
}
