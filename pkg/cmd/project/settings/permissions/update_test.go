package permissions

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
	interfaceMock "github.com/berrybytes/zocli/pkg/cmd/project/settings/permissions/mock"
	"github.com/spf13/cobra"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectPermissionsUpdateCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := &Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewProjectPermissionsUpdateCommand(o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"up", "u", "upd", "upda", "update"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewProjectPermissionsUpdateCommand(o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("Permission Id Flag", func(t *testing.T) {
		cmd := NewProjectPermissionsUpdateCommand(o)
		pid, _ := cmd.Flags().GetInt64("pid")

		assert.Equal(t, int64(-1), pid, "Expected default id to be -1 as default value")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewProjectPermissionsUpdateCommand(o)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewProjectPermissionsUpdateCommand(o)

		assert.Equal(t, "update", cmd.Use, "Expected command name to be 'update'")
	})

	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewProjectPermissionsUpdateCommand(o)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("pid")
		assert.NotNil(t, flag, "Expected 'pid' flag to be attached to the command")

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
			out:  grammar.ProjectPermissionUpdateHelp,
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
			getCmd := NewProjectPermissionsUpdateCommand(o)

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

func TestUpdatePermission(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrlNew := gomock.NewController(t)
	defer ctrlNew.Finish()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	printerRecorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = printerRecorder

	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := &Opts{F: f, I: interfaceRecorder}

	t.Run("name and id both", func(t *testing.T) {
		getCmd := NewProjectPermissionsUpdateCommand(o)
		getCmd.SetArgs([]string{"--id", "123", "--name", "abc"})
		interfaceRecorder.EXPECT().UpdatePermissions(getCmd, []string{}).Times(1)

		_, _ = getCmd.ExecuteC()
	})
	t.Run("both id and name", func(t *testing.T) {
		printerRecorder.EXPECT().Fatal(1, "provide either id or name, but not both").Times(1)
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewProjectPermissionsUpdateCommand(&o)
		o.ID = "123"
		o.Name = "abc"
		o.UpdatePermissions(&cobra.Command{}, []string{})
	})

	t.Run("get by id", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetPermission().Times(1)
		printerRecorder.EXPECT().Fatal(1, "no permissions found").Times(1)
		_ = NewProjectPermissionsUpdateCommand(&o)
		o.ID = "1"
		o.UpdatePermissions(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetPermission().Times(1)
		interfaceRecorder.EXPECT().GetProjectDetailByName().Times(1).Return(&api.Project{ID: 1})
		printerRecorder.EXPECT().Fatal(1, "no permissions found").Times(1)
		_ = NewProjectPermissionsUpdateCommand(&o)

		printerRecorder.EXPECT().Exit(0).Times(1)
		o.Name = "testProj"
		o.UpdatePermissions(&cobra.Command{}, []string{})
	})
	t.Run("nothing supplied", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		o.UpdatePermissions(&cobra.Command{}, []string{})
		assert.Equal(t, Opts{F: f, I: interfaceRecorder}, o)
	})
}

func Test_PushChanges(t *testing.T) {
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.UpdatePermissions

	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	printerRecorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = printerRecorder
	printerRecorder.EXPECT().Print(f.IO.ColorScheme().SuccessIcon(), " Successfully role changed to 'test'").Times(1)
	printerRecorder.EXPECT().Exit(0).Times(1)

	o := &Opts{F: f}

	o.ID = "1"
	o.Change.Id = 1
	o.UpdateRole = "test"
	o.RoleID = 1
	o.pushChanges()
}

func Test_CheckIfValidID(t *testing.T) {
	f := mock_factory.NewFactory()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	printerRecorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = printerRecorder

	t.Run("invalid permission id", func(t *testing.T) {
		printerRecorder.EXPECT().Fatal(1, "no such permission found.")
		o := &Opts{F: f, Permissions: api.Permissions{Permissions: []api.Permission{
			{Id: 1},
		}}}
		o.PermissionID = 2

		o.checkIfValidID()
	})
	t.Run("valid permission id", func(t *testing.T) {
		o := &Opts{F: f, Permissions: api.Permissions{Permissions: []api.Permission{
			{Id: 1},
		}}}
		o.PermissionID = 1

		o.checkIfValidID()
	})
}
