package fs

import (
	"context"
	"testing"

	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/printer"

	"github.com/stretchr/testify/assert"
)

func TestCheckIsLoggedIn(t *testing.T) {
	f := &factory.Factory{Ctx: context.Background(), Config: config.New(), Printer: printer.New(), Debug: printer.NewDebug(false)}
	CheckIsLoggedIn(f)

	tt := []struct {
		name             string
		ctx              context.Context
		conf             *config.Config
		err              string
		expectedLoggedIn bool
	}{
		{
			name:             "creation with nil",
			ctx:              nil,
			conf:             nil,
			err:              "",
			expectedLoggedIn: false,
		},
		{
			name:             "valid config",
			conf:             config.New(),
			ctx:              context.Background(),
			err:              "",
			expectedLoggedIn: f.LoggedIn,
		},
		{
			name: "invalid config file and folder",
			conf: &config.Config{
				ConfigFolder: "/tmp",
				AuthFile:     "/auth.yaml",
			},
			ctx:              context.Background(),
			err:              "",
			expectedLoggedIn: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := factory.New(tc.ctx, tc.conf)
			CheckIsLoggedIn(f)
			assert.Equal(t, tc.expectedLoggedIn, f.LoggedIn)
		})
	}
}
