package packages

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	interfaceMock "github.com/berrybytes/zocli/pkg/cmd/organization/advanced/cluster/packages/mock"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewUninstallCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewPackageUnInstallCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"u", "ui", "un", "unin", "unins", "uninst", "uninsta", "uninstal", "uin", "uins", "uinst", "uinsta", "uinstal", "remove", "r"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IDFlag", func(t *testing.T) {
		cmd := NewPackageUnInstallCommand(&o)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be \"\"")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewPackageUnInstallCommand(&o)

		assert.Equal(t, "uninstall", cmd.Use, "Expected command name to be 'uninstall'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewPackageUnInstallCommand(&o)

		flag := cmd.Flag("name")
		assert.NotNil(t, flag, "Expected 'name' flag to be attached to the command")
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
			out:  grammar.PackagesUnInstallHelp,
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
			interfaceCtrl := gomock.NewController(t)
			interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
			defer interfaceCtrl.Finish()
			o := Opts{F: f, I: interfaceRecorder}
			f.LoggedIn = false
			getCmd := NewPackageUnInstallCommand(&o)

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
