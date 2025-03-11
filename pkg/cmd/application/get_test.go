package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	interfaceMock "github.com/berrybytes/zocli/pkg/cmd/application/mock"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	mock "github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestNewProjectApplicationsGetCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewApplicationGetCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"g", "ge", "list", "lis", "retrieve"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewApplicationGetCommand(&o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewApplicationGetCommand(&o)
		name, _ := cmd.Flags().GetString("pname")

		assert.Equal(t, "", name, "Expected default pname to be an empty string")
	})

	t.Run("OutputFlag", func(t *testing.T) {
		cmd := NewApplicationGetCommand(&o)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewApplicationGetCommand(&o)

		assert.Equal(t, "get", cmd.Use, "Expected command name to be 'get'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewApplicationGetCommand(&o)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("pname")
		assert.NotNil(t, flag, "Expected 'pname' flag to be attached to the command")

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
			out:  grammar.ApplicationGetHelp,
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
			getCmd := NewApplicationGetCommand(&o)

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

func Test_GetApplication(t *testing.T) {
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
		expected := NewApplicationGetCommand(&o)
		mustEqual := expected
		interfaceRecorder.EXPECT().GetRunner(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})
	t.Run("both id and name", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewApplicationGetCommand(&o)
		o.ProjectID = "123"
		o.ProjectName = "abc"
		o.GetRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by id", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetApps().Times(1)
		interfaceRecorder.EXPECT().PrintApps().Times(1)
		_ = NewApplicationGetCommand(&o)
		o.ProjectID = "1"
		recorder.EXPECT().Exit(0).Times(1)
		o.GetRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetProjectDetailByName().Times(1).Return(&api.Project{ID: 1}).Times(1)
		interfaceRecorder.EXPECT().GetApps().Times(1)
		interfaceRecorder.EXPECT().PrintApps().Times(1)
		_ = NewApplicationGetCommand(&o)

		recorder.EXPECT().Exit(0).Times(1)
		o.ProjectName = "testProj"
		o.GetRunner(&cobra.Command{}, []string{})
	})

	t.Run("nothing supplied", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		o.GetRunner(&cobra.Command{}, []string{})
		assert.Equal(t, Opts{F: f, I: interfaceRecorder}, o)
	})
}

func Test_GetApps(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.ListApps

	o := Opts{F: f, I: interfaceRecorder, ProjectID: "1"}
	o.GetApps()
}

func Test_PrintApps(t *testing.T) {
	_ = []struct {
		name   string
		supply []api.Application
		expect func(recorder *mock_printer.MockPrinterInterface)
		output string
	}{
		{
			name: "single var JSON", supply: []api.Application{
				{Id: 1, Name: "test"},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res, _ := json.Marshal([]api.ApplicationPresenter{
					{Id: 1, Name: "test"},
				}[0])
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "json",
		},
		{
			name: "multiple var JSON", supply: []api.Application{
				{Id: 1, Name: "test"},
				{Id: 2, Name: "test"},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res, _ := json.Marshal(api.ApplicationPresenterList{
					Applications: []api.ApplicationPresenter{
						{Id: 1, Name: "test"},
						{Id: 2, Name: "test"},
					},
				})
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "json",
		},
		{
			name: "multiple var YAML", supply: []api.Application{
				{Id: 1, Name: "test"},
				{Id: 2, Name: "test"},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res := api.ApplicationPresenterList{
					Applications: []api.ApplicationPresenter{
						{Id: 1, Name: "test"},
						{Id: 2, Name: "test"},
					},
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
	// for _, test := range tests {
	// 	t.Run(test.name, func(t *testing.T) {
	// 		o := &Opts{F: mock_factory.NewFactory(), Output: test.output}
	// 		o.List = api.ApplicationList{Applications: test.supply}
	// 		ctrl := gomock.NewController(t)
	// 		defer ctrl.Finish()
	// 		recorder := mock_printer.NewMockPrinterInterface(ctrl)
	// 		o.F.Printer = recorder
	// 		test.expect(recorder)
	// 		o.PrintApps()
	// 	})
	// }

	t.Run("no output supplied", func(t *testing.T) {
		o := Opts{F: mock_factory.NewFactory()}
		o.PrintApps()
	})
}
