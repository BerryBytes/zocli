package defaults

import (
	"reflect"
	"testing"
)

func Test_headerAndBodyConv(t *testing.T) {
	testCases := []struct {
		name              string
		input             []map[string]interface{}
		expectedHeaders   map[string]string
		expectedBody      []byte
		expectedHeaderLen int
		expectedBodyLen   int
	}{
		{
			name:              "Empty input",
			input:             []map[string]interface{}{},
			expectedHeaderLen: 0,
			expectedBodyLen:   0,
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
			expectedHeaders: map[string]string{
				"Content-Type":  "text/csv",
				"Authorization": "Bearer abc123",
			},
			expectedBody:      []byte("Some content"),
			expectedBodyLen:   1,
			expectedHeaderLen: 1,
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
			expectedHeaders: map[string]string{
				"Authorization": "Bearer abc123",
			},
			expectedBody:      []byte("Some content"),
			expectedBodyLen:   1,
			expectedHeaderLen: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			headers, body := headersAndBodyConv(tc.input...)
			if tc.expectedHeaderLen == 0 {
				if len(headers) != tc.expectedHeaderLen {
					t.Errorf("headers len: got %v, want %v", len(headers), tc.expectedHeaderLen)
				}
			} else if !reflect.DeepEqual(headers, tc.expectedHeaders) {
				t.Errorf("headers: got %v, want %v", headers, tc.expectedHeaders)
			}
			if tc.expectedBodyLen == 0 {
				if len(body) != tc.expectedBodyLen {
					t.Errorf("body len: got %v, want %v", len(body), tc.expectedBodyLen)
				}
			} else if !reflect.DeepEqual(body, tc.expectedBody) {
				t.Errorf("body: got %v, want %v", body, tc.expectedBody)
			}
		})
	}
}
