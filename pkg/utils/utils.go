package utils

import (
	"encoding/json"

	"github.com/AlecAivazis/survey/v2"
	"github.com/berrybytes/zocli/pkg/utils/factory"
)

func ConvertType(input, output interface{}) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, output)
	if err != nil {
		return err
	}
	return nil
}

func ConfirmIfToProceed(message string, f *factory.Factory) {
	proceed := false
	prompt := &survey.Confirm{
		Message: message,
		Default: false,
	}

	err := survey.AskOne(prompt, &proceed)
	if err != nil {
		f.Printer.Fatal(5, "cannot proceed")
		return
	}

	if !proceed {
		f.Printer.Exit(0)
		return
	}
}
