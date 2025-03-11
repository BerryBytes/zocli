package permissions

import (
	"fmt"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/spf13/cobra"
)

type Opts struct {
	F              *factory.Factory
	ID             string
	Name           string
	Output         string
	PermissionID   int64
	ProjectID      int
	Permissions    api.Permissions
	Change         api.Permission
	Roles          api.Roles
	UpdateRole     string
	RoleID         int
	I              Interface
	PermissionName string
	UserEmail      string
	RoleName       string
}

func NewProjectPermissionsCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	permProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "permission",
		Aliases:               []string{"p", "pe", "perm", "per", "permiss"},
		Short:                 "project scoped permissions",
		Long:                  grammar.ProjectPermissionHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
		},
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	permProj.AddGroup(&cobra.Group{ID: "actions", Title: "Action commands"})
	o := &Opts{F: f}
	o.I = NewPermissionInterface(o)
	permProj.AddCommand(NewProjectPermissionsListCommand(o))
	permProj.AddCommand(NewProjectPermissionsUpdateCommand(o))
	permProj.AddCommand(NewProjectPermissionsDeleteCommand(o))
	permProj.AddCommand(NewProjectPermissionsAddCommand(o))
	return permProj
}

func (o *Opts) GetProjectDetailByName() *api.Project {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectByName(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:name>", o.Name)

	res := reqConfig.Request()
	if res == nil {
		return nil
	}

	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}

	return o.GetProjectDetail(fmt.Sprint(project.ID))
}

func (o *Opts) GetProjectDetail(id string) *api.Project {
	header := o.F.GetAuth()
	reqConfig := defaults.GetProjectDetail(o.F, map[string]interface{}{"headers": header})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", fmt.Sprint(id))
	res := reqConfig.Request()
	if res == nil {
		return nil
	}

	var project api.Project
	err := project.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, err)
	}
	return &project
}
