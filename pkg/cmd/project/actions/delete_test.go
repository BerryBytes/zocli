package actions

import (
	"bytes"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	mock_terminal "github.com/berrybytes/zocli/pkg/utils/terminal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProjectCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := mock_factory.NewFactory()
		cmd := NewProjectDeleteCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"d", "del", "de", "dele"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectDeleteCommand(f)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectDeleteCommand(f)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectDeleteCommand(f)

		assert.Equal(t, "delete", cmd.Use, "Expected command name to be 'delete'")
	})

	t.Run("CommandFlags", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectDeleteCommand(f)

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
			out:  grammar.ProjectDeleteHelp,
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
			deleteCmd := NewProjectDeleteCommand(f)

			deleteCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			deleteCmd.SetIn(stdin)
			deleteCmd.SetOut(stderr)
			deleteCmd.SetErr(stderr)

			_, err := deleteCmd.ExecuteC()
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

func TestDeleteIDAndNameBoth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f := mock_factory.NewFactory()
	f.LoggedIn = true
	f.Printer = recorder
	recorder.EXPECT().Fatal(1, "provide either id or name, but not both").Times(1)

	getCmd := NewProjectDeleteCommand(f)
	getCmd.SetArgs([]string{"--id", "23198", "--name", "fkjahskdfhaksd"})
	_, _ = getCmd.ExecuteC()
}

func TestFetchProjByID(t *testing.T) {
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.GetSingleProject

	f := mock_factory.NewFactory()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name   string
		id     string
		expect func(recorder *mock_printer.MockPrinterInterface)
	}{
		{
			name: "invalid id",
			id:   "123",
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				recorder.EXPECT().Fatal(8, "invalid id").Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			test.expect(recorder)
			f.Printer = recorder

			f.Term = mock_terminal.NewMockProvider()

			requester.Client = &mock.Client{}
			mock.GetDoFunc = mock.GetSingleProject

			o := &Opts{F: f}
			o.fetchProjByID(test.id)
		})
	}
}
