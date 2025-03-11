package apply

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	manifestprocessor "github.com/berrybytes/zocli/pkg/utils/manifestProcessor"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_NewApplyManifestCommand(t *testing.T) {
	f := &factory.Factory{}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewApplyManifestCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"app", "appl", "ap"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("FileFlag", func(t *testing.T) {
		cmd := NewApplyManifestCommand(f)
		output, _ := cmd.Flags().GetString("file")

		assert.Equal(t, "", output, "Expected default file to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewApplyManifestCommand(f)

		assert.Equal(t, "apply", cmd.Use, "Expected command name to be 'apply'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewApplyManifestCommand(f)

		flag := cmd.Flag("file")
		assert.NotNil(t, flag, "Expected 'file' flag to be attached to the command")
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
			out:  grammar.ManifestHelp,
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
			f.LoggedIn = false
			getCmd := NewApplyManifestCommand(f)

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

func Test_ManifestRunner(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	t.Run("invalid file", func(t *testing.T) {
		o := Opts{F: f}
		o.Processor = manifestprocessor.New(f)
		_ = NewApplyManifestCommand(f)
		o.File = "../.."
		recorder.EXPECT().Fatal(10, "cannot access file")
		o.ManifestRunner(&cobra.Command{}, []string{})
	})

	t.Run("nothing supplied", func(t *testing.T) {
		o := Opts{F: f}
		o.ManifestRunner(&cobra.Command{}, []string{})
		assert.Equal(t, Opts{F: f}, o)
	})
}
