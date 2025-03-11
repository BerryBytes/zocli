package defaults

import (
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

// Profile
//
// defaults for fetching profile
func Profile(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("profile"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}
