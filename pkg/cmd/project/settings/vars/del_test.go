package vars

import (
	"bytes"
	"testing"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectDelCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"del", "d", "de", "dele", "delete", "remove", "r"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("VariableIdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)
		id, _ := cmd.Flags().GetInt64("vid")

		assert.Equal(t, int64(-1), id, "Expected default variable id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("VariableNameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)
		name, _ := cmd.Flags().GetString("vname")

		assert.Equal(t, "", name, "Expected default variable name to be an empty string")
	})

	t.Run("ShowFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)
		name, _ := cmd.Flags().GetBool("show")

		assert.Equal(t, false, name, "Expected default show flag to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)

		assert.Equal(t, "delete", cmd.Use, "Expected command name to be 'delete'")
	})

	t.Run("CommandFlags", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsDeleteCommand(f)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("vid")
		assert.NotNil(t, flag, "Expected 'vid' flag to be attached to the command")

		flag = cmd.Flag("name")
		assert.NotNil(t, flag, "Expected 'name' flag to be attached to the command")

		flag = cmd.Flag("vname")
		assert.NotNil(t, flag, "Expected 'vname' flag to be attached to the command")
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
			out:  grammar.ProjectVarsDeleteHelp,
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
			getCmd := NewProjectVarsDeleteCommand(f)

			getCmd.SetArgs(tc.args)

			stdin := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			getCmd.SetIn(stdin)
			getCmd.SetOut(stderr)
			getCmd.SetErr(stderr)

			_, err := getCmd.ExecuteC()
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

func TestDelVar(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	t.Run("name and id both", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")

		getCmd := NewProjectVarsDeleteCommand(f)
		getCmd.SetArgs([]string{"--id", "123", "--name", "abc"})

		_, _ = getCmd.ExecuteC()
	})
}

func Test_Opts_deleteVar(t *testing.T) {
	o := &Opts{
		AllVars: []api.Variable{
			{Id: 1, Key: "test", Type: "secret", Value: "this is a test value"},
			{Id: 2, Key: "test2", Type: "secret", Value: "this is a test value"},
			{Id: 3, Key: "test3", Type: "secret", Value: "this is a test value"},
		},
		F: mock_factory.NewFactory(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	o.F.Printer = recorder
	recorder.EXPECT().Print(o.F.IO.ColorScheme().SuccessIcon(), " Successfully Deleted variable").Times(1)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.DeleteVariable

	deleteVar := api.Variable{Id: 2, Key: "test2", Type: "secret", Value: "this is a test value"}
	o.deleteVar(deleteVar)
	assert.NotContains(t, o.AllVars, deleteVar)
}
