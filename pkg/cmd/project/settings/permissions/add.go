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

func NewProjectPermissionsAddCommand(opts *Opts) *cobra.Command {
	listPermProj := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "add",
		Aliases:               []string{"a", "ad", "new"},
		Short:                 "add project scoped permissions to users",
		Long:                  grammar.ProjectPermissionAddHelp,
		SilenceErrors:         true,
		SilenceUsage:          true,
		Run:                   opts.AddPermissionRunner,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(opts.F)
		},
	}

	listPermProj.Flags().StringVarP(&opts.RoleName, "name", "n", "", "role name")
	listPermProj.Flags().StringVarP(&opts.UserEmail, "email", "e", "", "user email")
	listPermProj.Flags().IntVarP(&opts.ProjectID, "project", "p", 0, "project id")
	return listPermProj
}

// AddPermissionRunner
//
// This is the function that will be called when the command is run.
func (o *Opts) AddPermissionRunner(_ *cobra.Command, _ []string) {
	o.projectID()
	o.roleName()
	o.userEmail()
	o.AddPermission()
}

func (o *Opts) AddPermission() {
	body := []byte(`{"user_role_id": ` + fmt.Sprint(o.RoleID) + `, "email": "` + o.UserEmail + `", "project_id": ` + fmt.Sprint(o.ProjectID) + `}`)
	headers := o.F.GetAuth()
	reqConfig := defaults.AddProjectPermission(o.F, map[string]interface{}{"headers": headers, "body": body})
	_ = reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully given "+o.RoleName+" role to '"+o.UserEmail+"'\n")
	o.F.Printer.Exit(0)
}

// projectID
//
// This function is responsible for:
// 1. Getting the project id
// 2. Checking if the project id is valid
// 3. Assigning the project id to the ProjectID variable
func (o *Opts) projectID() {
	// if the project id is not supplied, then the active context is checked and used.
	if o.ProjectID == 0 && o.F.Config.ActiveContext != nil && o.F.Config.ActiveContext.DefaultProject != 0 {
		o.ProjectID = o.F.Config.ActiveContext.DefaultProject
	}
	if o.ProjectID == 0 {
		o.F.Printer.Print("Enter the project id: ")
		id, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, err.Error())
			return
		}
		validation.CheckValidID(o.F, strings.ReplaceAll(id, "\n", ""))
		o.ProjectID, _ = strconv.Atoi(strings.ReplaceAll(id, "\n", ""))
		return
	}
	o.F.Printer.Printf("Project id provided: " + strconv.Itoa(o.ProjectID))
}

// roleName
//
// This function is responsible for:
// 1. Getting the role name -> must be one of the roles that are available in the project
// 2. Calling the GetAllRoles function
// 3. Checking if the role name is valid
// 4. Assigning the role name to the RoleName variable
func (o *Opts) roleName() {
	o.F.Printer.Printf("Below are the roles that you can assign to any user,")
	o.I.GetAllRoles()
	if o.RoleName == "" {
		o.F.Printer.Print("Enter the role for the user: ")
		name, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, err.Error())
			return
		}
		o.RoleName = strings.ReplaceAll(name, "\n", "")
	} else {
		o.F.Printer.Printf("Role provided: " + o.RoleName)
	}

	found := false
	for _, role := range o.Roles.Roles {
		if strings.EqualFold(role.Name, o.RoleName) {
			found = true
			o.RoleID = role.Id
		}
	}

	if !found {
		o.F.Printer.Fatal(1, "invalid role provided")
		return
	}
}

// userEmail
//
// This function is responsible for:
// 1. Getting the user email -> must be a valid email
// 2. Assigning the user email to the UserEmail variable
func (o *Opts) userEmail() {
	if o.UserEmail == "" {
		o.F.Printer.Print("Enter the email address of the user: ")
		email, err := o.F.Term.ReadInput()
		if err != nil {
			o.F.Printer.Fatal(1, err.Error())
			return
		}
		o.UserEmail = strings.ReplaceAll(email, "\n", "")
	} else {
		o.F.Printer.Printf("Email provided: " + o.UserEmail)
	}
}
