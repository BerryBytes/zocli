package api

import (
	"testing"
)

func Test_Project_FromJson(t *testing.T) {
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
			lr := &Project{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_ProjectList_FromJson(t *testing.T) {
	projList := &ProjectList{}

	t.Run("valid JSON", func(t *testing.T) {
		validJson := []Project{
			{
				Name: "Project A",
			},
			{
				Name: "Project B",
			},
		}
		err := projList.FromJson(validJson)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(projList.Projects) != 2 {
			t.Errorf("Expected 2 projects, got %d", len(projList.Projects))
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		invalidJson := "invalid json"
		err := projList.FromJson(invalidJson)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func Test_Subscription_FromJson(t *testing.T) {
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
			lr := &Subscription{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_ResourceList_FromJson(t *testing.T) {
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
			lr := &ResourseList{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_Permission_FromJson(t *testing.T) {
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
			lr := &Permission{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_PermissionList_FromJson(t *testing.T) {
	projList := &Permissions{}

	t.Run("valid JSON", func(t *testing.T) {
		validJson := []Permission{
			{
				Id:     1,
				Email:  "test@mail.com",
				Active: true,
			},
			{
				Id:     2,
				Email:  "test2@mail.com",
				Active: false,
			},
		}
		err := projList.FromJson(validJson)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(projList.Permissions) != 2 {
			t.Errorf("Expected 2 projects, got %d", len(projList.Permissions))
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		invalidJson := "invalid json"
		err := projList.FromJson(invalidJson)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func Test_UserRole_FromJson(t *testing.T) {
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
			lr := &UserRole{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_Role_FromJson(t *testing.T) {
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
			lr := &Role{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_RoleList_FromJson(t *testing.T) {
	projList := &Roles{}

	t.Run("valid JSON", func(t *testing.T) {
		validJson := []Permission{
			{
				Id:     1,
				Email:  "test@mail.com",
				Active: true,
			},
			{
				Id:     2,
				Email:  "test2@mail.com",
				Active: false,
			},
		}
		err := projList.FromJson(validJson)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(projList.Roles) != 2 {
			t.Errorf("Expected 2 projects, got %d", len(projList.Roles))
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		invalidJson := "invalid json"
		err := projList.FromJson(invalidJson)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func Test_Variable_FromJson(t *testing.T) {
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
			lr := &Variable{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func Test_Resource_FromJson(t *testing.T) {
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
			lr := &Resource{}
			err := lr.FromJson(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
