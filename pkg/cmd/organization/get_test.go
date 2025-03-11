package organization

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	interfaceMock "github.com/berrybytes/zocli/pkg/cmd/organization/mock"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_NewOrganizationGetCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewOrganizationGetCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"g", "ge", "list", "lis", "retrieve"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("OutputFlag", func(t *testing.T) {
		cmd := NewOrganizationGetCommand(&o)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	t.Run("WideFlag", func(t *testing.T) {
		cmd := NewOrganizationGetCommand(&o)
		wide, _ := cmd.Flags().GetBool("wide")

		assert.Equal(t, false, wide, "Expected default wide to be false")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewOrganizationGetCommand(&o)

		assert.Equal(t, "get", cmd.Use, "Expected command name to be 'get'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewOrganizationGetCommand(&o)

		flag := cmd.Flag("wide")
		assert.NotNil(t, flag, "Expected 'wide' flag to be attached to the command")

		flag = cmd.Flag("out")
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
			out:  grammar.OrganizationGetHelp,
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
			getCmd := NewOrganizationGetCommand(&o)

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
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()

	t.Run("command initialization with prerunner", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		expected := NewOrganizationGetCommand(&o)
		mustEqual := expected
		interfaceRecorder.EXPECT().GetRunner(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})

	t.Run("nothing supplied", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetOrganizations().Times(1)
		interfaceRecorder.EXPECT().PrintOrganizationTable().Times(1)
		recorder.EXPECT().Exit(0).Times(1)
		o.GetRunner(&cobra.Command{}, []string{})
		assert.Equal(t, Opts{F: f, I: interfaceRecorder}, o)
	})
}
