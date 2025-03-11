package api

import "encoding/json"

type User struct {
	Id            int    `json:"id" yaml:"id"`
	FirstName     string `json:"first_name" yaml:"first_name"`
	LastName      string `json:"last_name" yaml:"last_name"`
	Email         string `json:"email" yaml:"email"`
	Company       string `json:"company" yaml:"company"`
	Designation   string `json:"designation" yaml:"designation"`
	EmailVerified bool   `json:"email_verified" yaml:"email_verified"`
	Active        bool   `json:"active" yaml:"active"`
	Quotas        struct {
		UserProject      int `json:"user_project" yaml:"user_project"`
		UserOrganization int `json:"user_organization" yaml:"user_organization"`
	} `json:"quotas" yaml:"quotas"`
}

func (u *User) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, u)
}
