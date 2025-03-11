package fs

import (
	"context"
	"errors"
	"os"

	"github.com/berrybytes/zocli/pkg/cmd/auth/login"
	"github.com/berrybytes/zocli/pkg/utils/factory"

	"gopkg.in/yaml.v3"
)

// CheckIsLoggedIn
//
// checks if a user is logged in by reading the
// authentication file and setting the necessary
// properties in the factory object.
// The function takes a pointer to a factory object (f)
// which is assumed to be properly initialized with the required properties.
//
// Parameters:
// - f: A pointer to a factory object that provides the necessary properties for authentication.
func CheckIsLoggedIn(f *factory.Factory) {
	if f.Ctx == nil {
		f.Ctx = context.Background()
	}
	_, cancel := context.WithCancel(f.Ctx)
	defer cancel()

	authFile := f.Config.ConfigFolder + f.Config.AuthFile
	if _, err := os.Stat(authFile); err != nil && errors.Is(err, os.ErrNotExist) {
		return
	}

	file, err := os.ReadFile(authFile)
	if err != nil {
		return
	}

	var saveDetails login.SaveConfig
	err = yaml.Unmarshal(file, &saveDetails)
	if err != nil {
		return
	}

	f.Debug.Debugf("seems like valid configuration for authentication is found with email: %s, auth token: %s and web token: %s", saveDetails.Email, saveDetails.AuthToken, saveDetails.WebToken)
	f.LoggedIn = true
	f.UserAuthToken = saveDetails.AuthToken
	f.UserWebToken = saveDetails.WebToken
	f.UserEmail = saveDetails.Email
}
