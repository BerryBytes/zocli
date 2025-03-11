package actions

import (
	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils/factory"
)

type Opts struct {
	ID      string
	Name    string
	F       *factory.Factory
	project *api.Project
}
