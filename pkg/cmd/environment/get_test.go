package environment

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/internal/config"
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

func Test_NewEnvironmentGetCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewEnvironmentGetCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"g", "ge", "list", "lis", "retrieve"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewEnvironmentGetCommand(&o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewEnvironmentGetCommand(&o)
		name, _ := cmd.Flags().GetString("aname")

		assert.Equal(t, "", name, "Expected default aname to be an empty string")
	})

	t.Run("OutputFlag", func(t *testing.T) {
		cmd := NewEnvironmentGetCommand(&o)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewEnvironmentGetCommand(&o)

		assert.Equal(t, "get", cmd.Use, "Expected command name to be 'get'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewEnvironmentGetCommand(&o)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("out")
		assert.NotNil(t, flag, "Expected 'out' flag to be attached to the command")

		flag = cmd.Flag("name")
		assert.NotNil(t, flag, "Expected 'name' flag to be attached to the command")

		flag = cmd.Flag("password")
		assert.NotNil(t, flag, "Expected 'password' flag to be attached to the command")
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
			out:  grammar.EnvironmentGetHelp,
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
			getCmd := NewEnvironmentGetCommand(&o)

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

func Test_GetRunner(t *testing.T) {
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
		expected := NewEnvironmentGetCommand(&o)
		mustEqual := expected
		interfaceRecorder.EXPECT().GetRunner(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})
	t.Run("both id and name", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentGetCommand(&o)
		o.EnvID = "123"
		o.EnvName = "abc"
		o.GetRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentGetCommand(&o)
		o.EnvName = "abc"
		recorder.EXPECT().Exit(0)
		recorder.EXPECT().Println("Please use id for now, as this is still under development")

		o.EnvName = "testProj"
		o.GetRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by context application", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}

		requester.Client = &mock.Client{}
		mock.GetDoFunc = mock.GetAllEnvironment

		recorder.EXPECT().Print(o.F.IO.ColorScheme().SuccessIcon(), " Using context Application Value.\n")
		o.F.Config.ActiveContext = &config.Context{DefaultApplication: 1}
		interfaceRecorder.EXPECT().PrintMultiEnvs().Times(1)
		recorder.EXPECT().Exit(0)
		// recorder.EXPECT().Fatal(9, "cannot unmarshal data")
		_ = NewEnvironmentGetCommand(&o)

		o.GetRunner(&cobra.Command{}, []string{})
	})
}
