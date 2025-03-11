package login

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"

	"golang.org/x/term"
)

func (l *Opts) webTokenFlow() error {
	err := l.checkIfAlreadyLoggedIn()
	if err != nil {
		return err
	}
	l.WebToken = l.F.Config.Populated.WebToken
	err = l.askToken()
	if err != nil {
		return err
	}
	return nil
}

// askToken
//
// prompts the user to enter the Token
func (l *Opts) askToken() error {
	if l.WebToken != "" {
		return l.LoginWithToken()
	}
	fmt.Print("Token: ")
	var webToken string
	var err error
	fd := int(os.Stdin.Fd())
	if term.IsTerminal(fd) {
		byteToken, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		webToken = string(byteToken)
		fmt.Println() //nolint:forbidigo // generic space for term
	} else {
		reader := bufio.NewReader(os.Stdin)
		webToken, err = reader.ReadString('\n')
		if err != nil {
			return err
		}
	}

	l.WebToken = strings.TrimSpace(webToken)
	return l.LoginWithToken()
}

// LoginWithToken
//
// try logging the user in with the webToken provided
func (l *Opts) LoginWithToken() error {
	headers := map[string]string{"X-PERSONAL-TOKEN": l.WebToken}
	reqConfig := defaults.Profile(l.F, map[string]interface{}{"headers": headers})

	res := reqConfig.Request()
	if res == nil {
		return nil
	}
	var profile api.ProfileResponse
	err := profile.FromJson(res.Data)
	if err != nil {
		l.F.Printer.Fatal(9, err)
	}

	l.LoginResponse = &api.LoginResponse{
		Data: api.Data{
			User: api.User{
				Id:    profile.User.Id,
				Email: profile.User.Email,
			},
			WebToken: l.WebToken,
		},
	}
	err = l.saveDetails()
	if err != nil {
		l.F.Printer.Fatal(10, err)
	}
	return err
	//return l.LoggedIn()
}
