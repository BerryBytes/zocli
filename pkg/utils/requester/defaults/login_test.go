package defaults

import (
	"bytes"
	"testing"

	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
)

func Test_BasicLogin(t *testing.T) {
	f := mock_factory.NewFactory()
	testCases := []struct {
		name              string
		input             []map[string]interface{}
		expectedURL       string
		expectedBody      []byte
		expectedMethod    string
		expectedHeaderLen int
	}{
		{
			name:              "Empty input",
			input:             []map[string]interface{}{},
			expectedURL:       f.Routes.GetRoute("login"),
			expectedBody:      []byte{},
			expectedMethod:    "POST",
			expectedHeaderLen: 0,
		},
		{
			name: "Valid body",
			input: []map[string]interface{}{
				{
					"body": []byte("Some content"),
				},
			},
			expectedURL:       f.Routes.GetRoute("login"),
			expectedBody:      []byte("Some content"),
			expectedMethod:    "POST",
			expectedHeaderLen: 1,
		},
		{
			name: "Multiple maps",
			input: []map[string]interface{}{
				{
					"body": []byte("Some content"),
				},
				{
					"headers": map[string]string{
						"Authorization": "Bearer abc123",
					},
					"body": []byte("Different content"),
				},
			},
			expectedURL:       f.Routes.GetRoute("login"),
			expectedBody:      []byte("Different content"),
			expectedMethod:    "POST",
			expectedHeaderLen: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rc := BasicLogin(f, tc.input...)
			if rc.URL != tc.expectedURL {
				t.Errorf("URL: got %v, want %v", rc.URL, tc.expectedURL)
			}
			body := rc.Body.(*bytes.Buffer).Bytes()
			if tc.expectedHeaderLen == 0 {
				if len(rc.Headers) != tc.expectedHeaderLen {
					t.Errorf("headers len: got %v, want %v", len(rc.Headers), tc.expectedHeaderLen)
				}
			} else if !bytes.Equal(body, tc.expectedBody) {
				t.Errorf("Body: got %s, want %s", body, tc.expectedBody)
			}
			if rc.Method != tc.expectedMethod {
				t.Errorf("Method: got %v, want %v", rc.Method, tc.expectedMethod)
			}
		})
	}
}

func Test_RequestSSO(t *testing.T) {
	f := mock_factory.NewFactory()
	testCases := []struct {
		name           string
		expectedURL    string
		expectedMethod string
	}{
		{
			name:           "Basic test",
			expectedURL:    f.Routes.GetRoute("sso-code"),
			expectedMethod: "POST",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rc := RequestSSO(f)
			if rc.URL != tc.expectedURL {
				t.Errorf("URL: got %v, want %v", rc.URL, tc.expectedURL)
			}
			if rc.Method != tc.expectedMethod {
				t.Errorf("Method: got %v, want %v", rc.Method, tc.expectedMethod)
			}
		})
	}
}

func Test_RequestSSOStatus(t *testing.T) {
	f := mock_factory.NewFactory()
	testCases := []struct {
		name            string
		input           []map[string]interface{}
		expectedURL     string
		expectedBody    []byte
		expectedMethod  string
		expectedBodyLen int
	}{
		{
			name:            "Empty input",
			input:           []map[string]interface{}{},
			expectedURL:     f.Routes.GetRoute("sso-code"),
			expectedBodyLen: 0,
			expectedMethod:  "GET",
		},
		{
			name: "Valid body",
			input: []map[string]interface{}{
				{
					"body": []byte("Some content"),
				},
			},
			expectedURL:     f.Routes.GetRoute("sso-code"),
			expectedBody:    []byte("Some content"),
			expectedMethod:  "GET",
			expectedBodyLen: 1,
		},
		{
			name: "Multiple maps",
			input: []map[string]interface{}{
				{
					"body": []byte("Some content"),
				},
				{
					"headers": map[string]string{
						"Authorization": "Bearer abc123",
					},
					"body": []byte("Different content"),
				},
			},
			expectedURL:     f.Routes.GetRoute("sso-code"),
			expectedBody:    []byte("Different content"),
			expectedMethod:  "GET",
			expectedBodyLen: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rc := RequestSSOStatus(f, tc.input...)
			if rc.URL != tc.expectedURL {
				t.Errorf("URL: got %v, want %v", rc.URL, tc.expectedURL)
			}
			body := rc.Body.(*bytes.Buffer).Bytes()
			if tc.expectedBodyLen == 0 {
				if len(rc.Headers) != tc.expectedBodyLen {
					t.Errorf("body len: got %v, want %v", len(body), tc.expectedBodyLen)
				}
			} else if !bytes.Equal(body, tc.expectedBody) {
				t.Errorf("Body: got %s, want %s", body, tc.expectedBody)
			}
			if rc.Method != tc.expectedMethod {
				t.Errorf("Method: got %v, want %v", rc.Method, tc.expectedMethod)
			}
		})
	}
}
