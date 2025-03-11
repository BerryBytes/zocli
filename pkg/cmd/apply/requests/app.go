package requests

import (
	"encoding/json"

	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
)

func (o *Opts) CreateApp(app *models.App) {
	err := app.ValidateCreation()
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
	appRequest := AppCreateRequest{
		Name:                app.MetaData.Name,
		ProjectID:           app.Spec.Project.ID,
		PluginID:            app.Spec.Plugin.ID,
		ClusterID:           app.Spec.Cluster.ID,
		ChartID:             app.Spec.ChartID,
		GitRepository:       app.Spec.GitRepository,
		GitRepoUrl:          &app.Spec.GitRepoUrl,
		GitService:          app.Spec.GitService,
		ImageUrl:            app.Spec.ImageUrl,
		ImageNamespace:      app.Spec.ImageNamespace,
		ImageRepo:           app.Spec.ImageRepo,
		Region:              app.Spec.Cluster.Region,
		ImageService:        app.Spec.ImageService,
		ServiceType:         int(app.Spec.ServiceType),
		OperatorPackageName: app.Spec.OperatorPackageName,
	}
	headers := o.F.GetAuth()
	body, _ := json.Marshal(appRequest)
	reqConfig := defaults.CreateApp(o.F, map[string]interface{}{"headers": headers, "body": body})
	o.F.IO.StopProgressIndicator()
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Print(" Successfully created app with name '" + app.MetaData.Name + "'\n")
	o.F.Printer.Exit(0)
}
