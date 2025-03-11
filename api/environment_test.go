package api

import "testing"

func Test_Environment_FromJson(t *testing.T) {
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
			lr := &Environment{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_SingleEnvironment_FromJson(t *testing.T) {
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
			lr := &SingleEnvironment{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_EnvironmentPresenter_FromJson(t *testing.T) {
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
			lr := &EnvironmentPresenter{}
			err := lr.FromJSON(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_EnvironmentOverview_FromJson(t *testing.T) {
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
			lr := &EnvironmentOverview{}
			err := lr.FromJSON(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
