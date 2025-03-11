package login

import (
	"testing"

	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/golang/mock/gomock"

	mockfactory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
)

func TestNoCreds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.Login

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithCreds: true,
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Fatal(3, "required Password").Times(1)
	_ = opts.LoginWithCreds()
}

func TestInvalidEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.Login

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithCreds: true,
		Email:     "invalid",
		Password:  "invalid",
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Fatal(3, "invalid email").Times(1)
	_ = opts.LoginWithCreds()
}

func TestValidCreds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.Login

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithCreds: true,
		Email:     "testmail@mail.com",
		Password:  "invalid",
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Print(opts.F.IO.ColorScheme().SuccessIcon(), " Logged in successfully as ", opts.Email).Times(1)
	_ = opts.LoginWithCreds()
}

func TestAskCredsValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.Login

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithCreds: true,
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Print(opts.F.IO.ColorScheme().SuccessIcon(), " Logged in successfully as ", "testmail@mail.com").Times(1)
	_ = opts.askCreds()
}

func TestCredsFlow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.Login

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithCreds: true,
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Print(opts.F.IO.ColorScheme().SuccessIcon(), " Logged in successfully as ", "testmail@mail.com").Times(1)
	_ = opts.credsFlow()
}
