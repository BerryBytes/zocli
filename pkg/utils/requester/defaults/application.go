package defaults

import (
	"bytes"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

func GetAppsList(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_apps"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func RenameApplication(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("app_rename"),
		Method:  "PUT",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func DeleteApplication(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("app_delete"),
		Method:  "DELETE",
		Headers: headers,
		F:       f,
	}
}

func GetSingleApplication(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("app_get_single"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func CreateApp(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("app_create"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func ApplicationGetByName(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("app_getbyname"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}
