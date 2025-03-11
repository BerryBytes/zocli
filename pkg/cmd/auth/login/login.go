package login

import (
	"fmt"
	"os"

	"github.com/berrybytes/zocli/pkg/utils/context"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/globalparser"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

type Opts struct {
	WebToken  string
	Email     string
	Password  string
	AuthToken string
	SsoCode   string
	F         *factory.Factory // f is the factory used for creating objects.

	WithToken   bool
	WithCreds   bool
	WithBrowser bool

	global *globalparser.GlobalFlags

	LoginResponse *api.LoginResponse
	Response      []byte
	cmd           *cobra.Command
}

// SaveConfig
//
// struct used to save the user's credentials.
type SaveConfig struct {
	Email     string `yaml:"email"`
	AuthToken string `yaml:"token"`
	ID        int    `yaml:"id"`
	WebToken  string `yaml:"xpersonaltoken"`
}

// NewLoginCommand
//
// returns a new instance of the login command.
// It takes a factory object f as input which is used to create other objects.
func NewLoginCommand(f *factory.Factory) *cobra.Command {
	opts := &Opts{F: f}
	login := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "login",
		Aliases:               []string{"l", "log", "signin", "sign-in", "sign_in", "sign_in", "signin"},
		Short:                 "login to 01cloud.com",
		Long:                  grammar.LoginHelp,
		RunE:                  opts.checkArgs,
	}

	login.PersistentFlags().BoolVarP(&opts.WithToken, "token", "t", false, "BOOL Auth token for your account")
	login.PersistentFlags().BoolVarP(&opts.WithBrowser, "sso", "s", false, "use browser to login")
	login.PersistentFlags().BoolVarP(&opts.WithCreds, "basic", "b", false, "BOOL use email and password to authenticate")

	// add a new flag set but mark it as hidden
	// as this approach is not recommended
	newToken := pflag.NewFlagSet("not recommended", pflag.ExitOnError)
	newToken.StringVarP(&opts.WebToken, "tokenVal", "T", "", "Token")
	newToken.StringVarP(&opts.Email, "email", "U", "", "Email")
	newToken.StringVarP(&opts.Password, "password", "P", "", "Password")

	newToken.VisitAll(func(flag *pflag.Flag) {
		flag.Hidden = true
	})

	login.Example = grammar.LoginExample
	login.Flags().AddFlagSet(newToken)
	opts.cmd = login
	return login
}

func (l *Opts) checkArgs(cmd *cobra.Command, args []string) error {
	l.global = globalparser.ParseGlobal(cmd, args, l.F)
	err := l.flow()
	return err
}

func (l *Opts) LoggedIn() error {
	if l.LoginResponse == nil {
		var loginResponse api.LoginResponse
		err := yaml.Unmarshal(l.Response, &loginResponse)
		if err != nil {
			return err
		}
		l.LoginResponse = &loginResponse
	}
	return l.saveDetails()
}

// saveDetails
//
// this function is responsible for saving the details of the user,
// i.e. email, token and userID on the token file on the config directory
// this function also calls the method for creating the default context
func (l *Opts) saveDetails() error {
	f, err := os.Create(l.F.Config.ConfigFolder + l.F.Config.AuthFile)
	if err != nil {
		return err
	}

	var save SaveConfig
	save.Email = l.LoginResponse.Data.User.Email
	save.AuthToken = l.LoginResponse.Data.AuthToken
	save.ID = l.LoginResponse.Data.User.Id
	save.WebToken = l.LoginResponse.WebToken

	yamlF, err := yaml.Marshal(save)
	if err != nil {
		return err
	}

	_, err = f.Write(yamlF)
	if err != nil {
		return err
	}
	l.F.Printer.Print(l.F.IO.ColorScheme().SuccessIcon(), " Logged in successfully as ", l.LoginResponse.User.Email)
	fmt.Println()
	// create a default context for the user with the orgranization 0 i.e. Default Organization
	context.SetFirstActiveContext(l.F)
	return nil
}
