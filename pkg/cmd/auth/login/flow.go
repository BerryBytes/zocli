package login

import (
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func (l *Opts) flow() error {
	switch {
	case l.WithBrowser:
		return l.browserFlow()

	default:
		return l.ask()
	}
}

func (l *Opts) ask() error {
	options := []string{
		"SSO : Login with browser using 01Cloud SSO (Single Sign On) Token",
		"Help : Print the help menu",
	}
	var qs = []*survey.Question{
		{
			Name: "login",
			Prompt: &survey.Select{
				Message: "Choose one of the following to login: ",
				Options: options,
				Default: options[0],
			},
		},
	}
	// the answers will be written to this struct
	answers := struct {
		LoginOption string `survey:"login"` // or you can tag fields to match a specific name
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		if err.Error() == "interrupt" {
			os.Exit(0)
		}
		return err
	}

	switch {
	case strings.Contains(answers.LoginOption, "SSO"):
		return l.browserFlow()
	case strings.Contains(answers.LoginOption, "H"):
		return l.cmd.Help()
	default:
		return l.ask()
	}
}
