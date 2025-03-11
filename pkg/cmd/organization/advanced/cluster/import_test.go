package cluster

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	"github.com/stretchr/testify/assert"
)

func Test_NewClusterImportCommand(t *testing.T) {
	f := mock_factory.NewFactory()
	o := Opts{F: f}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"i", "im", "imp", "impo", "impor", "copy"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("ProviderFlag", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)
		output, _ := cmd.Flags().GetString("provider")

		assert.Equal(t, "", output, "Expected default provider to be an empty string")
	})

	t.Run("nameFlag", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)
		output, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", output, "Expected default cluster name to be an empty string")
	})

	t.Run("regionFlag", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)
		output, _ := cmd.Flags().GetString("region")

		assert.Equal(t, "", output, "Expected default region name to be an empty string")
	})

	t.Run("zoneFlag", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)
		output, _ := cmd.Flags().GetString("zone")

		assert.Equal(t, "", output, "Expected default zone name to be an empty string")
	})

	t.Run("labelsFlag", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)
		output, _ := cmd.Flags().GetStringArray("labels")

		assert.Equal(t, []string{}, output, "Expected default labels to be an empty")
	})

	t.Run("mainFileFlag", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)
		output, _ := cmd.Flags().GetStringArray("mainfile")

		assert.Equal(t, []string{}, output, "Expected default mainfile to be an empty string")
	})

	t.Run("providerFileFlag", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)
		output, _ := cmd.Flags().GetStringArray("providerfile")

		assert.Equal(t, []string{}, output, "Expected default providerfile to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)

		assert.Equal(t, "import", cmd.Use, "Expected command name to be 'get'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewClusterImportCommand(&o)

		flag := cmd.Flag("provider")
		assert.NotNil(t, flag, "Expected 'provider' flag to be attached to the command")

		flag = cmd.Flag("name")
		assert.NotNil(t, flag, "Expected 'name' flag to be attached to the command")

		flag = cmd.Flag("region")
		assert.NotNil(t, flag, "Expected 'region' flag to be attached to the command")

		flag = cmd.Flag("zone")
		assert.NotNil(t, flag, "Expected 'zone' flag to be attached to the command")

		flag = cmd.Flag("labels")
		assert.NotNil(t, flag, "Expected 'labels' flag to be attached to the command")

		flag = cmd.Flag("mainfile")
		assert.NotNil(t, flag, "Expected 'mainfile' flag to be attached to the command")

		flag = cmd.Flag("providerfile")
		assert.NotNil(t, flag, "Expected 'providerfile' flag to be attached to the command")
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
			out:  grammar.ClusterImportHelp,
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
			getCmd := NewClusterImportCommand(&o)

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
