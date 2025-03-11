package login

import (
	"testing"
	"time"

	mock_printer "github.com/berrybytes/zocli/pkg/utils/printer/mock"
	"github.com/golang/mock/gomock"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	mockfactory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	"github.com/berrybytes/zocli/pkg/utils/requester"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"

	"github.com/stretchr/testify/assert"
)

func TestRequestSSOCode(t *testing.T) {
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.BrowserLogin

	l := Opts{
		F: mockfactory.NewFactory(),
	}
	expected := api.SSOCode{Code: "abcdef"}
	sso, err := l.requestSSOCode()
	if err != nil {
		assert.NoError(t, err)
	}
	assert.Equal(t, &expected, sso)
}

type testStruct struct {
	Name       string
	F          *factory.Factory
	Method     string
	Ssocode    string
	SsoIniTime time.Time
	ExitCode   int
	ErrMessage string
}

func TestValidCodeReceived(t *testing.T) {
	test := testStruct{
		Name:     "valid token received",
		F:        mockfactory.NewFactory(),
		Method:   "POST",
		Ssocode:  "abcdef",
		ExitCode: 0,
	}
	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.BrowserLogin

	reqConfig := requester.RequestConfig{
		URL:    test.F.Routes.GetRoute("sso-code"),
		Method: test.Method,
		F:      test.F,
	}

	baseRes := reqConfig.Request()
	assert.Equal(t, 1, baseRes.Success)
	assert.Equal(t, "Success", baseRes.Message)
	var sso api.SSOCode
	err := sso.FromJson(baseRes.Data)
	assert.Empty(t, err)
	assert.Equal(t, test.Ssocode, sso.Code)
}

func TestBrowserFlow(t *testing.T) {
	opts := Opts{
		F:         mockfactory.NewFactory(),
		WithCreds: true,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recorder := mock_printer.NewMockPrinterInterface(ctrl)
	opts.F.Printer = recorder
	opts.SsoCode = "abcdef"

	requester.Client = &mock.Client{}
	mock.GetDoFunc = mock.BrowserLogin

	go func() {
		time.Sleep(1 * time.Second)
		requester.Client = &mock.Client{}
		mock.GetDoFunc = mock.SSOStatus

		mock.GetDoFunc = mock.Profile
	}()

	recorder.EXPECT().Print("Login in using the browser using a dynamically generated code.").Times(1)
	recorder.EXPECT().Printf("\n! First copy your one-time code: %s\n", "abcdef").Times(1)
	recorder.EXPECT().Print("Press Enter to open " + opts.F.Routes.GetRoute("device-login") + " in your browser...").Times(1)
	recorder.EXPECT().Print("Waiting for code status.").MaxTimes(5)
	recorder.EXPECT().Print(".").Times(1)
	recorder.EXPECT().Print(opts.F.IO.ColorScheme().SuccessIcon(), " Logged in successfully as ", "testmail@mail.com").Times(1)
	_ = opts.browserFlow()
}
