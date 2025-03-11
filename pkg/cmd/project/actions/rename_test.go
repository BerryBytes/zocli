package actions

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	terminal_mock "github.com/berrybytes/zocli/pkg/utils/terminal/mock"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRenameProjectCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := mock_factory.NewFactory()
		cmd := NewProjectRenameCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"r", "rena", "renam", "renm", "re"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectRenameCommand(f)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectRenameCommand(f)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectRenameCommand(f)

		assert.Equal(t, "rename", cmd.Use, "Expected command name to be 'rename'")
	})

	t.Run("CommandFlags", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectRenameCommand(f)

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
			out:  grammar.ProjectRenameHelp,
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
			enableCmd := NewProjectRenameCommand(f)

			enableCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			enableCmd.SetIn(stdin)
			enableCmd.SetOut(stderr)
			enableCmd.SetErr(stderr)

			_, err := enableCmd.ExecuteC()
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

func TestRename(t *testing.T) {
	requester.Client = &mock.Client{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	requester.Client = &mock.Client{}

	tests := []struct {
		name        string
		expect      func(recorder *mock_printer.MockPrinterInterface)
		id          string
		projectName string
		getDoFunc   func(req *http.Request) (*http.Response, error)
	}{
		{
			name: "Invalid ID",
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				recorder.EXPECT().Fatal(1, "invalid id supplied").Times(1)
				recorder.EXPECT().Fatal(8, "invalid id").Times(1)
			},
			id:          "123abc",
			projectName: "abc",
			getDoFunc:   mock.RenameProject,
		},
		{
			name: "Success",
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				recorder.EXPECT().Exit(0).Times(1)
			},
			id:          "1",
			projectName: "test2Proj",
			getDoFunc:   mock.RenameProject,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			f := mock_factory.NewFactory()
			f.Printer = recorder
			test.expect(recorder)
			if test.name == "Success" {
				recorder.EXPECT().Print(f.IO.ColorScheme().SuccessIcon(), " Successfully name changed to 'test2Proj'").Times(1)
			}
			mock.GetDoFunc = mock.GetSingleProject
			go func() {
				time.Sleep(3) // nolint:staticcheck
				mock.GetDoFunc = test.getDoFunc
			}()

			opts := &Opts{F: f}
			opts.ID = test.id
			opts.Name = test.projectName
			opts.F.LoggedIn = true
			opts.rename(&cobra.Command{}, []string{})
		})
	}
}

func TestAskName(t *testing.T) {
	o := &Opts{F: mock_factory.NewFactory()}
	tests := []struct {
		name     string
		NameVal  string
		expected string
	}{
		{
			name:     "no name supplied",
			NameVal:  "",
			expected: "mock.email@gmail.com",
		},
		{
			name:     "name supplied",
			NameVal:  "test",
			expected: "test",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.NameVal != "" {
				o.Name = test.NameVal
			}
			o.askName()
			assert.Equal(t, test.expected, o.Name)
		})
	}
}

func TestAskID(t *testing.T) {
	o := &Opts{F: mock_factory.NewFactory()}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	o.F.Printer = recorder
	o.F.Term = &terminal_mock.MockCredentialProvider{}
	recorder.EXPECT().Fatal(1, "invalid id supplied").Times(1)

	tests := []struct {
		name     string
		IDVal    string
		expected string
	}{
		{
			name:     "invalid id supplied",
			IDVal:    "",
			expected: "mock.email@gmail.com",
		},
		{
			name:     "id supplied",
			IDVal:    "123",
			expected: "123",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.IDVal != "" {
				o.ID = test.IDVal
			}
			o.askID()
			assert.Equal(t, test.expected, o.ID)
		})
	}
}
