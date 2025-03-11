package defaults

import (
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

// BasicLogin
//
// defaults for basic login mechanism i.e. email and password
func BasicLogout(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	header, _ := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("logout"),
		Method:  "POST",
		Headers: header,
		F:       f,
	}
}
