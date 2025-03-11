package api

import (
	"encoding/json"
	"errors"
	"time"
)

type Project struct {
	ID             int          `json:"id"`
	Createdat      time.Time    `json:"createdat"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	ProjectCode    string       `json:"project_code"`
	Tags           string       `json:"tags"`
	Region         string       `json:"region"`
	Active         bool         `json:"active"`
	ClusterScope   int          `json:"cluster_scope"`
	Subscription   Subscription `json:"subscription"`
	SubscriptionId int          `json:"subscription_id"`
	User           User         `json:"user"`
	UserId         int          `json:"user_id"`
	OrganizationId int          `json:"organization_id"`
	Logging        string       `json:"logging"`
	Monitoring     string       `json:"monitoring"`
	DedicatedLb    bool         `json:"dedicated_lb"`
	Variables      []Variable   `json:"variables"`
}

func (p *Project) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &p)
}

type ProjectList struct {
	Projects []Project `json:"projects" yaml:"projects"`
}

func (p *ProjectList) FromJson(a interface{}) error {
	var projects []Project
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &projects)
	p.Projects = projects
	return err
}

type Variable struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (v *Variable) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &v)
}

type Subscription struct {
	Id             int       `json:"id"`
	Createdat      time.Time `json:"createdat"`
	Name           string    `json:"name"`
	Apps           int       `json:"apps"`
	DiskSpace      int       `json:"disk_space"`
	Memory         int       `json:"memory"`
	Cores          int       `json:"cores"`
	DataTransfer   int       `json:"data_transfer"`
	Price          int       `json:"price"`
	Weight         int       `json:"weight"`
	Active         bool      `json:"active"`
	OrganizationId int       `json:"organization_id"`
	CiBuild        int       `json:"ci_build"`
	Attributes     string    `json:"attributes"`
	CronJob        int       `json:"cron_job"`
	Backups        int       `json:"backups"`
	LoadBalancer   int       `json:"load_balancer"`
	PriceList      []Price   `json:"price_list"`
}

type Price struct {
	DataTransfer int `json:"data_transfer"`
	LoadBalancer int `json:"load_balancer"`
}

func (s *Subscription) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s)
}

type ResourseList struct {
	Pods     string `json:"pods"`
	Secrets  string `json:"secrets"`
	Services string `json:"services"`
}

func (r *ResourseList) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &r)
}

type Resource struct {
	Memory       int `json:"memory"`
	Disk         int `json:"disk"`
	Core         int `json:"core"`
	DataTransfer struct {
		DataTransfer struct {
			Receive  float64 `json:"receive"`
			Transmit float64 `json:"transmit"`
		} `json:"data_transfer"`
	} `json:"data_transfer"`
	TotalCiBuild int `json:"total_ci_build"`
	TotalCronJob int `json:"total_cron_job"`
	Apps         int `json:"apps"`
}

func (r *Resource) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &r)
}

type Permission struct {
	Id            int       `json:"id"`
	CreatedAt     time.Time `json:"createdat"`
	Email         string    `json:"email"`
	User          User      `json:"user"`
	UserId        int       `json:"user_id"`
	UserRole      UserRole  `json:"user_role"`
	UserRoleId    int       `json:"user_role_id"`
	ProjectId     int       `json:"project_id"`
	ApplicationId int       `json:"application_id"`
	EnvironmentId int       `json:"environment_id"`
	Active        bool      `json:"active"`
	Attributes    string    `json:"attributes"`
	GroupId       int       `json:"group_id"`
}

func (p *Permission) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &p)
}

type UserRole struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"createdat"`
	Name        string    `json:"name"`
	Code        int       `json:"code"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
}

func (r *UserRole) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &r)
}

type Permissions struct {
	Permissions []Permission
}

func (p *Permissions) FromJson(a interface{}) error {
	var perms []Permission
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &perms)
	p.Permissions = perms
	return err
}

type LoadBalancer struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CustomDomain string `json:"custom_domain"`
	Region       string `json:"region"`
	ClusterID    uint64 `json:"cluster_id"`
	ProjectID    uint64 `json:"project_id"`
}

type LoadBalancers struct {
	LoadBalancers []LoadBalancer
}

func (l *LoadBalancers) FromJson(a interface{}) error {
	var loads []LoadBalancer
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &loads)
	l.LoadBalancers = loads
	return err
}

func (l *LoadBalancer) ValidateCreation() error {
	if l.Name == "" {
		return errors.New("provide project name")
	}
	if l.ProjectID == 0 {
		return errors.New("provide name")
	}
	if l.Region == "" {
		return errors.New("provide region")
	}
	return nil
}

type Role struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"createdat"`
	Name        string    `json:"name"`
	Code        int       `json:"code"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
}

func (r *Role) FromJson(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &r)
}

type Roles struct {
	Roles []Role
}

func (r *Roles) FromJson(a interface{}) error {
	var roles []Role
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &roles)
	r.Roles = roles
	return err
}
