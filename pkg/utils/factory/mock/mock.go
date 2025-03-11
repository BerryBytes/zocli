package mock_factory

import (
	"context"
	"errors"
	"os"
	"runtime"

	"github.com/berrybytes/zocli/internal/browser"
	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/iostreams"
	"github.com/berrybytes/zocli/pkg/utils/printer"
	"github.com/berrybytes/zocli/pkg/utils/terminal/mock"
	ghBrowser "github.com/cli/browser"
)

func NewFactory() *factory.Factory {
	home := "/tmp"
	f := &factory.Factory{
		UserOS: runtime.GOOS,
		Config: &config.Config{
			ConfigFolder:  home + "/.01Cloud",
			AuthFile:      "/auth.yaml",
			ActiveContext: &config.Context{},
			ContextFile:   "/contexts.yaml",
		},
		Ctx:            context.Background(),
		UserHomeFolder: home,
		ConfigCreated:  false,
		Routes:         config.Load(),
		UserBrowser:    browser.New("", ghBrowser.Stdout, ghBrowser.Stderr),
		Term:           mock.NewMockProvider(),
		Debug:          printer.NewDebug(false),
	}

	testIO, _, _, _ := iostreams.Test()
	f.IO = testIO
	f.IO.SetStdoutTTY(false)
	f.IO.SetStderrTTY(false)
	f.IO.SetStdinTTY(false)

	// make the temporary folder
	if _, err := os.Stat(f.Config.ConfigFolder); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.LoggedIn = false
			err = os.Mkdir(f.Config.ConfigFolder, os.ModePerm)
			if err != nil {
				f.Debug.Debugf("error while creating temp folder. Err: %s", err.Error())
			}
		}
	}

	return f
}
