package login

import (
	"fmt"
	"strings"
	"time"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
)

// browserFlow
//
// this function is responsible for fetching a one-time code from the server
// and allowing the user to authenticate using the same code
//
// Flow:
//
//	GET one-time code from the server
//	Ask the user to enter the code on the browser (on the tab which just opened)
//		If the browser cannot be opened then, the user must be able to enter the
//		code on a default url and should be authenticated
//	Continuously fetch status of the code from the server to check,
//	if any user logged in using the same code
//	If yes, then fetch the code and details of the user
func (l *Opts) browserFlow() error {
	err := l.checkIfAlreadyLoggedIn()
	if err != nil {
		return err
	}
	sso, err := l.requestSSOCode()
	if err != nil {
		return err
	}

	if sso.Code == "" {
		l.F.Printer.Exit(2)
	}

	l.SsoCode = sso.Code
	l.F.Printer.Print("Login in using the browser using a dynamically generated code.")
	l.F.Printer.Printf("\n! First copy your one-time code: %s\n", sso.Code)
	l.F.Printer.Print("Press Enter to open " + l.F.Routes.GetRoute("device-login") + " in your browser...")
	_, err = l.F.Term.ReadInput()
	if err != nil {
		l.F.Printer.Error(err)
	}

	err = l.F.UserBrowser.Browse(l.F.Routes.GetRoute("device-login"))
	if err != nil {
		fmt.Println("Unable to open the browser.")
		fmt.Println("open the below URL manually on the browser")
		fmt.Println(l.F.Routes.GetRoute("device-login"))
		time.Sleep(1 * time.Second)
	}

	l.checkSSOCodeStatus()
	return err
}

// checkSSOCodeStatus
//
// this function is responsible for fetching the status of the sso code, and
// the time interval between the requests is 1 second
func (l *Opts) checkSSOCodeStatus() {
	time.Sleep(1 * time.Second)
	l.F.Printer.Print("Waiting for code status.")
	for {
		time.Sleep(1 * time.Second)
		l.F.Printer.Print(".")

		body := []byte(`{"code": "` + l.SsoCode + `"}`)
		reqConf := defaults.RequestSSOStatus(l.F, map[string]interface{}{"body": body})

		baseRes := reqConf.Request()
		if baseRes.Message != "Success" {
			time.Sleep(4 * time.Second)
			continue
		}

		l.gotCodeStatus(baseRes)
		break
	}
}

// gotCodeStatus
//
// this function is called after the sso code status returns the token as response.
func (l *Opts) gotCodeStatus(baseRes *api.BaseResponse) {
	var data api.Data
	err := data.FromJson(baseRes.Data)
	if err != nil {
		l.F.Printer.Fatal(9, err.Error())
	}

	headers := map[string]string{"Authorization": "basic " + data.AuthToken}
	reqConf := defaults.Profile(l.F, map[string]interface{}{"headers": headers})

	res := reqConf.Request()
	var profile api.ProfileResponse
	err = profile.FromJson(res.Data)
	if err != nil {
		l.F.Printer.Fatal(9, err)
	}

	l.LoginResponse = &api.LoginResponse{
		Data: api.Data{
			User: api.User{
				Id:    profile.User.Id,
				Email: profile.User.Email,
			},
		},
	}

	// switch the organization, so we can get new token
	reqConf = defaults.SwitchOrganization(l.F, map[string]interface{}{"headers": headers})
	reqConf.URL = strings.ReplaceAll(reqConf.URL, "<:id>", "0")
	res = reqConf.Request()
	var orgSwitch api.OrganizationSwitch
	err = orgSwitch.FromJson(res.Data)
	l.LoginResponse.AuthToken = orgSwitch.Token
	if err != nil {
		l.F.Printer.Fatal(9, "cannot unmarshal")
		return
	}

	// now save the details of the user as the token has also been fetched
	err = l.saveDetails()
	if err != nil {
		l.F.Printer.Fatal(10, err)
		return
	}
}

// requestSSOCode
//
// this function si responsible for fetching a unique SSO code for the new login
// NOTE: as the code gets created, if the user exists this application using CTRL + C
// we must intercept the keybindings and distroy the code by sending request to the server
func (l *Opts) requestSSOCode() (*api.SSOCode, error) {
	reqConfig := defaults.RequestSSO(l.F)

	baseRes := reqConfig.Request()
	var sso api.SSOCode
	err := sso.FromJson(baseRes.Data)
	if err != nil {
		l.F.Printer.Fatal(9, err)
	}
	return &sso, nil
}
