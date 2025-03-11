package models

import (
	"encoding/json"
	"time"
)

// MetaData
//
// this struct is responsible for holding the metadata any KIND manifest file,
// it will hold the id, created at, name and organization id,
// NOTE: the organization id is optional as by default it is 0 or the active organization on the context
type MetaData struct {
	ID             int       `yaml:"id" json:"id"`
	CreatedAt      time.Time `yaml:"createdat" json:"createdat"`
	Name           string    `yaml:"name" json:"name"`
	OrganizationId int       `yaml:"organization_id,omitempty" json:"organization_id,omitempty"`
}

type Subscription struct {
	ID   int    `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}

func (o *Subscription) FromJson(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteData, &o)
}

type Owner struct {
	ID    int    `yaml:"id" json:"id"`
	Email string `yaml:"email" json:"email"`
	Role  int    `yaml:"user_role,omitempty" json:"user_role,omitempty"`
}

func (o *Owner) FromJson(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteData, &o)
}

// IDOrName
//
// this struct is responsible for holding the id or name of any KIND manifest file,
// it will hold the id or name of the application, project, organization, etc...
// NOTE: the id and name are optional as by default they are 0 or empty
// And each KIND manifest file will have its own validation for this struct
type IDOrName struct {
	ID   int    `json:"id,omitempty" yaml:"id,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
