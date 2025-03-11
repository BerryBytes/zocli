package requests

import (
	"encoding/json"

	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
)

func (o *Opts) CreateProject(p *models.Project) {
	err := p.ValidateCreation()
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
	var projectCreate ProjectCreateRequest
	projectCreate.Name = p.MetaData.Name
	projectCreate.SubscriptionID = p.Spec.Subscription.ID
	projectCreate.Logging = p.Spec.Logging
	projectCreate.OptimizeCost = p.Spec.OptimizeCost
	projectCreate.DedicatedLb = p.Spec.DedicatedLB

	headers := o.F.GetAuth()
	body, _ := json.Marshal(projectCreate)
	reqConfig := defaults.CreateProject(o.F, map[string]interface{}{"headers": headers, "body": body})
	o.F.IO.StopProgressIndicator()
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Successfully created project with name '" + projectCreate.Name + "'\n")
	o.F.Printer.Exit(0)
}
