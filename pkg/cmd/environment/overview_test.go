package environment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
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
	"gopkg.in/yaml.v3"
)

func Test_NewEnvironmentOverviewCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewEnvironmentOverviewCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"o", "over", "ove", "overv"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewEnvironmentOverviewCommand(&o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewEnvironmentOverviewCommand(&o)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("OutputFlag", func(t *testing.T) {
		cmd := NewEnvironmentOverviewCommand(&o)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewEnvironmentOverviewCommand(&o)

		assert.Equal(t, "overview", cmd.Use, "Expected command name to be 'overview'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewEnvironmentOverviewCommand(&o)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("out")
		assert.NotNil(t, flag, "Expected 'out' flag to be attached to the command")

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
			out:  grammar.EnvironmentOverviewHelp,
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
			getCmd := NewEnvironmentOverviewCommand(&o)

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

func Test_OverviewRunner(t *testing.T) {
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
		expected := NewEnvironmentOverviewCommand(&o)
		mustEqual := expected
		interfaceRecorder.EXPECT().OverviewRunner(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})
	t.Run("both id and name", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentOverviewCommand(&o)
		o.EnvID = "123"
		o.EnvName = "abc"
		o.OverviewRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by id", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentOverviewCommand(&o)
		o.EnvID = "1"
		recorder.EXPECT().Exit(0)
		interfaceRecorder.EXPECT().GetEnvironmentOverview(o.EnvID)
		o.OverviewRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentOverviewCommand(&o)
		o.EnvName = "abc"
		recorder.EXPECT().Exit(0)
		recorder.EXPECT().Println("Please use id for now, as this is still under development")

		o.OverviewRunner(&cobra.Command{}, []string{})
	})

	t.Run("get without anything", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewEnvironmentOverviewCommand(&o)

		o.OverviewRunner(&cobra.Command{}, []string{})
	})
}

func Test_GetEnvironmentOverview(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockEnvironmentInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()

	o := Opts{F: f, I: interfaceRecorder}
	_ = NewEnvironmentOverviewCommand(&o)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.GetEnvironmentOverview

	interfaceRecorder.EXPECT().PrintOverviewTable().Times(1)
	o.GetEnvironmentOverview("1")
}

func Test_PrintOverview(t *testing.T) {
	tests := []struct {
		name   string
		supply api.EnvironmentOverview
		expect func(recorder *mock_printer.MockPrinterInterface)
		output string
	}{
		{
			name:   "single var JSON",
			supply: api.EnvironmentOverview{TotalCPU: 1},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res, _ := json.Marshal(api.EnvironmentOverview{
					TotalCPU: 1,
				})
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "json",
		},
		{
			name:   "multiple var YAML",
			supply: api.EnvironmentOverview{TotalCPU: 1},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res := api.EnvironmentOverview{
					TotalCPU: 1,
				}
				var final bytes.Buffer
				yamlEncoder := yaml.NewEncoder(&final)
				yamlEncoder.SetIndent(2)
				_ = yamlEncoder.Encode(&res)
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "yaml",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := &Opts{F: mock_factory.NewFactory(), Output: test.output}
			o.Overview = &test.supply
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			recorder.EXPECT().Exit(0).Times(1)
			o.F.Printer = recorder
			test.expect(recorder)
			o.PrintOverviewTable()
		})
	}

	t.Run("no output supplied", func(t *testing.T) {
		o := Opts{F: mock_factory.NewFactory()}
		o.Overview = &api.EnvironmentOverview{}
		o.PrintOverviewTable()
	})
}
