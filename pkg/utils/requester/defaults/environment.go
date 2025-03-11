package defaults

import (
	"bytes"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

func EnvironmentGet(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_detail"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func EnvironmentGetOverview(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_overview"),
		Method:  "POST",
		Headers: headers,
		Body:    bytes.NewBuffer(body),
		F:       f,
	}
}

func EnvironmentGetList(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_apps_env"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func EnvironmentGetByName(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_getbyname"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func EnvironmentStopByID(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_stop"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func EnvironmentStartByID(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_start"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func EnvironmentDeleteByID(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_delete"),
		Method:  "DELETE",
		Headers: headers,
		F:       f,
	}
}

func EnvironmentRenameByID(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_rename"),
		Method:  "PUT",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func EnvironmentCreate(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("environment_create"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}
