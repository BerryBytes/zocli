package permissions

import (
	"github.com/berrybytes/zocli/api"
	"github.com/spf13/cobra"
)

type Interface interface {
	GetPermission()
	PrintPermissions()
	GetProjectDetailByName() *api.Project
	GetProjectDetail(id string) *api.Project
	GetAllRoles()
	ListPermissions(*cobra.Command, []string)
	DeletePermission(*cobra.Command, []string)
	UpdatePermissions(*cobra.Command, []string)
}

func NewPermissionInterface(o *Opts) Interface {
	return o
}
