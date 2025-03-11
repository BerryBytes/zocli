package requests

type MemberCreateRequest struct {
	Email string `json:"email"`
	Role  int    `json:"role"`
}

type ProjectCreateRequest struct {
	Name           string `json:"name"`
	SubscriptionID int    `json:"subscription_id"`
	Description    string `json:"description"`
	ProjectCode    string `json:"project_code"`
	BaseDomain     string `json:"base_domain"`
	ClusterScope   int    `json:"cluster_scope"`
	Region         string `json:"region"`
	Logging        string `json:"logging"`
	Monitoring     string `json:"monitoring"`
	OptimizeCost   bool   `json:"optimize_cost"`
	DedicatedLb    bool   `json:"dedicated_lb"`
	Tags           string `json:"tags"`
}

type AppCreateRequest struct {
	Name                string  `json:"name"`
	ProjectID           int     `json:"project_id"`
	PluginID            uint64  `json:"plugin_id"`
	ClusterID           uint64  `json:"cluster_id,omitempty"`
	ChartID             string  `json:"chart_id,omitempty"`
	OwnerId             uint64  `json:"owner_id,omitempty"`
	GitRepository       string  `json:"git_repository"`
	GitRepoUrl          *string `json:"git_repo_url"`
	GitToken            string  `json:"git_token"`
	Region              string  `json:"region"`
	GitService          string  `json:"git_service"`
	ImageUrl            string  `json:"image_url"`
	ImageNamespace      string  `json:"image_namespace"`
	ImageRepo           string  `json:"image_repo"`
	ImageService        string  `json:"image_service"`
	ServiceType         int     `json:"service_type"` // 0-template/1-git/2-image/3-helm/4-operator
	OperatorPackageName string  `json:"operator_package_name"`
}

type OrganizationCreateRequest struct {
	Name                string `json:"name"`
	OrganizationPlainID int    `json:"organization_plan_id"`
	Domain              string `json:"domain"`
	Description         string `json:"description"`
}

type EnvCreateRequest struct {
	Name            string  `json:"name" yaml:"name"`
	ApplicationID   int     `json:"application_id" yaml:"application_id"`
	ResoruceID      int     `json:"resource_id" yaml:"resource_id"`
	PluginVersionID int     `json:"plugin_version_id" yaml:"plugin_version_id"`
	Replicas        int     `json:"replicas" yaml:"replicas"`
	Version         Version `json:"version" yaml:"version"`
}

type Version struct {
	Name string `json:"name" yaml:"name"`
	Repo string `json:"repo" yaml:"repo"`
	Tag  string `json:"tag" yaml:"tag"`
}
