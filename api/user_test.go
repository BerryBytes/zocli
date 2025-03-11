package api

import "testing"

func Test_User_FromJson(t *testing.T) {
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
				"id":             1,
				"first_name":     "John",
				"last_name":      "Doe",
				"email":          "john.doe@example.com",
				"company":        "Example Inc.",
				"designation":    "Software Developer",
				"email_verified": true,
				"active":         true,
				"quotas": map[string]interface{}{
					"user_project":      1,
					"user_organization": 2,
				},
			},
			wantErr: false,
		},
		{
			name: "Missing Quotas",
			input: map[string]interface{}{
				"id":             1,
				"first_name":     "John",
				"last_name":      "Doe",
				"email":          "john.doe@example.com",
				"company":        "Example Inc.",
				"designation":    "Software Developer",
				"email_verified": true,
				"active":         true,
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
		t.Run(tc.name, func(t *testing.T) {
			u := &User{}
			err := u.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr {
				if id, ok := tc.input.(map[string]interface{})["id"]; ok {
					if u.Id != id.(int) {
						t.Errorf("Unexpected User Id. Got %v, want %v", u.Id, tc.input.(map[string]interface{})["id"].(int))
					}
				}
			}
		})
	}
}
