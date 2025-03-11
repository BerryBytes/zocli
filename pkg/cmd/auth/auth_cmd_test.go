package auth

import (
	"bytes"
	"context"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/stretchr/testify/assert"
)

func TestAuthCmd(t *testing.T) {
	tt := []struct {
		name string
		args []string
		out  string
		err  string
	}{
		{
			name: "no args",
			args: []string{},
			out:  grammar.AuthHelp,
			err:  "",
		},
		{
			name: "help",
			out:  grammar.AuthHelp,
			args: []string{"--help"},
			err:  "",
		},
		{
			name: "random as flag input",
			out:  "",
			args: []string{"--random"},
			err:  "unknown flag: --random",
		},
		{
			name: "random inputs",
			out:  "",
			args: []string{"random input"},
			err:  "unknown command \"random input\" for \"auth\"",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := factory.New(context.Background(), config.New())
			f.LoggedIn = false
			authCmd := NewAuthCommand(f)

			authCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			authCmd.SetIn(stdin)
			authCmd.SetOut(stderr)
			authCmd.SetErr(stderr)

			_, err := authCmd.ExecuteC()
			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != "" {
				assert.Contains(t, stderr.String(), tc.out)
			}
		})
	}
}
