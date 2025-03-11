package defaults

import (
	"bytes"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

// GetOrganization
//
// this function is responsible for returning the necessary route and the method
// for fetching the details of a specific orgranization
//
// NOTE: this can only be called when an organization is added to the context other
// than the default organization, as the default organization i.e. of id 0 cannot
// be fetched
func GetOrganization(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	header, _ := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("organization_get"),
		Method:  "GET",
		Headers: header,
		F:       f,
	}
}

// GetOrganizations
//
// this function is responsible for returning the necessary route and the method
// for fetching the details of all orgranizations that are present on the scope of the user
func GetOrganizations(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	header, _ := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("organizations_get"),
		Method:  "GET",
		Headers: header,
		F:       f,
	}
}

func SwitchOrganization(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	header, _ := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("organization_switch"),
		Method:  "GET",
		Headers: header,
		F:       f,
	}
}

func DeleteOrganization(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	header, _ := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("organization_delete"),
		Method:  "DELETE",
		Headers: header,
		F:       f,
	}
}

func DeleteOrganizationMember(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	header, body := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("member_delete"),
		Method:  "DELETE",
		Headers: header,
		Body:    bytes.NewBuffer(body),
		F:       f,
	}
}

// CreateOrganization
//
// this  method is responsible for returning the RequestConfig instance which can be used by default
// for creating any organization
func CreateOrganization(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	header, body := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("organization_create"),
		Method:  "POST",
		Headers: header,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}
