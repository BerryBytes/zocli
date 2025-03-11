package permissions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectPermissionsDeleteCommand(opts *Opts) *cobra.Command {
	delPermProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "delete",
		Aliases:               []string{"d", "de", "del", "dele", "delete", "remove", "rem", "r"},
		Short:                 "delete project scoped permissions",
		Long:                  grammar.ProjectPermissionDeleteHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.I.DeletePermission,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)

			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
		},
	}

	delPermProj.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	delPermProj.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	delPermProj.Flags().Int64VarP(&opts.PermissionID, "pid", "p", -1, "permission id")
	return delPermProj
}

func (o *Opts) DeletePermission(cmd *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.I.GetPermission()
	}

	if o.Name != "" {
		proj := o.I.GetProjectDetailByName()
		o.ID = fmt.Sprint(proj.ID)
		o.I.GetPermission()
	}
	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.ID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
		o.I.GetPermission()
	}

	if o.ID == "" && o.Name == "" {
		_ = cmd.Help()
		o.F.Printer.Exit(0)
		return
	}

	if len(o.Permissions.Permissions) == 0 {
		o.F.Printer.Fatal(1, "no permissions found")
		return
	}

	if o.PermissionID == -1 {
		o.F.Printer.Print("Permission ID : ")
		var id string
		_, _ = fmt.Scanln(&id)
		id, _ = validation.GetIDOrName(id)
		if id == "" {
			o.F.Printer.Fatal(1, "invalid id")
			return
		}
		o.PermissionID, _ = strconv.ParseInt(id, 10, 64)
	}

	o.checkIfValidID()
	o.Delete()
}

func (o *Opts) Delete() {
	reqConfig := defaults.DeletePermission(o.F, map[string]interface{}{"headers": o.F.GetAuth()})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:permissionId>", fmt.Sprint(o.PermissionID))

	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully deleted permission with id '"+fmt.Sprint(o.PermissionID)+"'")
}
