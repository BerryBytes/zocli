package actions

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestEnableProjectCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := mock_factory.NewFactory()
		cmd := NewProjectEnableCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"e", "ena", "able", "enab"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectEnableCommand(f)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectEnableCommand(f)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectEnableCommand(f)

		assert.Equal(t, "enable", cmd.Use, "Expected command name to be 'enable'")
	})

	t.Run("CommandFlags", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectEnableCommand(f)

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
			out:  grammar.ProjectEnableHelp,
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
			enableCmd := NewProjectEnableCommand(f)

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

func TestEnableByID(t *testing.T) {
	requester.Client = &mock.Client{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	requester.Client = &mock.Client{}

	tests := []struct {
		name      string
		expect    func(recorder *mock_printer.MockPrinterInterface)
		id        string
		getDoFunc func(req *http.Request) (*http.Response, error)
	}{
		{
			name: "Invalid ID",
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				recorder.EXPECT().Fatal(1, "invalid id supplied").Times(1)
				recorder.EXPECT().Fatal(5, "cannot proceed")
				recorder.EXPECT().Fatal(3, "no such id found").Times(1)
			},
			id:        "123abc",
			getDoFunc: mock.EnableProjectByID,
		},
		{
			name: "No Such Project",
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				recorder.EXPECT().Fatal(5, "cannot proceed")
				recorder.EXPECT().Fatal(3, "no such id found").Times(1)
			},
			id:        "123",
			getDoFunc: mock.EnableProjectByID,
		},
		{
			name: "Success",
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				recorder.EXPECT().Fatal(5, "cannot proceed")
				recorder.EXPECT().Exit(0).Times(1)
			},
			id:        "1",
			getDoFunc: mock.EnableProjectByID,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			f := mock_factory.NewFactory()
			f.Printer = recorder
			test.expect(recorder)
			mock.GetDoFunc = test.getDoFunc

			if test.name == "Success" {
				recorder.EXPECT().Print(f.IO.ColorScheme().SuccessIcon(), " Successfully activated project.").Times(1)
			}

			opts := &Opts{F: f}
			opts.ID = test.id
			opts.F.LoggedIn = true
			opts.enable(&cobra.Command{}, []string{})
		})
	}
}

func TestBothIDAndname(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f := mock_factory.NewFactory()
	f.Printer = recorder

	recorder.EXPECT().Fatal(1, "provide either id or name, but not both").Times(1)
	opts := &Opts{F: f}
	opts.ID = "1234"
	opts.Name = "test"
	opts.enable(&cobra.Command{}, []string{})
}

func TestEnableByName(t *testing.T) {
	requester.Client = &mock.Client{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	requester.Client = &mock.Client{}
	f := mock_factory.NewFactory()

	tests := []struct {
		name      string
		expect    func(recorder *mock_printer.MockPrinterInterface)
		projName  string
		getDoFunc func(req *http.Request) (*http.Response, error)
	}{
		{
			name: "No Such Project",
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				recorder.EXPECT().Fatal(3, "no such project found").Times(1)
			},
			projName:  "abc",
			getDoFunc: mock.GetProjectByName,
		},
		//{
		//	name: "Success",
		//	expect: func(recorder *mock_printer.MockPrinterInterface) {
		//		recorder.EXPECT().Print(f.IO.ColorScheme().SuccessIcon(), " Successfully activated project.")
		//		recorder.EXPECT().Exit(0).Times(1)
		//	},
		//	projName:  "testProj",
		//	getDoFunc: mock.GetProjectByName,
		//},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			f.Printer = recorder
			test.expect(recorder)
			mock.GetDoFunc = test.getDoFunc
			// if test.name == "Success" {
			// 	go func() {
			// 		time.Sleep(3) // nolint:staticcheck
			// 		mock.GetDoFunc = mock.EnableProjectByID
			// 	}()
			// }

			opts := &Opts{F: f}
			opts.Name = test.projName
			opts.F.LoggedIn = true
			opts.enable(&cobra.Command{}, []string{})
		})
	}
}
