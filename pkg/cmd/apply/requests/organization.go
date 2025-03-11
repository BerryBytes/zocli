package requests

import (
	"encoding/json"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/cmd/organization"
	"github.com/berrybytes/zocli/pkg/utils/context"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
)

func (o *Opts) CreateOrganization(org *models.Organization) {
	err := org.ValidateCreation()
	if err != nil {
		o.F.IO.StopProgressIndicator()
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Print(" Invalid manifest file found.\n")
		o.F.Printer.Fatal(1, err.Error())
	}
	o.F.IO.StopProgressIndicator()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Valid manifest file found.\n")

	o.F.IO.StartProgressIndicator()
	orgCreate := o.parseVals(org)
	res := o.createOrganization(orgCreate)
	var orgSwitch api.OrganizationSwitch
	err = orgSwitch.FromJson(res.Data)
	if err != nil {
		o.F.IO.StopProgressIndicator()
		o.F.Printer.Fatal(9, "cannot unmarshal")
		return
	}

	o.changeContext(&orgSwitch)
}

// parseVals
//
// this function initializes the OrganizationCreateRequest from the organization models supplied
func (o *Opts) parseVals(org *models.Organization) *OrganizationCreateRequest {
	return &OrganizationCreateRequest{
		Name:                org.MetaData.Name,
		OrganizationPlainID: org.Spec.OrganizationPlan.ID,
		Domain:              org.Spec.Domain,
		Description:         org.Spec.Description,
	}
}

// createOrganization
//
// actual function which is responsible to call the requester for creating organization
func (o *Opts) createOrganization(org *OrganizationCreateRequest) *api.BaseResponse {
	headers := o.F.GetAuth()
	body, _ := json.Marshal(org)
	reqConfig := defaults.CreateOrganization(o.F, map[string]interface{}{"headers": headers, "body": body})
	o.F.IO.StopProgressIndicator()
	res := reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Successfully created organization with name '" + org.Name + "'\n")
	return res
}

// changeContext
//
// this function is responsible for changing the context active organization
// to the new organization which has just been created
func (o *Opts) changeContext(orgSwitch *api.OrganizationSwitch) {
	var orgUse organization.Opts
	orgUse.F = o.F
	orgUse.SaveTokenChanges(orgSwitch)
	o.F = context.SetActiveFalse(o.F)
	orgUse.SaveContextChanges(orgSwitch.Organization.ID, orgSwitch)
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Successfully changed the context to new organization.")
	o.F.Printer.Exit(0)
}
