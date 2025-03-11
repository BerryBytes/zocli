package api

import (
	"encoding/json"
)

type LoginResponse struct {
	Data
}

func (l *LoginResponse) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, l)
}

type Data struct {
	AuthToken string `json:"token" yaml:"token"`
	User      User   `json:"user" yaml:"user"`
	WebToken  string `json:"personalToken" yaml:"personalToken"`
}

func (d *Data) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, d)
}

type SSOCode struct {
	Code string `json:"code"`
}

func (s *SSOCode) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, s)
}
