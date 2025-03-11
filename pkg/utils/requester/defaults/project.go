package defaults

import (
	"bytes"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

func GetProjectList(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_list"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func GetProjectDetail(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_detail"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func GetProjectEnableByID(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_enable_id"),
		Method:  "POST",
		Headers: headers,
		Body:    bytes.NewBuffer(body),
		F:       f,
	}
}

func GetProjectByName(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_getbyname"),
		Method:  "GET",
		Headers: headers,
		Body:    bytes.NewBuffer(body),
		F:       f,
	}
}

func RenameProject(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_rename"),
		Method:  "PUT",
		Headers: headers,
		Body:    bytes.NewBuffer(body),
		F:       f,
	}
}

func DeleteProject(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_delete"),
		Method:  "DELETE",
		Headers: headers,
		Body:    bytes.NewBuffer(body),
		F:       f,
	}
}

func GetProjectResource(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_resource"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func UpdateVarProject(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_detail"),
		Method:  "PUT",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func ListProjectPermissions(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_permissions"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func ListProjectLoadbalancer(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("list_project_loadbalancer"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func GetByIdLoadbalancer(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("loadbalancer"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func CreateProjectLoadbalancer(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("create_project_loadbalancer"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func DeleteLoadbalancer(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("loadbalancer"),
		Method:  "DELETE",
		Headers: headers,
		F:       f,
	}
}

func ListProjectRoles(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("system_roles"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

func UpdateUserRoles(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_roles_update"),
		Method:  "PUT",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func DeletePermission(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_permission_delete"),
		Method:  "DELETE",
		Headers: headers,
		F:       f,
	}
}

func CreateProject(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_create"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func Addmember(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("member_add"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

func AddProjectPermission(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)
	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("project_permission_add"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}
