package permissions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
)

func NewProjectPermissionsUpdateCommand(opts *Opts) *cobra.Command {
	updatePermProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "update",
		Aliases:               []string{"up", "u", "upd", "upda", "update"},
		Short:                 "update project scoped Permissions",
		Long:                  grammar.ProjectPermissionUpdateHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.I.UpdatePermissions,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)

			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
		},
	}

	updatePermProj.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	updatePermProj.Flags().Int64VarP(&opts.PermissionID, "pid", "p", -1, "permission id")
	updatePermProj.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	return updatePermProj
}

func (o *Opts) UpdatePermissions(cmd *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.I.GetPermission()
	}

	if o.Name != "" {
		o.ID = fmt.Sprint(o.I.GetProjectDetailByName().ID)
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
	o.GetAllRoles()
	o.askUpdates()
	valid := o.checkValidRole()
	if !valid {
		o.F.Printer.Fatal(1, "no such role available")
	}
	o.pushChanges()
}

func (o *Opts) pushChanges() {
	headers := o.F.GetAuth()
	body := []byte(`{"project_id": ` + fmt.Sprint(o.ID) + `, "user_role_id": ` + fmt.Sprint(o.RoleID) + `}`)
	reqConfig := defaults.UpdateUserRoles(o.F, map[string]interface{}{"headers": headers, "body": body})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:permissionId>", fmt.Sprint(o.Change.Id))
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully role changed to '"+o.UpdateRole+"'")
	o.F.Printer.Exit(0)
}

func (o *Opts) checkIfValidID() {
	valid := false
	for _, permission := range o.Permissions.Permissions {
		if permission.Id == int(o.PermissionID) {
			valid = true
			o.Change = permission
		}
	}

	if !valid {
		o.F.Printer.Fatal(1, "no such permission found.")
	}
}

func (o *Opts) askUpdates() {
	input, _ := readline.New("Role : ")
	defer func(input *readline.Instance) {
		_ = input.Close()
	}(input)

	datax := o.Change.UserRole.Name
	data2 := []byte(datax)
	_, _ = input.WriteStdin(data2)

	updates, _ := input.Readline()
	o.UpdateRole = updates
}

func (o *Opts) checkValidRole() bool {
	valid := false
	for _, role := range o.Roles.Roles {
		if strings.EqualFold(role.Name, o.UpdateRole) {
			valid = true
			o.RoleID = role.Id
		}
	}

	return valid
}

func (o *Opts) GetAllRoles() {
	headers := o.F.GetAuth()
	reqConfig := defaults.ListProjectRoles(o.F, map[string]interface{}{"headers": headers})
	res := reqConfig.Request()
	if res == nil {
		return
	}
	err := o.Roles.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "unable to marshal")
	}

	o.printRoles()
}

func (o *Opts) printRoles() {
	tablePrinter := table.New(o.F, 0)
	tablePrinter.Separator()
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tablePrinter.HeaderRow(color, "role", "description")
	for _, role := range o.Roles.Roles {
		if !role.Active {
			continue
		}
		tablePrinter.AddField(role.Name)
		tablePrinter.AddField(role.Description)
		tablePrinter.EndRow()
	}
	_ = tablePrinter.Print()
	tablePrinter.Separator()
}
