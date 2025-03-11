package models

import (
	"errors"

	"github.com/berrybytes/zocli/api"
)

type Project struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`

	MetaData MetaData `yaml:"metadata" json:"metadata"`
	Spec     Spec     `yaml:"spec" json:"spec"`
}

type App struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`

	MetaData MetaData `yaml:"metadata" json:"metadata"`
	Spec     AppSpec  `yaml:"spec" json:"spec"`
}

// Env
//
// general rules parsed from server side:
// name -> allowed alphanumeric, underscore, hyphen and space only,
// resource id is -> required,
// application id -> required,
// plugin version id -> required,
// replicas -> must be greater than 0,
// either version or service type -> required,
type Env struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`

	// MetaData
	// this struct is responsible for holding the metadata of the environment,
	// it will hold the id, created at, name and organization id,
	MetaData MetaData `yaml:"metadata" json:"metadata"`
	Spec     EnvSpec  `yaml:"spec" json:"spec"`
}

// EnvSpec
//
// this struct is responsible for holding the spec of the environment,
// it will hold the application id, project id, resource id, plugin version id, replicas and version..etc.,
// these are the fields which will be available on the manifest file
type EnvSpec struct {
	Application     IDOrName `json:"application" yaml:"application"`
	Project         IDOrName `json:"project" yaml:"project"`
	ResourceID      int      `json:"resource_id" yaml:"resource_id"`
	PluginVersionID int      `json:"plugin_version_id" yaml:"plugin_version_id"`
	Replicas        int      `json:"replicas" yaml:"replicas"`
	Version         Version  `json:"version" yaml:"version"`
}

type Version struct {
	Name string `json:"name" yaml:"name"`
	Repo string `json:"repo" yaml:"repo"`
	Tag  string `json:"tag" yaml:"tag"`
}

// ValidateCreation
//
// this function is responsible for validating the environment creation
// it will check if the environment name is provided
// it will check if the application name or id is provided
// it will check if the project name or id is provided
// it will check if the version name is provided
// it will check if the version repo is provided
// it will check if the version tag is provided
// it will check if the replicas is provided and if not then default value will be 1
func (e *Env) ValidateCreation() error {
	if e.MetaData.Name == "" {
		return errors.New("provide environment name")
	}

	if e.Spec.Application.ID == 0 && e.Spec.Application.Name == "" {
		return errors.New("provide application name or id")
	}

	if e.Spec.Application.Name != "" && e.Spec.Project.ID == 0 && e.Spec.Project.Name == "" {
		return errors.New("provide project name or id inorder to get application id")
	}

	if e.Spec.Application.ID != 0 && e.Spec.Application.Name != "" {
		return errors.New("provide application name or id, and not both")
	}

	if e.Spec.Project.ID != 0 && e.Spec.Project.Name != "" {
		return errors.New("provide Project name or id, and not both")
	}

	if e.Spec.Replicas == 0 {
		e.Spec.Replicas = 1
	}

	if e.Spec.Version.Name == "" {
		return errors.New("provide version name")
	}

	if e.Spec.Version.Repo == "" {
		return errors.New("provide repo name")
	}

	if e.Spec.Version.Tag == "" {
		return errors.New("provide tag")
	}

	return nil
}

type ServiceType int

const (
	TEMPLATE ServiceType = iota
	GIT
	DOCKER
	HELM
	OPERATOR
)

func (p *Project) ValidateCreation() error {
	if p.MetaData.Name == "" {
		return errors.New("provide project name")
	}
	if p.Spec.Subscription.ID == 0 && p.Spec.Subscription.Name == "" {
		p.Spec.Subscription.ID = 1
	}
	if p.Spec.Logging == "" {
		p.Spec.Logging = "01Logs"
	}
	if p.Spec.Monitoring == "" {
		p.Spec.Monitoring = "Prometheus"
	}
	return nil
}

func (p *App) ValidateCreation() error {
	if p.MetaData.Name == "" {
		return errors.New("provide project name")
	}
	if p.Spec.Project.ID == 0 {
		return errors.New("invalid project")
	}
	if p.Spec.Plugin.ID == 0 {
		return errors.New("invalid plugin")
	}
	if p.Spec.Cluster.Region == "" {
		return errors.New("required app region")
	}
	return nil
}

type Spec struct {
	Subscription Subscription   `yaml:"subscription" json:"subscription"`
	Logging      string         `yaml:"logging" json:"logging"`
	Monitoring   string         `yaml:"monitoring" json:"monitoring"`
	OptimizeCost bool           `yaml:"optimize_cost" json:"optimize_cost"`
	DedicatedLB  bool           `yaml:"dedicated_lb" json:"dedicated_lb"`
	Variables    []api.Variable `yaml:"variables,omitempty" json:"variables,omitempty"`
	Owner        Owner          `yaml:"owner" json:"owner"`
}

type AppSpec struct {
	Cluster             Cluster     `yaml:"cluster" json:"cluster"`
	ServiceType         ServiceType `yaml:"service_type" json:"service_type"`
	GitService          string      `yaml:"git_service"  json:"git_service"`
	GitRepoUrl          string      `yaml:"git_repo_url" json:"git_repo_url"`
	GitRepository       string      `yaml:"git_repository" json:"git_repository"`
	Plugin              Plugin      `yaml:"plugin" json:"plugin"`
	Project             IDOrName    `yaml:"project" json:"project"`
	ChartID             string      `yaml:"chart_id,omitempty" json:"chart_id,omitempty"`
	ImageNamespace      string      `yaml:"image_namespace,omitempty" json:"image_namespace,omitempty"`
	ImageRepo           string      `yaml:"image_repo,omitempty" json:"image_repo,omitempty"`
	ImageService        string      `yaml:"imageService,omitempty" json:"image_service,omitempty"`
	ImageUrl            string      `yaml:"image_service,omitempty" json:"image_url,omitempty"`
	OperatorPackageName string      `yaml:"operator_package_name,omitempty" json:"operator_package_name,omitempty"`
}

type Cluster struct {
	ID     uint64 `yaml:"id" json:"id"`
	Name   string `yaml:"name" json:"name"`
	Region string `yaml:"region" json:"region"`
}

type Plugin struct {
	ID   uint64 `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}
