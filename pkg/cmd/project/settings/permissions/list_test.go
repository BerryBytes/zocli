package permissions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	interfaceMock "github.com/berrybytes/zocli/pkg/cmd/project/settings/permissions/mock"
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

func TestNewProjectPermissionsListCommand(t *testing.T) {
	f := &factory.Factory{}
	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()
	o := Opts{F: f, I: interfaceRecorder}
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewProjectPermissionsListCommand(&o)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"l", "li", "lis", "get", "g", "ge"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		cmd := NewProjectPermissionsListCommand(&o)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		cmd := NewProjectPermissionsListCommand(&o)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("OutputFlag", func(t *testing.T) {
		cmd := NewProjectPermissionsListCommand(&o)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	// Test case 6: Check command name
	t.Run("CommandName", func(t *testing.T) {
		cmd := NewProjectPermissionsListCommand(&o)

		assert.Equal(t, "list", cmd.Use, "Expected command name to be 'list'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		cmd := NewProjectPermissionsListCommand(&o)

		flag := cmd.Flag("id")
		assert.NotNil(t, flag, "Expected 'id' flag to be attached to the command")

		flag = cmd.Flag("name")
		assert.NotNil(t, flag, "Expected 'name' flag to be attached to the command")

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
			out:  grammar.ProjectPermissionListHelp,
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
			getCmd := NewProjectPermissionsListCommand(&o)

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

func Test_ListPermission(t *testing.T) {
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
		expected := NewProjectPermissionsListCommand(&o)
		expected.SetArgs([]string{"--id", "123", "--name", "abc"})
		mustEqual := expected
		interfaceRecorder.EXPECT().ListPermissions(expected, []string{}).Times(1)
		assert.Equal(t, mustEqual, expected)
		_, _ = expected.ExecuteC()
	})
	t.Run("both id and name", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")
		o := Opts{F: f, I: interfaceRecorder}
		_ = NewProjectPermissionsListCommand(&o)
		o.ID = "123"
		o.Name = "abc"
		o.ListPermissions(&cobra.Command{}, []string{})
	})

	t.Run("get by id", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetPermission().Times(1)
		_ = NewProjectPermissionsListCommand(&o)
		o.ID = "1"
		recorder.EXPECT().Exit(0).Times(1)
		o.ListPermissions(&cobra.Command{}, []string{})
	})

	t.Run("get by name", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		interfaceRecorder.EXPECT().GetPermission().Times(1)
		interfaceRecorder.EXPECT().GetProjectDetailByName().Times(1).Return(&api.Project{ID: 1})
		_ = NewProjectPermissionsListCommand(&o)

		recorder.EXPECT().Exit(0).Times(1)
		o.Name = "testProj"
		o.ListPermissions(&cobra.Command{}, []string{})
	})
	t.Run("nothing supplied", func(t *testing.T) {
		o := Opts{F: f, I: interfaceRecorder}
		o.ListPermissions(&cobra.Command{}, []string{})
		assert.Equal(t, Opts{F: f, I: interfaceRecorder}, o)
	})
}

func Test_GetPermission(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	interfaceCtrl := gomock.NewController(t)
	interfaceRecorder := interfaceMock.NewMockInterface(interfaceCtrl)
	defer interfaceCtrl.Finish()

	interfaceRecorder.EXPECT().PrintPermissions().Times(1)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.ListProjectPermissions

	o := Opts{F: f, I: interfaceRecorder, ID: "1"}
	o.GetPermission()
}

func TestOpts_PrintPermissions(t *testing.T) {
	tests := []struct {
		name   string
		supply []api.Permission
		expect func(recorder *mock_printer.MockPrinterInterface)
		output string
	}{
		{
			name: "single var JSON", supply: []api.Permission{
				{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res, _ := json.Marshal([]api.Permission{
					{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
				}[0])
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "json",
		},
		{
			name: "multiple vars JSON",
			supply: []api.Permission{
				{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
				{Id: 2, Email: "test2@mail.com", UserId: 2, Active: true},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res, _ := json.Marshal([]api.Permission{
					{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
					{Id: 2, Email: "test2@mail.com", UserId: 2, Active: true},
				})
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "json",
		},
		{
			name: "single var YAML", supply: []api.Permission{
				{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res := []api.Permission{
					{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
				}[0]
				var final bytes.Buffer
				yamlEncoder := yaml.NewEncoder(&final)
				yamlEncoder.SetIndent(2)
				_ = yamlEncoder.Encode(&res)
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
			output: "yaml",
		},
		{
			name: "multiple vars YAML",
			supply: []api.Permission{
				{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
				{Id: 2, Email: "test2@mail.com", UserId: 2, Active: true},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res := []api.Permission{
					{Id: 1, Email: "test@mail.com", UserId: 1, Active: true},
					{Id: 2, Email: "test2@mail.com", UserId: 2, Active: true},
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
			o := Opts{Output: test.output, F: mock_factory.NewFactory(), Permissions: api.Permissions{Permissions: test.supply}}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			o.F.Printer = recorder
			test.expect(recorder)
			o.PrintPermissions()
		})
	}
	t.Run("without any output args", func(t *testing.T) {
		o := Opts{F: mock_factory.NewFactory(), Permissions: api.Permissions{Permissions: []api.Permission{{Id: 1}}}}
		o.PrintPermissions()
	})
}
