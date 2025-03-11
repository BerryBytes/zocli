package middleware

import (
	"testing"

	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/golang/mock/gomock"
)

func TestNotLoggedIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	f := mock_factory.NewFactory()

	tests := []struct {
		name   string
		logged bool
	}{
		{
			name:   "not logged in",
			logged: false,
		},
		{
			name:   "logged in",
			logged: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f.Printer = recorder
			f.LoggedIn = test.logged
			if test.logged == false {
				recorder.EXPECT().Fatal(1, "not logged in").Times(1)
			} else {
				recorder.EXPECT().Fatal(1, "not logged in").Times(0)
			}
			LoggedIn(f)
		})
	}
}
