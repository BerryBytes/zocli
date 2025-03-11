package factory

import (
	"context"
	"testing"

	"github.com/berrybytes/zocli/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestFactoryCreation(t *testing.T) {
	tt := []struct {
		name           string
		ctx            context.Context
		conf           *config.Config
		err            string
		expectedConfig *config.Config
	}{
		{
			name: "creation with nil",
			ctx:  nil,
			conf: nil,
			err:  "",
			expectedConfig: &config.Config{
				ConfigFolder:  "/tmp",
				AuthFile:      "/auth.yaml",
				Populated:     config.ConfigPopulated{},
				ContextFile:   "/contexts.yaml",
				ConfPopulated: true,
			},
		},
		{
			name:           "invalid value",
			ctx:            context.TODO(),
			conf:           config.New(),
			err:            "",
			expectedConfig: config.New(),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := New(tc.ctx, tc.conf)
			assert.Equal(t, tc.expectedConfig, f.Config)
		})
	}
}
