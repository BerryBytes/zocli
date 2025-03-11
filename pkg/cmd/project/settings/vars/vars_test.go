package vars

import (
	"testing"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	mockfactory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectVarsDetail(t *testing.T) {
	f := mockfactory.NewFactory()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f.Printer = recorder

	t.Run("valid id", func(t *testing.T) {
		o := &Opts{
			F: f,
		}

		requester.Client = &mock.Client{}
		mock.GetDoFunc = mock.GetSingleProject
		o.getProjectDetail("1")
	})
	t.Run("invalid id", func(t *testing.T) {
		o := &Opts{
			F: f,
		}

		recorder.EXPECT().Fatal(8, "invalid id")
		requester.Client = &mock.Client{}
		mock.GetDoFunc = mock.GetSingleProject
		o.getProjectDetail("12")
	})

	t.Run("by invalid name", func(t *testing.T) {
		o := &Opts{
			F:    f,
			Name: "abc",
		}

		recorder.EXPECT().Fatal(3, "no such project found")
		requester.Client = &mock.Client{}
		mock.GetDoFunc = mock.GetProjectByName
		o.getProjectDetailByName()
	})

	// t.Run("by valid name", func(t *testing.T) {
	// 	o := &Opts{
	// 		F:    f,
	// 		Name: "testProj",
	// 	}
	//
	// 	requester.Client = &mock.Client{}
	// 	mock.GetDoFunc = mock.GetProjectByName
	// 	go func() {
	// 		time.Sleep(3) // nolint:staticcheck
	// 		mock.GetDoFunc = mock.GetSingleProject
	// 	}()
	// 	o.getProjectDetailByName()
	// })
}

func TestNewProjectVarCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsCommand(f)

		assert.NotNil(t, cmd, "Expect command to be not nil")
		assert.Equal(t, []string{"v", "va", "var", "variable", "variables", "vari"}, cmd.Aliases, "Expect aliases to be equal")
		assert.True(t, cmd.DisableFlagsInUseLine, "Expect flags in use line to be disabled")
	})

	// Test case 6: Check command name
	t.Run("CommandName", func(t *testing.T) {
		f := &factory.Factory{}
		cmd := NewProjectVarsCommand(f)

		assert.Equal(t, "vars", cmd.Use, "Expected command name to be 'vars'")
	})
}
