package logout

import (
	"fmt"
	"os"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/spf13/cobra"
)

type Opts struct {
	f *factory.Factory

	NoInvalidate bool
}

func NewLogOutCmd(f *factory.Factory) *cobra.Command {
	opts := &Opts{f: f}
	logout := &cobra.Command{
		Use:   "logout",
		Short: "Logout of this device",
		Long:  grammar.LogOutHelp,
		RunE:  opts.logoutCmd,
	}

	logout.PersistentFlags().BoolVarP(&opts.NoInvalidate, "device-only", "d", false, "Do not invalidate the session")

	return logout
}

func (l *Opts) logoutCmd(cmd *cobra.Command, _ []string) error {
	if !l.f.LoggedIn {
		l.f.Printer.Fatal(1, "not logged in")
	}

	if !l.NoInvalidate {
		l.requestLogout()
	}
	err := os.RemoveAll(l.f.Config.ConfigFolder + l.f.Config.AuthFile)
	if err != nil {
		l.f.Printer.Fatal(1, err)
	}
	fmt.Println(l.f.IO.ColorScheme().SuccessIcon(), " Successfully Logged out.")
	return nil
}

func (l *Opts) requestLogout() {
	header := map[string]string{"Authorization": "Basic " + l.f.UserAuthToken}
	reqConfig := defaults.BasicLogout(l.f, map[string]interface{}{"headers": header})

	_ = reqConfig.Request()
}
