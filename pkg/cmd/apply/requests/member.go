package requests

import (
	"encoding/json"

	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
)

func (o *Opts) CreateMember(m *models.OrganizationMember) {
	err := m.ValidateCreation()
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
	var req MemberCreateRequest
	req.Email = m.Spec.Email
	req.Role = models.SetRole(m.Spec.Role)

	headers := o.F.GetAuth()
	body, _ := json.Marshal(req)
	reqConfig := defaults.Addmember(o.F, map[string]interface{}{"headers": headers, "body": body})
	o.F.IO.StopProgressIndicator()
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Printf("Successfully added %s as %s role. \n", req.Email, m.Spec.Role)
	o.F.Printer.Exit(0)
}
