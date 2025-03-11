package manifestprocessor

import (
	"log"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
)

// MakeManifest
//
// generate the manifest file from the data which is received
// from the request response, supplied as the argument
func (o *Opts) MakeManifest(model string, data interface{}) interface{} {
	switch strings.ToLower(model) {
	case "project":
		pro, Ok := data.(api.Project)
		if Ok {
			return o.makeProjectManifest(&pro)
		}
	case "organization":
		org, Ok := data.(api.Organization)
		if Ok {
			return o.makeOrganizationManifest(&org)
		}
	case "member":
		mem, Ok := data.(api.OrganizationMember)
		if Ok {
			return o.makeOrganizationMemberManifest(&mem)
		}
	case "application":
		app, Ok := data.(api.Application)
		if Ok {
			return o.makeApplicationManifest(&app)
		}
	case "environment":
	default:
		o.F.Printer.Fatal(1, "no such kind of manifest found")
	}
	return nil
}

// makeProjectManifest
//
// creates a manifest file from the project data received
func (o *Opts) makeProjectManifest(data *api.Project) *models.Project {
	manifestProject := models.Project{}
	manifestProject.MetaData.ID = data.ID
	manifestProject.MetaData.CreatedAt = data.Createdat
	manifestProject.MetaData.Name = data.Name
	manifestProject.MetaData.OrganizationId = data.OrganizationId
	err := utils.ConvertType(data, &manifestProject.Spec)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal data")
	}
	manifestProject.Spec.Variables = data.Variables
	manifestProject.ApiVersion = "app.01cloud.io/v1"
	manifestProject.Kind = "project"
	return &manifestProject
}

// makeOrganizationManifest
//
// creates a manifest file from the Organization data received
func (o *Opts) makeOrganizationManifest(data *api.Organization) *models.Organization {
	manifestOrg := models.Organization{}
	manifestOrg.MetaData.ID = data.ID
	manifestOrg.MetaData.Name = data.Name
	manifestOrg.MetaData.CreatedAt = data.Createdat
	err := manifestOrg.Spec.FromJson(data)
	if err != nil {
		log.Fatal(err)
	}
	manifestOrg.ApiVersion = "app.01cloud.io/v1"
	manifestOrg.Kind = "organization"

	return &manifestOrg
}

// makeOrganizationMemberManifest
//
// creates a manifest file from the Organization Member data received
func (o *Opts) makeOrganizationMemberManifest(data *api.OrganizationMember) *models.OrganizationMember {
	manifestOrgMem := models.OrganizationMember{}
	manifestOrgMem.MetaData.ID = data.ID
	manifestOrgMem.MetaData.Name = "members"
	manifestOrgMem.MetaData.CreatedAt = data.Createdat
	manifestOrgMem.Spec = models.MemberSpec{
		Email: data.Owner.Email,
		Role:  models.GetRole(data.UserRole),
	}
	err := utils.ConvertType(data, &manifestOrgMem.Spec)
	if err != nil {
		log.Fatal(err)
	}
	manifestOrgMem.ApiVersion = "app.01cloud.io/v1"
	manifestOrgMem.Kind = "member"

	return &manifestOrgMem
}

// makeAppManifest
//
// creates a manifest file from the app data received
func (o *Opts) makeApplicationManifest(data *api.Application) *models.App {
	manifestApp := models.App{}
	manifestApp.MetaData.ID = data.Id
	manifestApp.MetaData.CreatedAt = data.Createdat
	manifestApp.MetaData.Name = data.Name
	err := utils.ConvertType(data, &manifestApp.Spec)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot unmarshal data")
	}
	manifestApp.ApiVersion = "app.01cloud.io/v1"
	manifestApp.Kind = "application"
	return &manifestApp
}
