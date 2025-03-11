package requests

import (
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
)

type Opts struct {
	F *factory.Factory
}

type Interface interface {
	CreateProject(*models.Project)
	CreateOrganization(*models.Organization)
	CreateApp(p *models.App)
	CreateMember(m *models.OrganizationMember)
	CreateEnv(e *models.Env)
}

func NewInterface(f *factory.Factory) Interface {
	return &Opts{F: f}
}
