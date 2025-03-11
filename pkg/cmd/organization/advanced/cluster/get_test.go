package cluster

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	"github.com/stretchr/testify/assert"
)

func Test_NewClusterGetCommand(t *testing.T) {
	f := mock_factory.NewFactory()
	o := Opts{F: f}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewClusterGetCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"g", "ge", "list", "lis", "retrieve"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("OutputFlag", func(t *testing.T) {
		cmd := NewClusterGetCommand(&o)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewClusterGetCommand(&o)

		assert.Equal(t, "get", cmd.Use, "Expected command name to be 'get'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewClusterGetCommand(&o)

		flag := cmd.Flag("out")
		assert.NotNil(t, flag, "Expected 'out' flag to be attached to the command")
	})
	tt := []struct {
		name   string
		args   []string
		out    string
		err    string
		exists bool
	}{
		{
			name: "help",
			out:  grammar.ClusterGetHelp,
			args: []string{"--help"},
			err:  "",
		},
		{
			name: "random as flag input",
			out:  "",
			args: []string{"--random"},
			err:  "unknown flag: --random",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := mock_factory.NewFactory()
			o := Opts{F: f}
			f.LoggedIn = false
			getCmd := NewClusterGetCommand(&o)

			getCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			getCmd.SetIn(stdin)
			getCmd.SetOut(stderr)
			getCmd.SetErr(stderr)

			_, err := getCmd.ExecuteC()
			fmt.Println(err)
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
