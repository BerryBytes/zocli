package login

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// checkIfAlreadyLoggedIn checks if the user is already logged in by loading saved details
func (l *Opts) checkIfAlreadyLoggedIn() error {
	// Load saved credentials from YAML file
	if err := l.loadDetails(); err == nil {
		// Check if required credentials are present
		if l.F.UserAuthToken != "" && l.F.Auth0Token != "" && l.F.UserEmail != "" {
			l.F.LoggedIn = true
		}
	}

	if !l.F.LoggedIn {
		return nil
	}

	proceed := false
	prompt := &survey.Confirm{
		Message: "You are already logged in using " + l.F.UserEmail + ". Re-Authenticate?",
		Default: false,
	}

	err := survey.AskOne(prompt, &proceed)
	if err != nil {
		l.F.Printer.Fatalf(5, "cannot proceed.\nErr: %v\n", err)
	}
	if proceed {
		err := os.RemoveAll(l.F.Config.ConfigFolder + l.F.Config.AuthFile)
		if err != nil {
			l.F.Printer.Fatal(1, err)
		}
		// Clear Factory credentials to reset state
		l.F.UserAuthToken = ""
		l.F.UserWebToken = ""
		l.F.UserEmail = ""
		l.F.Auth0Token = ""
		l.F.LoggedIn = false
		l.LoginResponse = nil
		l.Auth0Token = ""
		fmt.Println(l.F.IO.ColorScheme().SuccessIcon(), " Successfully logged out from old session.")
		return nil
	}
	os.Exit(0)
	return nil
}
