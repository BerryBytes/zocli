package api

import "time"

type UserRoleEnum int

const (
	Admin UserRoleEnum = 1 + iota
	Members
)

type OrganizationMembersList struct {
	Members []OrganizationMember
}

type OrganizationMember struct {
	ID             int       `json:"id"`
	Createdat      time.Time `json:"createdat"`
	Owner          User      `json:"user,omitempty"`
	UserID         int       `json:"user_id"`
	Organization   any       `json:"organization"`
	OrganizationID int       `json:"organization_id"`
	UserRole       int       `json:"user_role"`
}
