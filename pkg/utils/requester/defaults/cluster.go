package defaults

import (
	"bytes"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester"
)

// CreateCluster
//
// @Description: create cluster
// @param f
// @param args (headers, body)
// @return *requester
// Description: default requester config for creating cluster
func CreateCluster(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("cluster_create"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

// GetCluster
//
// @Description: get cluster
// @param f
// @param args (headers, body)
// @return *requester
// Description: default requester config for getting cluster
func GetCluster(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("create_cluster"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

// InstallPackage
//
// @Description: install package
// @param f
// @param args (headers, body)
// @return *requester
// Description: default requester config for deleting packages of cluster
func InstallPackage(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("package_install"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

// UninstallPackage
//
// @Description: uninstall package
// @param factory
// @param args (headers, body) for request
// @return *requestorConfig
// Description: default requester config for uninstalling packages of cluster
func UninstallPackage(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, body := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("package_uninstall"),
		Method:  "POST",
		Headers: headers,
		F:       f,
		Body:    bytes.NewBuffer(body),
	}
}

// StatusForPackage
//
// @Description: get status for packages
// @param factory
// @param args (headers, body) for request
// @return *requestorConfig
// Description: default requester config for getting status for packages
func StatusForPackage(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("cluster_status"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}

// PackageConfig
//
// @Description: get config for all packages that exists on 01cloud
// @param *factory
// @param args (headers, body) for request
// @return *requestorConfig
// Description: default requester config for getting config for all packages
func PackageConfig(f *factory.Factory, args ...map[string]interface{}) *requester.RequestConfig {
	headers, _ := headersAndBodyConv(args...)

	return &requester.RequestConfig{
		URL:     f.Routes.GetRoute("package_config"),
		Method:  "GET",
		Headers: headers,
		F:       f,
	}
}
