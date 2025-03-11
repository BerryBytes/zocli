package vars

import (
	"bytes"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectUpdateCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsUpdateCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"u", "up", "upda", "upd"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsUpdateCommand(f)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("VariableIdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsUpdateCommand(f)
		id, _ := cmd.Flags().GetInt64("vid")

		assert.Equal(t, int64(-1), id, "Expected default variable id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsUpdateCommand(f)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("ShowFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsUpdateCommand(f)
		name, _ := cmd.Flags().GetBool("show")

		assert.Equal(t, false, name, "Expected default show flag to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsUpdateCommand(f)

		assert.Equal(t, "update", cmd.Use, "Expected command name to be 'update'")
	})

	t.Run("CommandFlags", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsUpdateCommand(f)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("vid")
		assert.NotNil(t, flag, "Expected 'vid' flag to be attached to the command")

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
			out:  grammar.ProjectVarsUpdateHelp,
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
			getCmd := NewProjectVarsUpdateCommand(f)

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

func TestUpdateVar(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	t.Run("name and id both", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")

		getCmd := NewProjectVarsUpdateCommand(f)
		getCmd.SetArgs([]string{"--id", "123", "--name", "abc"})

		_, _ = getCmd.ExecuteC()
	})
}
