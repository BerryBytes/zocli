package login

import (
	"fmt"
	"log"
	"strings"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
)

func (l *Opts) credsFlow() error {
	err := l.checkIfAlreadyLoggedIn()
	if err != nil {
		return err
	}
	l.Email = l.F.Config.Populated.Email
	l.Password = l.F.Config.Populated.Password
	err = l.askCreds()
	if err != nil {
		log.Fatal(2, err)
	}
	return nil
}

// // askCreds
// //
// // prompts the user to enter the email and password
func (l *Opts) askCreds() error {
	if l.Email == "" {
		fmt.Print("Email: ")
		text, err := l.F.Term.ReadInput()
		if err != nil {
			return err
		}
		l.Email = strings.TrimSpace(strings.Trim(text, "\n"))
	}

	if l.Password == "" {
		fmt.Print("Password: ")
		bytePass, err := l.F.Term.ReadPassword()
		if err != nil {
			return err
		}
		fmt.Println()
		l.Password = bytePass
	}

	return l.LoginWithCreds()
}

// LoginWithCreds
//
// try logging the user in with the credentials provided
func (l *Opts) LoginWithCreds() error {
	body := []byte(`{"email": "` + l.Email + `", "password": "` + l.Password + `"}`)
	reqConfig := defaults.BasicLogin(l.F, map[string]interface{}{"body": body})

	baseRes := reqConfig.Request()
	if baseRes == nil {
		return nil
	}
	var loginRes api.LoginResponse
	err := loginRes.FromJson(baseRes.Data)
	if err != nil {
		l.F.Printer.Fatal(9, err)
	}
	l.LoginResponse = &loginRes
	return l.LoggedIn()
}
