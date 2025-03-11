package defaults

import (
	"reflect"
	"testing"

	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
)

func Test_TokenLogin(t *testing.T) {
	f := mock_factory.NewFactory()
	testCases := []struct {
		name              string
		input             []map[string]interface{}
		expectedURL       string
		expectedHeaders   map[string]string
		expectedMethod    string
		expectedHeaderLen int
	}{
		{
			name:              "Empty input",
			input:             []map[string]interface{}{},
			expectedURL:       f.Routes.GetRoute("profile"),
			expectedHeaderLen: 0,
			expectedMethod:    "GET",
		},
		{
			name: "Valid headers and body",
			input: []map[string]interface{}{
				{
					"headers": map[string]string{
						"Content-Type":  "text/csv",
						"Authorization": "Bearer abc123",
					},
					"body": []byte("Some content"),
				},
			},
			expectedURL:       f.Routes.GetRoute("profile"),
			expectedHeaderLen: 1,
			expectedHeaders: map[string]string{
				"Content-Type":  "text/csv",
				"Authorization": "Bearer abc123",
			},
			expectedMethod: "GET",
		},
		{
			name: "Multiple maps",
			input: []map[string]interface{}{
				{
					"headers": map[string]string{
						"Content-Type": "text/csv",
					},
					"body": []byte("Some content"),
				},
				{
					"headers": map[string]string{
						"Authorization": "Bearer abc123",
					},
				},
			},
			expectedURL:       f.Routes.GetRoute("profile"),
			expectedHeaderLen: 1,
			expectedHeaders: map[string]string{
				"Authorization": "Bearer abc123",
			},
			expectedMethod: "GET",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rc := Profile(f, tc.input...)
			if rc.URL != tc.expectedURL {
				t.Errorf("URL: got %v, want %v", rc.URL, tc.expectedURL)
			}
			if tc.expectedHeaderLen == 0 {
				if len(rc.Headers) != tc.expectedHeaderLen {
					t.Errorf("headers len: got %v, want %v", len(rc.Headers), tc.expectedHeaderLen)
				}
			} else if !reflect.DeepEqual(rc.Headers, tc.expectedHeaders) {
				t.Errorf("Headers: got %v, want %v", rc.Headers, tc.expectedHeaders)
			}
			if rc.Method != tc.expectedMethod {
				t.Errorf("Method: got %v, want %v", rc.Method, tc.expectedMethod)
			}
		})
	}
}
