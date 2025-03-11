package api

import (
	"encoding/json"
	"time"
)

type Environment struct {
	CronJob        []any  `json:"CronJob"`
	InitContainers any    `json:"InitContainers"`
	Storage        any    `json:"Storage"`
	Action         string `json:"action"`
	Active         bool   `json:"active"`
	Addons         []any  `json:"addons"`
	Application    struct {
		Cluster struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"cluster"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"application"`
	ApplyImmediately   bool        `json:"apply_immediately"`
	Attributes         any         `json:"attributes"`
	AutoScaler         interface{} `json:"auto_scaler"`
	CiRequest          any         `json:"ci_request"`
	CloneEnvironment   any         `json:"clone_environment"`
	Createdat          time.Time   `json:"createdat"`
	DeploymentStrategy struct {
		RollingUpdate struct {
			MaxSurge int `json:"maxSurge"`
		} `json:"rollingUpdate"`
	} `json:"deployment_strategy"`
	ErrorMessage       any         `json:"error_message"`
	ExternalLogging    any         `json:"external_logging"`
	ExternalSecret     any         `json:"external_secret"`
	FileManagerEnabled any         `json:"file_manager_enabled"`
	GitBranch          string      `json:"git_branch"`
	GitRepositoryInfo  any         `json:"git_repository_info"`
	GitURL             string      `json:"git_url"`
	ID                 int         `json:"id"`
	ImageTag           string      `json:"image_tag"`
	ImageURL           string      `json:"image_url"`
	LoadBalancer       any         `json:"load_balancer"`
	LoadBalancerID     int         `json:"load_balancer_id"`
	Name               string      `json:"name"`
	OperatorPayload    any         `json:"operator_payload"`
	OtherVersion       interface{} `json:"other_version"`
	Parent             any         `json:"parent"`
	ParentID           int         `json:"parent_id"`
	PluginVersion      struct {
		ID     int `json:"id"`
		Plugin struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"plugin"`
		Version string `json:"version"`
	} `json:"plugin_version"`
	Replicas        int `json:"replicas"`
	RepositoryImage any `json:"repository_image"`
	Resource        struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Memory         int    `json:"memory"`
		OrganizationID int    `json:"organization_id"`
		Weight         int    `json:"weight"`
	} `json:"resource"`
	Schedules any `json:"schedules"`
	Scripts   struct {
		Build      string `json:"build"`
		CiSteps    any    `json:"ci_steps"`
		Dockerfile string `json:"dockerfile"`
		Run        string `json:"run"`
		SubDir     string `json:"sub_dir"`
	} `json:"scripts"`
	ServiceType int `json:"service_type"`
	Setting     struct {
	} `json:"setting"`
	Status string `json:"status"`
}

func (e *Environment) FromJson(d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &e)
}

type SingleEnvironment struct {
	Environment Environment `json:"environment"`
	Metadata    Metadata    `json:"metadata"`
	Overview    []Overview  `json:"overview"`
}

func (e *SingleEnvironment) FromJson(d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &e)
}

type MultipleEnvironment struct {
	Environments []Environment
}

func (m *MultipleEnvironment) FromJson(d interface{}) error {
	var envs []Environment
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &envs)
	if err != nil {
		return err
	}
	m.Environments = envs
	return nil
}

type Metadata struct {
	Envs []struct {
		Env []struct {
			Name      string `json:"name"`
			Value     string `json:"value,omitempty"`
			ValueFrom struct {
				SecretKeyRef struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"secretKeyRef"`
			} `json:"valueFrom,omitempty"`
		} `json:"env"`
		EnvFrom         interface{} `json:"envFrom"`
		Image           string      `json:"image"`
		ImagePullPolicy string      `json:"imagePullPolicy"`
		Name            string      `json:"name"`
		Ports           []struct {
			ContainerPort int    `json:"containerPort"`
			Name          string `json:"name"`
		} `json:"ports"`
		Resources struct {
			Limits struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"limits"`
			Requests struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"requests"`
		} `json:"resources"`
		SecurityContext struct {
			AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation"`
			Capabilities             struct {
				Drop []string `json:"drop"`
			} `json:"capabilities"`
			RunAsNonRoot bool `json:"runAsNonRoot"`
			RunAsUser    int  `json:"runAsUser"`
		} `json:"securityContext"`
		VolumeMounts []struct {
			MountPath string `json:"mountPath"`
			Name      string `json:"name"`
			SubPath   string `json:"subPath"`
		} `json:"volumeMounts"`
	} `json:"envs"`
	Info struct {
		FirstDeployed time.Time `json:"first_deployed"`
		LastDeployed  time.Time `json:"last_deployed"`
		Deleted       time.Time `json:"deleted "`
		Description   string    `json:"description"`
		Status        string    `json:"status"`
		Notes         string    `json:"notes"`
	} `json:"info"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	ReleaseID string `json:"releaseId"`
	Version   int    `json:"version"`
}

type Overview struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	IsRunning bool   `json:"is_running"`
	Type      string `json:"type,omitempty"`
}

type EnvironmentPresenter struct {
	Environment SingularEnvironmentPresenter `json:"environment"`
	Overview    []Overview                   `json:"overview"`
}

type SingularEnvironmentPresenter struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Status            string `json:"status"`
	GitBranch         string `json:"git_branch,omitempty"`
	GitRepositoryInfo any    `json:"git_repository_info,omitempty"`
	GitURL            string `json:"git_url,omitempty"`
	ImageTag          string `json:"image_tag,omitempty"`
	ImageURL          string `json:"image_url,omitempty"`
}

func (e *EnvironmentPresenter) FromJSON(d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &e)
}

type EnvironmentsPresenter struct {
	Environments []SingularEnvironmentPresenter `json:"environments"`
}

func (e *EnvironmentsPresenter) FromJSON(d interface{}) error {
	var all []SingularEnvironmentPresenter
	a, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = json.Unmarshal(a, &all)
	if err != nil {
		return err
	}
	e.Environments = all
	return nil
}

type EnvironmentOverview struct {
	EnvName   string `json:"env_name"`
	CPUUsages []struct {
		Metric struct {
			Namespace string `json:"namespace"`
		} `json:"metric"`
		Values [][]interface{} `json:"values"`
	} `json:"cpu_usages"`
	DataTransfer struct {
		Receive  interface{} `json:"receive"`
		Transfer interface{} `json:"transfer"`
	} `json:"data_transfer"`
	DiskUsages []struct {
		Values [][]interface{} `json:"values"`
	} `json:"disk_usages"`
	MemoryUsages []struct {
		Metric struct {
			Namespace string `json:"namespace"`
		} `json:"metric"`
		Values [][]interface{} `json:"values"`
	} `json:"memory_usages"`
	TotalCPU    int `json:"total_cpu"`
	TotalMemory int `json:"total_memory"`
	TotalPv     int `json:"total_pv"`
}

func (e *EnvironmentOverview) FromJSON(d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &e)
}
