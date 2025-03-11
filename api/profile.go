package api

import "encoding/json"

type ProfileResponse struct {
	User
}

func (p *ProfileResponse) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, p)
}
