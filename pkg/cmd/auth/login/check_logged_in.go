package login

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func (l *Opts) checkIfAlreadyLoggedIn() error {
	if !l.F.LoggedIn {
		return nil
	}

	proceed := false
	prompt := &survey.Confirm{
		Message: "You are already Logged in using " + l.F.UserEmail + ". Re-Authenticate ?",
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
		fmt.Println(l.F.IO.ColorScheme().SuccessIcon(), " Successfully Logged out from old session.")
		return nil
	}
	os.Exit(0)
	return nil
}
