package member

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/spf13/cobra"
)

func NewOrganizationMemberDeleteCommand(o *Opts) *cobra.Command {
	del := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d", "del", "de", "dele"},
		Short:   "organizations members details",
		Long:    grammar.OrganizationMemberDeleteHelp,
		GroupID: "basic",
		Run:     o.I.DeleteRunner,
		PreRun: func(_ *cobra.Command, args []string) {
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	return del
}

type DeleteRequest struct {
	Email string `json:"email" yaml:"email"`
}

func (o *Opts) DeleteRunner(cmd *cobra.Command, _ []string) {
	o.GetOrganizationMembers()
	o.PrintOrganizationMembersTable()
	o.confirmAndDelete()
	_ = cmd.Help()
}

func (o *Opts) confirmAndDelete() {
	fmt.Print("Enter the email to remove/delete member :: ")
	text, err := o.F.Term.ReadInput()
	if err != nil {
		o.F.Printer.Fatal(5, err)
	}
	email := strings.TrimSpace(strings.Trim(text, "\n"))
	if email == "" {
		o.F.Printer.Fatal(1, "required email address")
	}
	proceed := false
	prompt := &survey.Confirm{
		Message: "Do you sure want to delete/remove the organization member ?",
		Default: false,
	}

	err = survey.AskOne(prompt, &proceed)
	if err != nil {
		o.F.Printer.Fatalf(5, "cannot proceed.\nErr: %v\n", err)
	}

	if !proceed {
		o.F.Printer.Exit(0)
	}
	req := DeleteRequest{
		Email: email,
	}
	header := o.F.GetAuth()
	body, _ := json.Marshal(req)
	reqConfig := defaults.DeleteOrganizationMember(o.F, map[string]interface{}{"headers": header, "body": body})
	res := reqConfig.Request()
	if res == nil {
		return
	}

	o.F.Printer.Println(o.F.IO.ColorScheme().SuccessIcon(), " Successfully removed '"+email+"'")
	o.F.Printer.Exit(0)
}
