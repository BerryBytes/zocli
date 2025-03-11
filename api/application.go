package api

import (
	"encoding/json"
	"time"
)

type ServiceType int

const (
	TEMPLATE ServiceType = iota
	GIT
	DOCKER
	HELM
	OPERATOR
)

func GetServiceType(t ServiceType) string {
	switch t {
	case TEMPLATE:
		return "TEMPLATE"
	case GIT:
		return "GIT"
	case DOCKER:
		return "DOCKER"
	case HELM:
		return "HELM"
	case OPERATOR:
		return "OPERATOR"
	default:
		return "INVALID"
	}
}

type ApplicationPresenter struct {
	Id                  int               `json:"id"`
	Active              bool              `json:"active"`
	ClusterId           int               `json:"cluster_id"`
	Createdat           time.Time         `json:"createdat"`
	GitRepoUrl          string            `json:"git_repo_url"`
	GitRepository       string            `json:"git_repository"`
	GitRepositoryInfo   GitRepositoryInfo `json:"git_repository_info"`
	GitService          string            `json:"git_service"`
	ImageNamespace      string            `json:"image_namespace"`
	ImageRepo           string            `json:"image_repo"`
	ImageService        string            `json:"image_service"`
	ImageUrl            string            `json:"image_url"`
	Name                string            `json:"name"`
	OperatorPackageName string            `json:"operator_package_name"`
	OwnerId             int               `json:"owner_id"`
	PluginId            int               `json:"plugin_id"`
	ProjectId           int               `json:"project_id"`
	ServiceType         ServiceType       `json:"service_type"`
}

func (ap *ApplicationPresenter) FromJSON(a interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &ap)
}

type ApplicationPresenterList struct {
	Applications []ApplicationPresenter `json:"applications" yaml:"applications"`
}

type Application struct {
	Active              bool              `json:"active"`
	Cluster             Cluster           `json:"cluster"`
	ClusterId           int               `json:"cluster_id"`
	Createdat           time.Time         `json:"createdat"`
	Env                 []Env             `json:"env"`
	GitRepoUrl          string            `json:"git_repo_url"`
	GitRepository       string            `json:"git_repository"`
	GitRepositoryInfo   GitRepositoryInfo `json:"git_repository_info"`
	GitService          string            `json:"git_service"`
	Id                  int               `json:"id"`
	ImageNamespace      string            `json:"image_namespace"`
	ImageRepo           string            `json:"image_repo"`
	ImageService        string            `json:"image_service"`
	ImageUrl            string            `json:"image_url"`
	Name                string            `json:"name"`
	OperatorPackageName string            `json:"operator_package_name"`
	Owner               User              `json:"owner"`
	OwnerId             int               `json:"owner_id"`
	Plugin              Plugin            `json:"plugin"`
	PluginId            int               `json:"plugin_id"`
	Project             Project           `json:"project"`
	ProjectId           int               `json:"project_id"`
	ServiceType         ServiceType       `json:"service_type"`
}

func (a *Application) FromJson(d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &a)
}

type ApplicationList struct {
	Applications []Application
}

func (a *ApplicationList) FromJson(d interface{}) error {
	var apps []Application
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &apps)
	a.Applications = apps
	return err
}

type Plugin struct {
	Active      bool      `json:"active"`
	Attributes  string    `json:"attributes"`
	Createdat   time.Time `json:"createdat"`
	Description string    `json:"description"`
	Id          int       `json:"id"`
	Image       string    `json:"image"`
	IsAddOn     bool      `json:"is_add_on"`
	MinCpu      int       `json:"min_cpu"`
	MinMemory   int       `json:"min_memory"`
	Name        string    `json:"name"`
	SourceUrl   string    `json:"source_url"`
	SupportCi   bool      `json:"support_ci"`
}

type Cluster struct {
	Active              bool        `json:"active"`
	CloudStorage        interface{} `json:"cloud_storage"`
	ClusterRequestId    int         `json:"cluster_request_id"`
	Color               string      `json:"color"`
	Createdat           time.Time   `json:"createdat"`
	DnsId               int         `json:"dns_id"`
	Id                  int         `json:"id"`
	ImageRegistryId     int         `json:"image_registry_id"`
	Labels              string      `json:"labels"`
	Name                string      `json:"name"`
	Nodes               int         `json:"nodes"`
	OrganizationId      int         `json:"organization_id"`
	Provider            string      `json:"provider"`
	ProvisionPercentage float64     `json:"provision_percentage"`
	PvCapacity          int         `json:"pv_capacity"`
	Region              string      `json:"region"`
	StorageClass        string      `json:"storage_class"`
	TotalMemory         int         `json:"total_memory"`
	Weight              int         `json:"weight"`
	Zone                string      `json:"zone"`
}

type Env struct {
	Action             string            `json:"action"`
	Active             bool              `json:"active"`
	Application        Application       `json:"application"`
	ApplicationId      int               `json:"application_id"`
	ApplyImmediately   bool              `json:"apply_immediately"`
	AutoScaler         AutoScaler        `json:"auto_scaler"`
	Createdat          time.Time         `json:"createdat"`
	DeploymentStrategy interface{}       `json:"deployment_strategy"`
	GitBranch          string            `json:"git_branch"`
	GitRepositoryInfo  GitRepositoryInfo `json:"git_repository_info"`
	GitUrl             string            `json:"git_url"`
	Id                 int               `json:"id"`
	ImageTag           string            `json:"image_tag"`
	ImageUrl           string            `json:"image_url"`
	LoadBalancerId     int               `json:"load_balancer_id"`
	Name               string            `json:"name"`
	OtherVersion       interface{}       `json:"other_version"`
	ParentId           int               `json:"parent_id"`
	PluginVersion      PluginVersion     `json:"plugin_version"`
	PluginVersionId    int               `json:"plugin_version_id"`
	Replicas           int               `json:"replicas"`
	Resource           Resource          `json:"resource"`
	ResourceId         int               `json:"resource_id"`
	ServiceType        int               `json:"service_type"`
	Setting            interface{}       `json:"setting"`
	Status             string            `json:"status"`
	Variables          interface{}       `json:"variables"`
	Version            interface{}       `json:"version"`
}

type AutoScaler struct {
	AdvancedScheduling      interface{} `json:"advanced_scheduling"`
	Enabled                 bool        `json:"enabled"`
	HorizontalPodAutoscaler struct {
		MaxReplicas int `json:"max_replicas"`
		MinReplicas int `json:"min_replicas"`
	} `json:"horizontal_pod_autoscaler"`
	VerticalPodAutoscaler struct {
		ControlledValues string `json:"controlled_values"`
		MaxAllowedCpu    string `json:"max_allowed_cpu"`
		MaxAllowedMemory string `json:"max_allowed_memory"`
		MinAllowedCpu    string `json:"min_allowed_cpu"`
		MinAllowedMemory string `json:"min_allowed_memory"`
		Mode             string `json:"mode"`
		UpdateMode       string `json:"update_mode"`
	} `json:"vertical_pod_autoscaler"`
}

type PluginVersion struct {
	Plugin   Plugin `json:"plugin"`
	PluginId int    `json:"plugin_id"`
	Version  string `json:"version"`
}

type GitRepositoryInfo struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	HtmlUrl  string `json:"html_url"`
	CloneUrl string `json:"clone_url"`
	GitUrl   string `json:"git_url"`
}
