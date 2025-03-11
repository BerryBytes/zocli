package defaults

import (
	"bytes"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

// BasicLogin
//
// defaults for basic login mechanism i.e. email and password
func BasicLogin(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	_, body := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:    f.Routes.GetRoute("login"),
		Method: "POST",
		Body:   bytes.NewBuffer(body),
		F:      f,
	}
}

// RequestSSO
//
// defaults for requesting the sso token to the server
func RequestSSO(f *factory.Factory, _ ...map[string]interface{}) *requester.RequestConfig {
	return &requester.RequestConfig{
		URL:    f.Routes.GetRoute("sso-code"),
		Method: "POST",
		F:      f,
	}
}

// RequestSSOStatus
//
// defaults for requesting the sso token status to the server
func RequestSSOStatus(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	_, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:    f.Routes.GetRoute("sso-code"),
		Method: "GET",
		Body:   bytes.NewBuffer(body),
		F:      f,
	}
}
