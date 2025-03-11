package vars

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
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

func TestNewProjectGetCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsGetCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"ge", "g"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	t.Run("IdFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsGetCommand(f)
		id, _ := cmd.Flags().GetString("id")

		assert.Equal(t, "", id, "Expected default id to be an empty string")
	})

	t.Run("NameFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsGetCommand(f)
		name, _ := cmd.Flags().GetString("name")

		assert.Equal(t, "", name, "Expected default name to be an empty string")
	})

	t.Run("OutputFlag", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsGetCommand(f)
		output, _ := cmd.Flags().GetString("out")

		assert.Equal(t, "", output, "Expected default output to be an empty string")
	})

	// Test case 6: Check command name
	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsGetCommand(f)

		assert.Equal(t, "get", cmd.Use, "Expected command name to be 'get'")
	})

	// Test case 7: Check if flags are correctly attached to the command
	t.Run("CommandFlags", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsGetCommand(f)

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
			out:  grammar.ProjectVarsGetHelp,
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
			getCmd := NewProjectVarsGetCommand(f)

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

func TestGetVars(t *testing.T) {
	f := mock_factory.NewFactory()
	f.LoggedIn = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	t.Run("name and id both", func(t *testing.T) {
		recorder.EXPECT().Fatal(1, "provide either id or name, but not both")

		getCmd := NewProjectVarsGetCommand(f)
		getCmd.SetArgs([]string{"--id", "123", "--name", "abc"})

		_, _ = getCmd.ExecuteC()
	})

	t.Run("get by id", func(t *testing.T) {
		o := &Opts{F: f, ID: "1"}
		recorder.EXPECT().Exit(0).Times(1)
		requester.Client = &mock.Client{}
		mock.GetDoFunc = mock.GetSingleProject
		o.getVars(&cobra.Command{}, []string{})
	})

	//t.Run("get by name", func(t *testing.T) {
	//	o := &Opts{F: f, Name: "testProj"}
	//	recorder.EXPECT().Exit(0).Times(1)
	//	requester.Client = &mock.Client{}
	//	mock.GetDoFunc = mock.GetProjectByName
	//	go func() {
	//		time.Sleep(3) // nolint:staticcheck
	//		mock.GetDoFunc = mock.GetSingleProject
	//	}()
	//	o.getVars(&cobra.Command{}, []string{})
	//})
}

func TestPrintVarsJSONOutput(t *testing.T) {
	tests := []struct {
		name   string
		supply []api.Variable
		expect func(recorder *mock_printer.MockPrinterInterface)
	}{
		{
			name: "single var",
			supply: []api.Variable{
				{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res, _ := json.Marshal([]api.Variable{
					{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
				}[0])
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
		},
		{
			name: "multiple vars",
			supply: []api.Variable{
				{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
				{Id: 2, Key: "TestKey2", Value: "TestValue2", Type: "SecretType"},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res, _ := json.Marshal([]api.Variable{
					{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
					{Id: 2, Key: "TestKey2", Value: "TestValue2", Type: "SecretType"},
				})
				var final bytes.Buffer
				_ = json.Indent(&final, res, "", "\t")
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := Opts{Output: "json", F: mock_factory.NewFactory()}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			o.F.Printer = recorder
			test.expect(recorder)
			o.printVars(test.supply)
		})
	}
}

func TestPrintVarsYAMLOutput(t *testing.T) {
	tests := []struct {
		name   string
		supply []api.Variable
		expect func(recorder *mock_printer.MockPrinterInterface)
	}{
		{
			name: "single var",
			supply: []api.Variable{
				{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res := []api.Variable{
					{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
				}[0]
				var final bytes.Buffer
				yamlEncoder := yaml.NewEncoder(&final)
				yamlEncoder.SetIndent(2)
				_ = yamlEncoder.Encode(&res)
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
		},
		{
			name: "multiple vars",
			supply: []api.Variable{
				{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
				{Id: 2, Key: "TestKey2", Value: "TestValue2", Type: "SecretType"},
			},
			expect: func(recorder *mock_printer.MockPrinterInterface) {
				var a []interface{}
				res := []api.Variable{
					{Id: 1, Key: "TestKey", Value: "TestValue", Type: "secret"},
					{Id: 2, Key: "TestKey2", Value: "TestValue2", Type: "SecretType"},
				}
				var final bytes.Buffer
				yamlEncoder := yaml.NewEncoder(&final)
				yamlEncoder.SetIndent(2)
				err := yamlEncoder.Encode(&res)
				fmt.Println(err)
				a = append(a, final.String())
				recorder.EXPECT().Print(a).Times(1)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := Opts{Output: "yaml", F: mock_factory.NewFactory() /* You need to set this to a correct instance */}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			recorder := mock_printer.NewMockPrinterInterface(ctrl)
			o.F.Printer = recorder
			test.expect(recorder)
			o.printVars(test.supply)
		})
	}
}
