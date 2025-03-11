package fs

import (
	"context"
	"testing"

	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/stretchr/testify/assert"
)

func TestCheckConfigDir(t *testing.T) {
	tt := []struct {
		name   string
		ctx    context.Context
		conf   *config.Config
		err    string
		exists bool
	}{
		{
			name:   "valid config",
			conf:   config.New(),
			ctx:    context.Background(),
			err:    "",
			exists: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := factory.New(tc.ctx, tc.conf)

			if tc.conf != nil && tc.err == "" {
				CheckConfigDir(f)
				exists := assert.DirExists(t, tc.conf.ConfigFolder)
				assert.Equal(t, exists, tc.exists)
			}
		})
	}
}
