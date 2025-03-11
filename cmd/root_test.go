package cmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	cleanFactory := factory.New(context.Background(), config.New())

	verboseFactory := factory.New(context.Background(), config.New())
	verboseFactory.Verbose = true

	noInteractiveFactory := factory.New(context.Background(), config.New())
	noInteractiveFactory.NoInteractive = true

	quietFactory := factory.New(context.Background(), config.New())
	quietFactory.Quiet = true

	noConfigFactory := factory.New(context.Background(), config.New())
	noConfigFactory.ConfigCreated = false

	invalidConfigFactory := factory.New(context.Background(), config.New())
	invalidConfigFactory.ConfigCreated = false
	invalidConfigFactory.Config.ConfigFolder = "/root/invalid/folder"

	tt := []struct {
		name string
		args []string
		out  string
		err  string
		f    *factory.Factory
	}{
		{
			name: "no args",
			args: []string{},
			out:  grammar.RootHelp,
			err:  "",
			f:    cleanFactory,
		},
		{
			name: "help",
			out:  grammar.RootHelp,
			args: []string{"--help"},
			err:  "",
			f:    cleanFactory,
		},
		{
			name: "random input",
			out:  "",
			args: []string{"--random input"},
			err:  "unknown flag: --random input",
			f:    cleanFactory,
		},
		{
			name: "check verbose variable long hand",
			out:  "",
			args: []string{"--verbose"},
			err:  "",
			f:    verboseFactory,
		},
		{
			name: "check verbose variable short hand",
			out:  "",
			args: []string{"-v"},
			err:  "",
			f:    verboseFactory,
		},
		{
			name: "check quiet variable long hand",
			out:  "",
			args: []string{"--quiet"},
			err:  "",
			f:    quietFactory,
		},
		{
			name: "check quiet variable short hand",
			out:  "",
			args: []string{"-q"},
			err:  "",
			f:    quietFactory,
		},
		{
			name: "check no-interactive variable long hand",
			out:  "",
			args: []string{"--no-interactive"},
			err:  "",
			f:    noInteractiveFactory,
		},
		{
			name: "no config created should trigger go-routine",
			out:  "",
			args: []string{},
			err:  "",
			f:    noConfigFactory,
		},
		{
			name: "invalid config folder to be created",
			out:  "",
			args: []string{},
			err:  "",
			f:    invalidConfigFactory,
		},
	}

	root := Cmd

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := tc.f

			root.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			root.SetIn(stdin)
			root.SetOut(stderr)

			err := Execute(f)
			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != "" {
				assert.Contains(t, stderr.String(), tc.out)
			} else if tc.out == "" && tc.err == "" {
				assert.Equal(t, f, tc.f)
			} else if tc.err == "" {
				assert.Empty(t, stderr.String())
			}

			root.ResetFlags()
		})
	}
}
