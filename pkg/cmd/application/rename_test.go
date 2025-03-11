package application

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	interfaceMock "github.com/berrybytes/zocli/pkg/cmd/application/mock"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectApplicationRenameCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}

	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewApplicationRenameCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"r", "re", "renam", "ren"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewApplicationRenameCommand(&o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("ApplicationID", func(t *testing.T) {
		cmd := NewApplicationRenameCommand(&o)
		id, _ := cmd.Flags().GetString("aid")

		assert.Equal(t, "", id, "Expected default application id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewApplicationRenameCommand(&o)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("ApplicationNameFlag", func(t *testing.T) {
		cmd := NewApplicationRenameCommand(&o)
		name, _ := cmd.Flags().GetString("appname")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		cmd := NewApplicationRenameCommand(&o)

		assert.Equal(t, "rename", cmd.Use, "Expected command name to be 'list'")
	})

	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewApplicationRenameCommand(&o)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("name")
		assert.NotNil(t, flag, "Expected 'name' flag to be attached to the command")

		flag = cmd.Flag("pid")
		assert.NotNil(t, flag, "Expected 'pid' flag to be attached to the command")

		flag = cmd.Flag("pname")
		assert.NotNil(t, flag, "Expected 'pname' flag to be attached to the command")
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
			out:  grammar.ApplicationRenameHelp,
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
			getCmd := NewApplicationRenameCommand(&o)

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

func Test_RenameApplication(t *testing.T) {
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
		expected := NewApplicationRenameCommand(&o)
		mustEqual := expected
		interfaceRecorder.EXPECT().RenameRunner(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})
	t.Run("both id and name", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewApplicationRenameCommand(&o)
		o.ProjectID = "123"
		o.ProjectName = "abc"
		o.RenameRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by id", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetApps().Times(1)
		interfaceRecorder.EXPECT().PrintApps().Times(1)
		recorder.EXPECT().Print("Application ID : ")
		recorder.EXPECT().Fatal(1, o.F.IO.ColorScheme().FailureIcon()+" Error reading input")
		_ = NewApplicationRenameCommand(&o)
		o.ProjectID = "1"
		o.RenameRunner(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		o.List.Applications = append(o.List.Applications, api.Application{Id: 1})
		interfaceRecorder.EXPECT().GetProjectDetailByName().Times(1).Return(&api.Project{ID: 1}).Times(1)
		interfaceRecorder.EXPECT().GetApps().Times(1)
		interfaceRecorder.EXPECT().PrintApps().Times(1)
		interfaceRecorder.EXPECT().RenameApplication().Times(1)
		recorder.EXPECT().Fatal(1, "Only 3 to 30 alphanumeric, underscore, space & hyphen characters allowed!").Times(1)
		_ = NewApplicationRenameCommand(&o)

		o.ProjectName = "testProj"
		o.ApplicationID = "1"
		o.RenameRunner(&cobra.Command{}, []string{})
	})

	t.Run("nothing supplied", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		o.RenameRunner(&cobra.Command{}, []string{})
		assert.Equal(t, Opts{F: f, I: interfaceRecorder}, o)
	})
}

func Test_RenameApplicationFn(t *testing.T) {
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.RenameApplication

	o := &Opts{F: mock_factory.NewFactory(), ApplicationName: "testProj", ApplicationID: "1"}

	ctrl := gomock.NewController(t)
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	o.F.Printer = recorder
	recorder.EXPECT().Print(o.F.IO.ColorScheme().SuccessIcon(), " Application renamed to '"+o.ApplicationName+"'\n")
	o.RenameApplication()
}
