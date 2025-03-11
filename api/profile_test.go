package api

import "testing"

func Test_ProfileResponse_FromJson(t *testing.T) {
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
				"name": "email",
				"id":   1,
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
			p := &ProfileResponse{}
			err := p.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
