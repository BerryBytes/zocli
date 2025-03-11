package config

import (
	"os"
)

type Routes map[string]string

var routes Routes

func Load() *Routes {
	routes = make(Routes)

	frontEndBase := os.Getenv("CLI_FRONT_BASE")
	if frontEndBase == "" {
		frontEndBase = "https://console.01cloud.io/"
	}

	baseURL := os.Getenv("CLI_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.01cloud.io/"
	}

	// login routes
	routes["login"] = baseURL + "user/login"
	routes["sso-code"] = baseURL + "user/login/sso"
	routes["device-login"] = frontEndBase + "login/sso"

	routes["profile"] = baseURL + "profile"
	routes["project_list"] = baseURL + "projects"
	routes["project_detail"] = baseURL + "project/<:id>"
	routes["project_enable_id"] = baseURL + "project/<:id>/activation"
	routes["project_getbyname"] = baseURL + "getbyname/project/<:name>"
	routes["project_rename"] = baseURL + "rename/project/<:id>"
	routes["project_delete"] = routes["project_detail"]
	routes["project_create"] = baseURL + "project"
	routes["project_resource"] = routes["project_detail"] + "/resource"
	routes["project_permissions"] = routes["project_detail"] + "/users"
	routes["project_permission_add"] = routes["project_detail"] + "/user"
	routes["list_project_loadbalancer"] = routes["project_detail"] + "/loadbalancers"
	routes["create_project_loadbalancer"] = baseURL + "loadbalancer"
	routes["loadbalancer"] = baseURL + "loadbalancer/<:id>"
	routes["system_roles"] = baseURL + "roles"
	routes["project_roles_update"] = routes["project_detail"] + "/user/<:permissionId>"
	routes["project_permission_delete"] = routes["project_roles_update"]
	// note the id must be at last like, projects/1
	routes["project_apps"] = baseURL + "project/<:id>" + "/applications?search=&plugin_id=all&region=all"
	routes["app_rename"] = baseURL + "rename/application/<:id>"
	routes["app_create"] = baseURL + "application"
	routes["app_delete"] = baseURL + "application/<:id>"
	routes["app_get_single"] = routes["app_delete"]
	routes["app_getbyname"] = baseURL + "getbyname/application/<:name>"
	// note the id must be at last like, projects/1/applications/1/env
	routes["project_apps_env"] = baseURL + "application/<:id>" + "/environments"
	routes["environment_detail"] = baseURL + "environment/<:id>"
	routes["environment_getbyname"] = baseURL + "getbyname/environment/<:name>"
	routes["environment_overview"] = baseURL + "environment/<:id>/overview"
	routes["environment_stop"] = baseURL + "environment/<:id>/stop"
	routes["environment_start"] = baseURL + "environment/<:id>/start"
	routes["environment_delete"] = baseURL + "environment/<:id>"
	routes["environment_rename"] = baseURL + "rename/environment/<:id>"
	routes["environment_create"] = baseURL + "environment"
	routes["logout"] = baseURL + "user/logout"
	routes["organization_get"] = baseURL + "organization"
	routes["organization_delete"] = baseURL + "organization"
	routes["organization_create"] = baseURL + "organization"
	routes["organizations_get"] = baseURL + "organizations"
	routes["organization_switch"] = baseURL + "organization/<:id>/switch"
	routes["member_add"] = baseURL + "organization/members"
	routes["member_delete"] = baseURL + "organization/members"

	routes["cluster_create"] = baseURL + "import-cluster"
	routes["create_cluster"] = baseURL + "create-cluster"
	routes["cluster_status"] = baseURL + "create-cluster/<:id>/package-status"
	routes["package_config"] = baseURL + "package-config"
	routes["package_install"] = baseURL + "create-cluster/<:id>/install-package"
	routes["package_uninstall"] = baseURL + "create-cluster/<:id>/uninstall-package"

	return &routes
}

func (r *Routes) GetRoute(route string) string {
	return routes[route]
}
