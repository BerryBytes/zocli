package models

import (
	"encoding/json"
	"errors"
)

type Organization struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`

	MetaData MetaData `yaml:"metadata" json:"metadata"`
	Spec     OrgSpec  `yaml:"spec" json:"spec"`
}

type OrgSpec struct {
	OrganizationPlan Subscription `yaml:"subscription" json:"subscription"`
	Domain           string       `yaml:"domain" json:"domain"`
	Description      string       `json:"description"`
}

func (o *OrgSpec) FromJson(data interface{}) error {
	var allMap map[string]interface{}
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &allMap)
	if err != nil {
		return err
	}

	err = o.OrganizationPlan.FromJson(allMap["organization_plan"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Organization) ValidateCreation() error {
	if o.MetaData.Name == "" {
		return errors.New("provide organization name")
	}
	if o.Spec.OrganizationPlan.ID == 0 {
		return errors.New("provide organization plan id")
	}
	return nil
}
