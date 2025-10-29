package factory

import (
	"context"
	"net/http"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"

	"github.com/AlecAivazis/survey/v2"
	"github.com/berrybytes/zocli/internal/browser"
	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/iostreams"
	"github.com/berrybytes/zocli/pkg/utils/printer"
	"github.com/berrybytes/zocli/pkg/utils/terminal"
	cliBrowser "github.com/cli/browser"
	// "github.com/sirupsen/logrus"
)

// Factory
//
// all the root commands flags will be saved here,
// along with all other default fields
// and then, will be transferred to all the sub-command
type Factory struct {
	// determines the verbose mode of the CLI
	Verbose bool
	// determines the quiet mode of the CLI
	Quiet bool
	// determines the interactive mode of the CLI
	NoInteractive bool

	// determines if the user is logged in or not
	LoggedIn bool
	// configuration for running the CLI
	Config *config.Config
	// context for running the CLI
	Ctx context.Context
	// user's operating system
	UserOS string
	// user's home folder
	UserHomeFolder string
	// determines if the configuration folder has been created or not
	ConfigCreated bool
	// http client for sending requests
	HTTPClient http.Client
	// routes for sending requests to 01cloud-api-server
	Routes *config.Routes

	UserAuthToken string
	UserWebToken  string
	UserEmail     string
	WebTokenUsed  bool

	// default browser of the user
	UserBrowser *browser.Browser
	Promopter   survey.Prompt

	Term terminal.Provider
	IO   *iostreams.IOStreams

	Debug      printer.DebugInterface
	Printer    printer.PrinterInterface
	Auth0Token string
}

type FactoryInterface interface {
	createDefaultConfig()
	CleanUp()
	CleanUpTest()
}

type MyConfig struct {
	Auth0Token string `yaml:"auth0_token"`
}

// New
//
// creates a new factory instance.
// which will contain all the default values for the necessary fields
func New(ctx context.Context, conf *config.Config) *Factory {
	f := &Factory{
		Config:     conf,
		Ctx:        ctx,
		UserOS:     runtime.GOOS,
		HTTPClient: http.Client{},
		Routes:     config.Load(),
		Term:       terminal.New(runtime.GOOS),
		IO:         iostreams.System(),
		Printer:    printer.New(),
	}

	if f.Config == nil {
		f.createDefaultConfig()
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// as there was error getting the user home directory
		// change the config folder path to temp
		// as nor then, none of the files can be created at runtime
		return f
	}
	f.UserHomeFolder = homeDir
	f.UserBrowser = browser.New("", cliBrowser.Stdout, cliBrowser.Stderr)
	f.Debug = printer.NewDebug(f.Config.Populated.Debug)

	return f
}

func (f *Factory) createDefaultConfig() {
	f.Config = config.New()
	f.Config.ConfigFolder = "/tmp"
	f.Config.AuthFile = "/auth.yaml"
}

func (f *Factory) CleanUp() {
	err := os.RemoveAll(f.Config.ConfigFolder)
	if err != nil {
		f.Printer.Errorf("Err: %s", err.Error())
	}
}

func (f *Factory) CleanUpTest() {
	err := os.RemoveAll(f.Config.ConfigFolder)
	if err != nil {
		f.Debug.Debug(err)
	}
}

// loadAuthFromDisk tries to populate Factory fields from ~/.01cloud/token.yaml
func (f *Factory) loadAuthFromDisk() error {
	path := f.Config.ConfigFolder + f.Config.AuthFile // e.g. /home/ash/.01cloud/token.yaml
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Only the fields we need to build headers
	var cfg struct {
		Email      string `yaml:"email"`
		AuthToken  string `yaml:"token"`
		WebToken   string `yaml:"xpersonaltoken"`
		Auth0Token string `yaml:"auth0_token"`
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}

	if cfg.Auth0Token != "" {
		f.Auth0Token = cfg.Auth0Token
	}
	if cfg.AuthToken != "" {
		f.UserAuthToken = cfg.AuthToken
	}
	if cfg.Email != "" {
		f.UserEmail = cfg.Email
	}
	return nil
}

func (f *Factory) GetAuth() map[string]string {
	f.WebTokenUsed = false

	// Ensure tokens are loaded from disk if not already present in Factory
	if f.Auth0Token == "" || f.UserAuthToken == "" {
		_ = f.loadAuthFromDisk() // ignore error; downstream calls will handle missing auth
	}

	header := map[string]string{
		"Authorization": "Basic " + f.UserAuthToken,
		"X-CUSTOM-AUTH": f.Auth0Token,
	}

	return header
}
