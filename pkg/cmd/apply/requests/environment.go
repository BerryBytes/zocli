package requests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
)

// CreateEnv
//
// this function is responsible for creating the environment
// it will get the information from the manifest file
// and it will create the environment
// NOTE: BY DEFAULT THE APPLICATIONS ARE SEARCHED ON THE DEFAULT ORGANIZATION WHICH IS ACTIVE ON THE CONTEXT
func (o *Opts) CreateEnv(e *models.Env) {
	err := e.ValidateCreation()
	if err != nil {
		o.F.IO.StopProgressIndicator()
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Print(" Invalid manifest file found.\n")
		o.F.Printer.Fatal(1, err.Error())
	}
	if e.Spec.Application.Name != "" {
		// get application by name and use that id for the creation of the environment
		o.GetApplicationByName(e)
	}
	o.F.IO.StopProgressIndicator()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Valid manifest file found.\n")

	o.F.IO.StartProgressIndicator()
	var envCreate EnvCreateRequest
	envCreate.Name = e.MetaData.Name
	envCreate.ApplicationID = e.Spec.Application.ID
	envCreate.ResoruceID = e.Spec.ResourceID
	envCreate.PluginVersionID = e.Spec.PluginVersionID
	envCreate.Replicas = e.Spec.Replicas
	envCreate.Version.Name = e.Spec.Version.Name
	envCreate.Version.Repo = e.Spec.Version.Repo
	envCreate.Version.Tag = e.Spec.Version.Tag

	headers := o.F.GetAuth()
	body, _ := json.Marshal(envCreate)
	reqConfig := defaults.EnvironmentCreate(o.F, map[string]interface{}{"headers": headers, "body": body})
	o.F.IO.StopProgressIndicator()
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Successfully created environment with name '" + envCreate.Name + "'\n")
	o.F.Printer.Exit(0)
}

// GetApplicationByName
//
// this function is responsible for getting the application id by name,
// if the application name is provided on the manifest file,
// if the project name is provided, then the application will be searched on that project
// NOTE: BY DEFAULT THE APPLICATIONS ARE SEARCHED ON THE DEFAULT ORGANIZATION WHICH IS ACTIVE ON THE CONTEXT
func (o *Opts) GetApplicationByName(e *models.Env) {
	// get application by name and use that id for the creation of the environment
	headers := o.F.GetAuth()
	reqConfig := defaults.ApplicationGetByName(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:name>", e.Spec.Application.Name)
	if e.Spec.Project.Name != "" {
		reqConfig.URL = reqConfig.URL + "?proj=" + e.Spec.Project.Name
	}
	if e.Spec.Project.ID != 0 {
		reqConfig.URL = reqConfig.URL + "?proj_id=" + fmt.Sprintf("%d", e.Spec.Project.ID)
	}

	o.F.IO.StopProgressIndicator()
	res := reqConfig.Request()
	var application api.Application
	err := application.FromJson(res.Data)
	o.F.IO.StartProgressIndicator()
	if err != nil {
		o.F.IO.StopProgressIndicator()
		o.F.Printer.Print(o.F.IO.ColorScheme().FailureIcon())
		o.F.Printer.Print(" Invalid manifest file found.\n")
		o.F.Printer.Fatal(1, err.Error())
		return
	}
	e.Spec.Application.ID = application.Id
}
