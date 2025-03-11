package permissions

import (
	"fmt"
	"strings"
	"time"

	"github.com/berrybytes/zocli/pkg/utils/formatter"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/berrybytes/zocli/pkg/utils/validation"
	"github.com/spf13/cobra"
)

func NewProjectPermissionsListCommand(opts *Opts) *cobra.Command {
	listPermProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "list",
		Aliases:               []string{"l", "li", "lis", "get", "g", "ge"},
		Short:                 "list project scoped permissions",
		Long:                  grammar.ProjectPermissionListHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.I.ListPermissions,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)

			if len(args) != 0 {
				opts.ID, opts.Name = validation.GetIDOrName(args[0])
			}
		},
	}

	listPermProj.Flags().StringVarP(&opts.ID, "id", "i", "", "project id")
	listPermProj.Flags().StringVarP(&opts.Name, "name", "n", "", "project name")
	listPermProj.Flags().StringVarP(&opts.Output, "out", "o", "", "output in json or yaml")
	return listPermProj
}

func (o *Opts) ListPermissions(cmd *cobra.Command, _ []string) {
	if o.ID != "" && o.Name != "" {
		o.F.Printer.Fatal(1, "provide either id or name, but not both")
		return
	}

	if o.ID != "" {
		validation.CheckValidID(o.F, o.ID)
		o.I.GetPermission()
		o.F.Printer.Exit(0)
		return
	}

	if o.Name != "" {
		proj := o.I.GetProjectDetailByName()
		o.ID = fmt.Sprint(proj.ID)
		o.I.GetPermission()
		o.F.Printer.Exit(0)
		return
	}

	if o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Project Value.\n")
		o.ID = fmt.Sprint(o.F.Config.ActiveContext.DefaultProject)
		o.I.GetPermission()
		o.F.Printer.Exit(0)
		return
	}

	_ = cmd.Help()
}

func (o *Opts) GetPermission() {
	headers := o.F.GetAuth()
	reqConfig := defaults.ListProjectPermissions(o.F, map[string]interface{}{"headers": headers})
	reqConfig.URL = strings.ReplaceAll(reqConfig.URL, "<:id>", o.ID)
	res := reqConfig.Request()

	err := o.Permissions.FromJson(res.Data)
	if err != nil {
		o.F.Printer.Fatal(9, "cannot marshal")
	}
	o.I.PrintPermissions()
}

func (o *Opts) PrintPermissions() {
	if o.Output == "json" {
		if len(o.Permissions.Permissions) == 1 {
			formatter.PrintJson(o.F, o.Permissions.Permissions[0])
			return
		}
		formatter.PrintJson(o.F, o.Permissions.Permissions)
		return
	} else if o.Output == "yaml" {
		if len(o.Permissions.Permissions) == 1 {
			formatter.PrintYaml(o.F, o.Permissions.Permissions[0])
			return
		}
		formatter.PrintYaml(o.F, o.Permissions.Permissions)
		return
	}

	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tablePrinter := table.New(o.F, 0)
	tablePrinter.HeaderRow(color, "id", "name", "email", "role", "Since")

	for _, user := range o.Permissions.Permissions {
		tablePrinter.AddField(fmt.Sprint(user.Id))
		tablePrinter.AddField(user.User.FirstName + " " + user.User.LastName)
		tablePrinter.AddField(user.Email)
		tablePrinter.AddField(user.UserRole.Name)
		now := time.Now()
		tablePrinter.AddField(table.RelativeTimeAgo(now, user.CreatedAt))
		tablePrinter.EndRow()
	}

	_ = tablePrinter.Print()
}
