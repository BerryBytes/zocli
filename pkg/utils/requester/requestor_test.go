package requester

// func TestNilURL(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	recorder := mock_printer.NewMockPrinterInterface(ctrl)
//
// 	reqConfig := RequestConfig{
// 		F: mockfactory.NewFactory(),
// 	}
// 	reqConfig.F.Printer = recorder
// 	reqConfig.F.LoggedIn = true
// 	recorder.EXPECT().Fatal(3, "cannot request on blank URL").Times(1)
// 	_ = reqConfig.Request()
// }
//
// func TestInvalidURL(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	recorder := mock_printer.NewMockPrinterInterface(ctrl)
//
// 	reqConfig := RequestConfig{
// 		F:   mockfactory.NewFactory(),
// 		URL: "ws:google.com",
// 	}
// 	reqConfig.F.Printer = recorder
// 	reqConfig.F.LoggedIn = true
// 	recorder.EXPECT().Fatal(3, "Get \"ws:google.com\": unsupported protocol scheme \"ws\"").Times(1)
// 	_ = reqConfig.Request()
// }
//
// func TestNilBody(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	recorder := mock_printer.NewMockPrinterInterface(ctrl)
//
// 	reqConfig := RequestConfig{
// 		F:      mockfactory.NewFactory(),
// 		URL:    mockfactory.NewFactory().Routes.GetRoute("login"),
// 		Method: "POST",
// 	}
// 	reqConfig.F.Printer = recorder
// 	reqConfig.F.LoggedIn = true
// 	recorder.EXPECT().Fatal(8, "unexpected end of JSON input").Times(1)
// 	_ = reqConfig.Request()
// }
