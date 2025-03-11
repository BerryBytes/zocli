package login

import (
	"testing"

	mockfactory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
	"github.com/golang/mock/gomock"
)

func TestInvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.TokenLogin

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithToken: true,
		WebToken:  "invalid token",
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Fatal(3, "no such token").Times(1)
	_ = opts.LoginWithToken()
}

func TestNillToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.TokenLogin

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithToken: true,
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Fatal(3, "required Token").Times(1)
	_ = opts.LoginWithToken()
}

func TestValidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.TokenLogin

	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithToken: true,
		WebToken:  "validToken",
	}
	opts.F.Printer = recorder
	recorder.EXPECT().Print(opts.F.IO.ColorScheme().SuccessIcon(), " Logged in successfully as ", "testmail@mail.com").Times(1)
	_ = opts.LoginWithToken()
}
