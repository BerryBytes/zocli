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

func Test_NewEnvironmentDeleteCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewEnvironmentDeleteCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"d", "del", "de", "dele", "delete"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewEnvironmentDeleteCommand(&o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewEnvironmentDeleteCommand(&o)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})
	t.Run("CommandName", func(t *testing.T) {
		cmd := NewEnvironmentDeleteCommand(&o)

		assert.Equal(t, "delete", cmd.Use, "Expected command name to be 'delete'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewEnvironmentDeleteCommand(&o)

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
			out:  grammar.EnvironmentDeleteHelp,
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
			deleteCmd := NewEnvironmentDeleteCommand(&o)

			deleteCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			deleteCmd.SetIn(stdin)
			deleteCmd.SetOut(stderr)
			deleteCmd.SetErr(stderr)

			_, err := deleteCmd.ExecuteC()
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

func Test_DeleteRunner(t *testing.T) {
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
		expected := NewEnvironmentDeleteCommand(&o)
		mustEqual := expected
		interfaceRecorder.EXPECT().DeleteRunner(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})

	t.Run("both id and name", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentDeleteCommand(&o)
		o.EnvID = "123"
		o.EnvName = "abc"
		o.DeleteRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentDeleteCommand(&o)
		o.EnvName = "abc"
		recorder.EXPECT().Exit(0)
		recorder.EXPECT().Println("Please use id for now, as this is still under development")

		o.EnvName = "testProj"
		o.DeleteRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by id", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		o.EnvID = "1"
		interfaceRecorder.EXPECT().DeleteEnvByID().Times(1)

		o.DeleteRunner(&cobra.Command{}, []string{})
	})
}

func Test_DeleteEnvByID(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.DeleteEnvironment

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()

	interfaceRecorder.EXPECT().AskDeleteConfirmation()
	recorder.EXPECT().Print(f.IO.ColorScheme().SuccessIcon())
	recorder.EXPECT().Exit(0)
	recorder.EXPECT().Print("Successfully deleted environment\n")
	o := Opts{F: f, I: interfaceRecorder}
	o.EnvID = "1"
	o.DeleteEnvByID()
}
