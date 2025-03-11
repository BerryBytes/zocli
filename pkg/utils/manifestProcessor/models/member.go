package models

import (
	"errors"
	"strings"
)

type OrganizationMember struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`

	MetaData MetaData   `yaml:"metadata" json:"metadata"`
	Spec     MemberSpec `yaml:"spec" json:"spec"`
}

type MemberSpec struct {
	Email string `yaml:"email" json:"email"`
	Role  string `yaml:"role" json:"role"`
}

func (p *OrganizationMember) ValidateCreation() error {
	if p.MetaData.Name == "" {
		return errors.New("required name")
	}
	if p.Spec.Email == "" {
		return errors.New("required email")
	}
	if p.Spec.Role == "" {
		return errors.New("required role")
	}
	return nil
}

func GetRole(id int) string {
	switch id {
	case 1:
		return "admin"
	case 2:
		return "member"
	default:
		return ""
	}
}

func SetRole(role string) int {
	switch strings.ToLower(role) {
	case "admin":
		return 1
	case "member":
		return 2
	default:
		return 2
	}
}
