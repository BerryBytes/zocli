package api

import "testing"

func Test_LoginResponse_FromJson(t *testing.T) {
	var testData = []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "Empty map",
			input:   map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "Valid data",
			input: map[string]interface{}{
				"Data": map[string]interface{}{
					"field1": "value1",
					"field2": "value2",
				},
			},
			wantErr: false,
		},
		{
			name:    "Invalid data",
			input:   "invalid data",
			wantErr: true,
		},
	}

	for _, tc := range testData {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			lr := &LoginResponse{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_Data_FromJson(t *testing.T) {
	var testData = []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "Empty map",
			input:   map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "Valid data",
			input: map[string]interface{}{
				"token": "token1",
				"user": map[string]interface{}{
					"name": "email",
					"id":   1,
				},
				"xpersonaltoken": "personalToken1",
			},
			wantErr: false,
		},
		{
			name:    "Invalid data",
			input:   "invalid data",
			wantErr: true,
		},
	}

	for _, tc := range testData {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			d := &Data{}
			err := d.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_SSOToken_FromJson(t *testing.T) {
	var testData = []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "Empty map",
			input:   map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "Valid data",
			input: map[string]interface{}{
				"code": "ssoToken1",
			},
			wantErr: false,
		},
		{
			name:    "Invalid data",
			input:   "invalid data",
			wantErr: true,
		},
	}

	for _, tc := range testData {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			s := &SSOCode{}
			err := s.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
