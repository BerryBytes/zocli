package status

import (
	"fmt"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/spf13/cobra"
)

type Opts struct {
	f *factory.Factory
}

func NewStatusCmd(f *factory.Factory) *cobra.Command {
	opts := &Opts{f}
	status := &cobra.Command{
		Use:     "status",
		Aliases: []string{"stat", "me"},
		Short:   "View authentication status",
		Long:    grammar.StatusHelp,
		RunE:    opts.statusCmd,
	}
	return status
}

// statusCmd
//
// this function is responsible for checking if the token is valid
// and also prints either the web token is used or the auth token
func (s *Opts) statusCmd(cmd *cobra.Command, _ []string) error {
	headers := s.f.GetAuth()
	if s.f.WebTokenUsed {
		fmt.Println(s.f.IO.ColorScheme().SuccessIcon(), "Web Token found")
	} else {
		fmt.Println(s.f.IO.ColorScheme().SuccessIcon(), "Auth Token found")
	}
	reqConfig := defaults.Profile(s.f, map[string]interface{}{"headers": headers})
	res := reqConfig.Request()
	var profile api.ProfileResponse
	err := profile.FromJson(res.Data)
	if err != nil {
		s.f.Printer.Fatal(9, err)
	}
	s.printProfile(&profile)
	s.f.Printer.Fatal(1, "not logged in")
	return cmd.Help()
}

// printProfile
//
// this function prints the profile of the user on the tabular structure
// and also prints the quotas of the user
func (s *Opts) printProfile(profile *api.ProfileResponse) {
	fmt.Println(s.f.IO.ColorScheme().SuccessIcon(), " Token is valid")
	fmt.Println(s.f.IO.ColorScheme().SuccessIcon(), " Fetched Profile")
	fmt.Println()

	color := s.f.IO.ColorScheme().ColorFromString("blue")
	table := table.New(s.f, 0)
	table.HeaderRow(color, "profile", "")

	table.AddField("Name :")
	table.AddField(profile.FirstName + profile.LastName)
	table.EndRow()
	table.AddField("Email :")
	table.AddField(profile.Email)
	table.EndRow()
	table.AddField("Company :")
	table.AddField(profile.Company)
	table.EndRow()

	table.HeaderRow(color, "quotas", "")

	table.AddField("Projects :")
	table.AddField(fmt.Sprintf("%d", profile.Quotas.UserProject))
	table.EndRow()
	table.AddField("Organizations :")
	table.AddField(fmt.Sprintf("%d", profile.Quotas.UserOrganization))
	table.EndRow()
	err := table.Print()
	if err != nil {
		s.f.Printer.Error(err)
	}

	s.f.Printer.Exit(0)
}
