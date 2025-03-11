package api

import (
	"encoding/json"
	"time"
)

type OrganizationList struct {
	Organizations []Organization
}

func (o *OrganizationList) FromJson(data interface{}) error {
	var orgs []Organization
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &orgs)
	if err != nil {
		return err
	}
	o.Organizations = orgs
	return nil
}

type Organization struct {
	ID                 int                  `json:"id"`
	Createdat          time.Time            `json:"createdat"`
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	Domain             string               `json:"domain"`
	Image              string               `json:"image"`
	User               User                 `json:"user"`
	UserID             int                  `json:"user_id"`
	OrganizationPlan   *OrganizationPlan    `json:"organization_plan"`
	OrganizationPlanID int                  `json:"organization_plan_id"`
	Plugins            []Plugin             `json:"plugins"`
	Members            []OrganizationMember `json:"members"`
}

type OrganizationSwitch struct {
	Organization Organization `json:"organization"`
	Token        string       `json:"token"`
}

func (o *OrganizationSwitch) FromJson(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteData, &o)
}

type OrganizationPlan struct {
	ID         int       `json:"id"`
	Createdat  time.Time `json:"createdat"`
	Name       string    `json:"name"`
	Cluster    int       `json:"cluster"`
	Memory     int       `json:"memory"`
	Cores      int       `json:"cores"`
	NoOfUser   int       `json:"no_of_user"`
	Price      int       `json:"price"`
	Weight     int       `json:"weight"`
	Attributes string    `json:"attributes"`
	Active     bool      `json:"active"`
}

func (o *Organization) FromJson(data interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteData, &o)
}

type Packages struct {
	Chart       string   `json:"chart"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	Name        string   `json:"name"`
	Namespace   string   `json:"namespace"`
	Optional    bool     `json:"optional"`
	Title       string   `json:"title"`
	RequiredDNS bool     `json:"required_dns,omitempty"`
	Needs       []string `json:"needs,omitempty"`
}

type Repositories struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Templates struct {
	Name     string   `json:"name"`
	Packages []string `json:"packages"`
}

type PackageConfig struct {
	Packages     []Packages     `json:"packages"`
	Repositories []Repositories `json:"repositories"`
	Templates    []Templates    `json:"templates"`
}

type ClusterList struct {
	ID                    int         `json:"id"`
	Createdat             time.Time   `json:"createdat"`
	ClusterName           string      `json:"cluster_name"`
	ClusterVersion        string      `json:"cluster_version"`
	Region                string      `json:"region"`
	Zone                  interface{} `json:"zone"`
	Provider              string      `json:"provider"`
	ProviderName          string      `json:"provider_name"`
	VpcName               string      `json:"vpc_name"`
	NetworkCidr           string      `json:"network_cidr"`
	NetworkPolicy         bool        `json:"network_policy"`
	PvcWriteMany          bool        `json:"pvc_write_many"`
	Active                bool        `json:"active"`
	Status                string      `json:"status"`
	Type                  string      `json:"type"`
	TLS                   string      `json:"tls"`
	NfsDetail             interface{} `json:"nfs_detail"`
	RegionalCluster       bool        `json:"regional_cluster"`
	RemoveDefaultNodePool bool        `json:"remove_default_node_pool"`
	NodeGroupCount        int         `json:"node_group_count"`
	NodeGroupDetail       interface{} `json:"node_group_detail"`
	OrganizationID        int         `json:"organization_id"`
	Cluster               struct {
		ID                  int         `json:"id"`
		Createdat           time.Time   `json:"createdat"`
		Name                string      `json:"name"`
		Context             string      `json:"context"`
		ConfigPath          string      `json:"configPath"`
		Region              string      `json:"region"`
		Provider            string      `json:"provider"`
		Zone                string      `json:"zone"`
		DNSID               int         `json:"dns_id"`
		Labels              string      `json:"labels"`
		Nodes               int         `json:"nodes"`
		PvCapacity          int         `json:"pv_capacity"`
		Weight              int         `json:"weight"`
		ImageRegistryID     int         `json:"image_registry_id"`
		Active              bool        `json:"active"`
		OrganizationID      int         `json:"organization_id"`
		TotalMemory         int         `json:"total_memory"`
		CloudStorage        interface{} `json:"cloud_storage"`
		ProvisionPercentage float64     `json:"provision_percentage"`
		ClusterRequest      interface{} `json:"cluster_request"`
		ClusterRequestID    int         `json:"cluster_request_id"`
		Color               string      `json:"color"`
		StorageClass        string      `json:"storage_class"`
	} `json:"cluster"`
	ClusterID       int    `json:"cluster_id"`
	ProjectID       string `json:"project_id"`
	Credentials     string `json:"credentials"`
	AccessKey       string `json:"access_key"`
	SecretKey       string `json:"secret_key"`
	SubnetCidrRange string `json:"subnet_cidr_range"`
	ErrorMessage    struct {
		ClusterID     int       `json:"cluster_id"`
		Code          int       `json:"code"`
		EnvironmentID int       `json:"environment_id"`
		Message       string    `json:"message"`
		Source        string    `json:"source"`
		Time          time.Time `json:"time"`
	} `json:"error_message"`
}
