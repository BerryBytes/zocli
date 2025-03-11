package get

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/internal/config"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	manifestprocessor "github.com/berrybytes/zocli/pkg/utils/manifestProcessor"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectGetCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectGetCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"g", "ge"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectGetCommand(f)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectGetCommand(f)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("WideFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectGetCommand(f)
		wide, _ := cmd.Flags().GetBool("wide")

		assert.Equal(t, false, wide, "Expected default wide to be false")
	})
	t.Run("OutputFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectGetCommand(f)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	// Test case 6: Check command name
	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectGetCommand(f)

		assert.Equal(t, "get", cmd.Use, "Expected command name to be 'get'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectGetCommand(f)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("name")
		assert.NotNil(t, flag, "Expected 'name' flag to be attached to the command")

		flag = cmd.Flag("wide")
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
			out:  grammar.ProjectGetHelp,
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
			getCmd := NewProjectGetCommand(f)

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

func TestIDAndNameBoth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f := mock_factory.NewFactory()
	f.LoggedIn = true
	f.Printer = recorder
	recorder.EXPECT().Fatal(1, "provide either id or name, but not both").Times(1)

	getCmd := NewProjectGetCommand(f)
	getCmd.SetArgs([]string{"--id", "23198", "--name", "fkjahskdfhaksd"})
	_, _ = getCmd.ExecuteC()
}

func TestWideData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f := mock_factory.NewFactory()
	f.Printer = recorder
	f.LoggedIn = true
	recorder.EXPECT().Exit(0).Times(1)
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.GetSingleProject

	getCmd := NewProjectGetCommand(f)
	getCmd.SetArgs([]string{"--id", "1", "-w"})
	_, _ = getCmd.ExecuteC()
}
func TestInvalidProjectIDFromServer(t *testing.T) {
	tests := []struct {
		name       string
		fatalMsg   string
		normalExit bool
		exitCode   int
		id         string
	}{
		{
			name:     "invalid project id from server",
			fatalMsg: "invalid id",
			exitCode: 8,
			id:       "200",
		},
		{
			name:       "valid project id from server",
			exitCode:   0,
			normalExit: true,
			id:         "1",
		},
		{
			name:     "invalid id supplied",
			fatalMsg: "invalid id supplied",
			exitCode: 1,
			id:       "abc",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			recorder := mock_printer.NewMockPrinterInterface(ctrl)

			requester.Client = &mock.Client{}
			mock.GetDoFunc = mock.GetSingleProject

			f := mock_factory.NewFactory()
			f.Printer = recorder
			f.LoggedIn = true
			if test.normalExit {
				recorder.EXPECT().Exit(test.exitCode).Times(1)
			} else {
				recorder.EXPECT().Fatal(test.exitCode, test.fatalMsg).Times(1)
			}
			if test.id == "abc" {
				recorder.EXPECT().Fatal(8, "invalid id").Times(1)
			}
			getCmd := NewProjectGetCommand(f)
			getCmd.SetArgs([]string{"--id", test.id})
			_, _ = getCmd.ExecuteC()
		})
	}

}

func TestGetAllProjects(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.GetAllProjects

	f := mock_factory.NewFactory()
	f.Printer = recorder
	f.LoggedIn = true
	recorder.EXPECT().Exit(0).Times(1)

	getCmd := NewProjectGetCommand(f)
	_, _ = getCmd.ExecuteC()
}

func TestOpts_PrintProjects(t *testing.T) {
	tests := []struct {
		name   string
		supply []api.Project
		expect func(recorder *mock_printer.MockPrinterInterface)
		output string
	}{
		{
			name: "single var JSON", supply: []api.Project{
				{UserId: 1, Active: true},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				f := factory.New(context.Background(), config.New())
				manifest := manifestprocessor.New(f)
				var a []interface{}
				res, _ := json.Marshal(manifest.MakeManifest("project", []api.Project{
					{UserId: 1, Active: true},
				}[0]))
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "json",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := Opts{Output: test.output, F: mock_factory.NewFactory(), List: api.ProjectList{Projects: test.supply}}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			o.F.Printer = recorder
			test.expect(recorder)
			o.printProjects()
		})
	}
	t.Run("without any output args", func(t *testing.T) {
		o := Opts{F: mock_factory.NewFactory(), List: api.ProjectList{Projects: []api.Project{}}}
		o.printProjects()
	})
}
