package requester

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/berrybytes/zocli/pkg/utils/requester/mock"
)

type RequestConfig struct {
	URL     string
	Headers map[string]string
	Method  string
	Body    io.Reader
	F       *factory.Factory
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func CleanClient() {
	Client = &http.Client{}
}

func (rc *RequestConfig) Request() *api.BaseResponse {
	rc.F.IO.StartProgressIndicator()
	client := Client
	if rc.URL == "" {
		rc.F.IO.StopProgressIndicator()
		rc.F.Printer.Fatal(3, "cannot request on blank URL")
		return nil
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequestWithContext(rc.F.Ctx, rc.Method, rc.URL, rc.Body)
	if err != nil {
		rc.F.IO.StopProgressIndicator()
		rc.F.Printer.Fatal(3, err.Error())
		return nil
	}

	rc.F.Debug.Debug(req.URL, " ======URL======")
	for key, value := range rc.Headers {
		req.Header.Add(key, value)
	}
	// addition of default Headers
	hostname, _ := os.Hostname()
	req.Header.Add("User-Agent", "cli "+rc.F.UserOS)
	if hostname != "" {
		req.Header.Add("Os-Hostname", hostname)
	}
	rc.F.Debug.Debug("from the requester REQUEST: ", *req)

	rc.F.Debug.Debug("body for the request", req.Body)
	res, err := client.Do(req)
	if mock.GetDoFunc != nil {
		fmt.Println(req.URL, " +==========")
		time.Sleep(10) // nolint:staticcheck
	}
	if err != nil {
		rc.F.IO.StopProgressIndicator()
		rc.F.Printer.Fatal(3, err.Error())
		return nil
	}

	var response api.BaseResponse
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		rc.F.IO.StopProgressIndicator()
		rc.F.Debug.Debug("from the requester RESPONSE: ", response)
		rc.F.Printer.Fatal(5, err.Error())
		return nil
	}

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		rc.F.IO.StopProgressIndicator()
		rc.F.Debug.Debug("from the requester RESPONSE: ", response)
		rc.F.Printer.Fatal(5, err.Error())
		return nil
	}

	defer res.Body.Close()

	if response.Success == 0 {
		rc.F.Debug.Debug("from the requester RESPONSE: ", response)
		if response.Message == "" {
			rc.F.IO.StopProgressIndicator()
			rc.F.Printer.Fatal(7)
		}

		rc.F.IO.StopProgressIndicator()
		rc.F.Printer.Fatal(8, response.Message)
		return nil
	}
	rc.F.Debug.Debug("from the requester RESPONSE: ", response)
	rc.F.IO.StopProgressIndicator()
	return &response
}
