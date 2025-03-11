package environment

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	interfaceMock "github.com/berrybytes/zocli/pkg/cmd/environment/mock"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_NewEnvironmentStopCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewEnvironmentStopCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"s", "st", "sto", "halt", "hal"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewEnvironmentStopCommand(&o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewEnvironmentStopCommand(&o)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})
	t.Run("CommandName", func(t *testing.T) {
		cmd := NewEnvironmentStopCommand(&o)

		assert.Equal(t, "stop", cmd.Use, "Expected command name to be 'stop'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewEnvironmentStopCommand(&o)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("name")
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
			out:  grammar.EnvironmentStopHelp,
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
			interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
			defer interfaceCtrl.Finish()
			o := Opts{F: f, I: interfaceRecorder}
			f.LoggedIn = false
			stopCmd := NewEnvironmentStopCommand(&o)

			stopCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			stopCmd.SetIn(stdin)
			stopCmd.SetOut(stderr)
			stopCmd.SetErr(stderr)

			_, err := stopCmd.ExecuteC()
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

func Test_StopRunner(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()

	t.Run("command initialization with prerunner", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		expected := NewEnvironmentStopCommand(&o)
		mustEqual := expected
		interfaceRecorder.EXPECT().StopRunner(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})

	t.Run("both id and name", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentStopCommand(&o)
		o.EnvID = "123"
		o.EnvName = "abc"
		o.StopRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentStopCommand(&o)
		o.EnvName = "abc"
		recorder.EXPECT().Exit(0)
		recorder.EXPECT().Println("Please use id for now, as this is still under development")

		o.EnvName = "testProj"
		o.StopRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by id", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		o.EnvID = "1"
		interfaceRecorder.EXPECT().StopEnvByID().Times(1)

		o.StopRunner(&cobra.Command{}, []string{})
	})
}

func Test_StopEnvByID(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.StopEnvironment

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	recorder.EXPECT().Println("Sent command for stopping environment")
	recorder.EXPECT().Exit(0).Times(1)
	o := Opts{F: f}
	o.EnvID = "1"
	o.StopEnvByID()
}
