package fs

import (
	"errors"
	"os"

	"github.com/berrybytes/zocli/pkg/utils/factory"
)

// CheckConfigDir
//
// check if the config directory is present and if not create it
func CheckConfigDir(f *factory.Factory) {
	if _, err := os.Stat(f.Config.ConfigFolder); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.LoggedIn = false
			err = os.Mkdir(f.Config.ConfigFolder, os.ModePerm)
			if err != nil {
				f.Printer.Fatal(4, err)
			}
		}
	}
}
